package udevgo

import (
	"encoding/json"
	"go/build"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
)

type Pkg struct {
	build.Package
	Errs udev.SrcMsgs

	//	the below all copied over from `go list` src because that cmd outputs this stuff but our above build.Package doesn't have it:

	Deps        []string        `json:",omitempty"` // all (recursively) imported dependencies
	Target      string          `json:",omitempty"` // install path
	Shlib       string          `json:",omitempty"` // the shared library that contains this package (only set when -linkshared)
	StaleReason string          `json:",omitempty"` // why is Stale true?
	Stale       bool            `json:",omitempty"` // would 'go install' do anything for this package?
	Standard    bool            `json:",omitempty"` // is this package part of the standard Go library?
	Incomplete  bool            `json:",omitempty"` // was there an error loading this package or dependencies?
	Error       *PackageError   `json:",omitempty"` // error loading this package (not dependencies)
	DepsErrors  []*PackageError `json:",omitempty"` // errors loading dependencies
}

//	copied over from `go list` src because that cmd outputs this stuff but one cannot import it from anywhere
type PackageError struct {
	ImportStack []string // shortest path from package named on command line to this one
	Pos         string   // position of error (if present, file:line:col)
	Err         string   // the error itself
}

var (
	PkgsByDir map[string]*Pkg
	PkgsByImP map[string]*Pkg
	_PkgsErrs []*Pkg

	pkgsMutex       sync.Mutex
	ShortenImpPaths *strings.Replacer
)

// func ShortImpP(pkgimppath string) string {
// 	if len(SnipImp) > 0 && ustr.Pref(pkgimppath, SnipImp) {
// 		if PkgsByImP != nil {
// 			if pkg := PkgsByImP[pkgimppath]; pkg != nil {
// 				return pkg.Name
// 			}
// 		}
// 		return strings.TrimPrefix(pkgimppath, SnipImp)
// 	}
// 	return pkgimppath
// }

// func ShortenImPs(val string) string {
// 	if len(SnipImp) > 0 {
// 		l := len(SnipImp)
// 		for {
// 			if i := ustr.Idx(val, SnipImp); i < 0 {
// 				return val
// 			} else if j := ustr.Idx(val[i+l:], "."); j < 0 {
// 				return val
// 			} else {
// 				val = val[:i] + ShortImpP(val[i:i+l+j]) + val[i+l+j:]
// 			}
// 		}
// 	}
// 	return val
// }

func AllFinalDependants(origpkgimppaths []string) (depimppaths []string) {
	opkgs := map[string]*Pkg{}
	for _, origpkgimppath := range origpkgimppaths {
		if pkg := PkgsByImP[origpkgimppath]; pkg != nil {
			opkgs[pkg.ImportPath] = pkg
		}
	}
	//	grab all dependants of each origpkg
	for _, opkg := range opkgs {
		for _, depimppath := range opkg.Dependants() {
			if _, ignore := opkgs[depimppath]; !ignore {
				uslice.StrAppendUnique(&depimppaths, depimppath)
			}
		}
	}
	//	shake out unnecessary mentions
	depimppaths = ShakeOutIntermediateDeps(depimppaths)
	return
}

func ShakeOutIntermediateDeps(pkgimppaths []string) []string {
	pkgimppaths = uslice.StrWithout(pkgimppaths, false, "")
	for oncemore := true; oncemore; {
		oncemore = false
		for _, pkgimppath := range pkgimppaths {
			if pkg := PkgsByImP[pkgimppath]; pkg != nil {
				for _, subimppath := range pkg.Deps {
					if uslice.StrHas(pkgimppaths, subimppath) {
						pkgimppaths = uslice.StrWithout(pkgimppaths, false, subimppath)
						oncemore = true
						break
					}
				}
			}
		}
	}
	return pkgimppaths
}

func ShakeOutIntermediateDepsViaDir(dirrelpaths []string, basedirpath string) []string {
	pkgimppaths := ShakeOutIntermediateDeps(uslice.StrMap(dirrelpaths, func(drp string) string {
		dir := filepath.Join(basedirpath, drp)
		if runtime.GOOS == "windows" {
			dir = strings.ToLower(dir)
		}
		return PkgsByDir[dir].ImportPath
	}))
	return uslice.StrMap(pkgimppaths, func(pkgimppath string) (dirrelpath string) {
		dirrelpath, _ = filepath.Rel(basedirpath, PkgsByImP[pkgimppath].Dir)
		return
	})
}

func DependantsOn(pkgdirpath string) (pkgimppaths []string) {
	if pkg := PkgsByDir[pkgdirpath]; pkg != nil {
		pkgimppaths = pkg.Dependants()
	}
	return
}

func ImportersOf(pkgdirpath string, basedirpath string) (pkgimppaths []string) {
	if pkg := PkgsByDir[pkgdirpath]; pkg != nil {
		pkgimppaths = pkg.Importers(basedirpath)
	}
	return
}

func (me *Pkg) Dependants() (pkgimppaths []string) {
	pkgsMutex.Lock()
	defer pkgsMutex.Unlock()
	for _, pkg := range PkgsByDir {
		if uslice.StrHas(pkg.Deps, me.ImportPath) {
			pkgimppaths = append(pkgimppaths, pkg.ImportPath)
		}
	}
	return
}

func (me *Pkg) Importers(basedirpath string) (pkgimppaths []string) {
	pkgsMutex.Lock()
	defer pkgsMutex.Unlock()
	for _, pkg := range PkgsByDir {
		if uslice.StrHas(pkg.Imports, me.ImportPath) {
			if len(basedirpath) == 0 {
				pkgimppaths = append(pkgimppaths, pkg.ImportPath)
			} else if reldirpath, _ := filepath.Rel(basedirpath, pkg.Dir); len(reldirpath) > 0 {
				pkgimppaths = append(pkgimppaths, reldirpath)
			}
		}
	}
	return
}

func RefreshPkgs() {
	pkgsbydir, pkgsbyimp, pkgserrs := map[string]*Pkg{}, map[string]*Pkg{}, []*Pkg{}

	if cmdout, _, err := urun.CmdExecStdin("", "", "go", "list", "-e", "-json", "all"); err != nil {
		panic(err)
	} else if jsonobjstrs := ustr.Split(ustr.Trim(cmdout), "}\n{"); len(jsonobjstrs) > 0 {
		jsonobjstrs[0] = jsonobjstrs[0][1:]
		idxlast := len(jsonobjstrs) - 1
		jsonobjstrs[idxlast] = jsonobjstrs[idxlast][:len(jsonobjstrs[idxlast])-1]
		for _, jsonobjstr := range jsonobjstrs {
			var pkg Pkg
			if err := json.Unmarshal([]byte("{"+jsonobjstr+"}"), &pkg); err != nil {
				panic(err)
			} else {
				if runtime.GOOS == "windows" {
					pkg.Dir = strings.ToLower(pkg.Dir)
				}
				pkgsbydir[pkg.Dir] = &pkg
				pkgsbyimp[pkg.ImportPath] = &pkg
				if pkgerror := pkg.Error; pkgerror != nil {
					var pkgerrs udev.SrcMsgs
					if multerrs := ustr.Split(pkgerror.Err, "\n"); len(multerrs) > 0 {
						pkgerrs = append(pkgerrs, udev.SrcMsgsFromLns(multerrs)...)
					} else if errpos := strings.Split(pkgerror.Pos, ":"); len(errpos) >= 3 {
						fpath := errpos[0]
						pos1ln, _ := strconv.Atoi(errpos[1])
						pos1ch, _ := strconv.Atoi(errpos[2])
						pkgerrs = append(pkgerrs, &udev.SrcMsg{Pos1Ln: pos1ln - 1, Pos1Ch: pos1ch - 1, Ref: fpath, Msg: pkgerror.Err})
					} else {
						fpath := errpos[0]
						pkgerrs = append(pkgerrs, &udev.SrcMsg{Ref: fpath, Msg: pkgerror.Err})
					}
					pkg.Errs = pkgerrs
					pkgserrs = append(pkgserrs, &pkg)
				}
			}
		}

		repls := make([]string, 0, len(pkgsbyimp)*2)
		for imp, pkg := range pkgsbyimp {
			repls = append(repls, imp, pkg.Name)
		}
		ShortenImpPaths = strings.NewReplacer(repls...)

		pkgsMutex.Lock()
		defer pkgsMutex.Unlock()
		PkgsByDir, PkgsByImP, _PkgsErrs = pkgsbydir, pkgsbyimp, pkgserrs
	}
	return
}
