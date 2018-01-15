package udevgo

import (
	"encoding/json"
	"errors"
	"go/build"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
)

type Pkg struct {
	ApproxLoC int // 0 unless/until calling CountLoC()
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

	dependants        []string
	importers         []string
	goFilePathsNoTest []string
	goFilePathsAll    []string
}

func (me *Pkg) IsSortedPriorTo(pkg interface{}) bool { return me.ImportPath < pkg.(*Pkg).ImportPath }
func (me *Pkg) String() string                       { return me.ImportPath }
func (me *Pkg) IsSortedPriorToByDeps(cmp *Pkg) bool {
	return !uslice.StrHas(me.Deps, cmp.ImportPath)
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
	PkgsErrs  []*Pkg

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

// func AllFinalDependants(origpkgimppaths []string) (depimppaths []string) {
// 	opkgs := map[string]*Pkg{}
// 	for _, origpkgimppath := range origpkgimppaths {
// 		if pkg := PkgsByImP[origpkgimppath]; pkg != nil {
// 			opkgs[pkg.ImportPath] = pkg
// 		}
// 	}
// 	//	grab all dependants of each origpkg
// 	for _, opkg := range opkgs {
// 		for _, depimppath := range opkg.Dependants() {
// 			if _, ignore := opkgs[depimppath]; !ignore {
// 				uslice.StrAppendUnique(&depimppaths, depimppath)
// 			}
// 		}
// 	}
// 	//	shake out unnecessary mentions
// 	depimppaths = ShakeOutIntermediateDeps(depimppaths)
// 	return
// }

// func ShakeOutIntermediateDeps(pkgimppaths []string) []string {
// 	pkgimppaths = uslice.StrWithout(pkgimppaths, false, "")
// 	for oncemore := true; oncemore; {
// 		oncemore = false
// 		for _, pkgimppath := range pkgimppaths {
// 			if pkg := PkgsByImP[pkgimppath]; pkg != nil {
// 				for _, subimppath := range pkg.Deps {
// 					if uslice.StrHas(pkgimppaths, subimppath) {
// 						pkgimppaths = uslice.StrWithout(pkgimppaths, false, subimppath)
// 						oncemore = true
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return pkgimppaths
// }

// func ShakeOutIntermediateDepsViaDir(dirrelpaths []string, basedirpath string) []string {
// 	pkgimppaths := ShakeOutIntermediateDeps(uslice.StrMap(dirrelpaths, func(drp string) string {
// 		dir := filepath.Join(basedirpath, drp)
// 		if runtime.GOOS == "windows" {
// 			dir = strings.ToLower(dir)
// 		}
// 		return PkgsByDir[dir].ImportPath
// 	}))
// 	return uslice.StrMap(pkgimppaths, func(pkgimppath string) (dirrelpath string) {
// 		dirrelpath, _ = filepath.Rel(basedirpath, PkgsByImP[pkgimppath].Dir)
// 		return
// 	})
// }

// func DependantsOn(pkgdirpath string) (pkgimppaths []string) {
// 	if pkg := PkgsByDir[pkgdirpath]; pkg != nil {
// 		pkgimppaths = pkg.Dependants()
// 	}
// 	return
// }

// func ImportersOf(pkgdirpath string, basedirpath string) (pkgimppaths []string) {
// 	if pkg := PkgsByDir[pkgdirpath]; pkg != nil {
// 		pkgimppaths = pkg.Importers(basedirpath)
// 	}
// 	return
// }

func (me *Pkg) Dependants() []string {
	if me.dependants == nil {
		pkgsMutex.Lock()
		defer pkgsMutex.Unlock()
		me.dependants = []string{}
		for _, pkg := range PkgsByDir {
			if uslice.StrHas(pkg.Deps, me.ImportPath) {
				me.dependants = append(me.dependants, pkg.ImportPath)
			}
		}
	}
	return me.dependants
}

func (me *Pkg) Importers() []string {
	if me.importers == nil {
		pkgsMutex.Lock()
		defer pkgsMutex.Unlock()
		me.importers = []string{}
		for _, pkg := range PkgsByDir {
			if uslice.StrHas(pkg.Imports, me.ImportPath) {
				me.importers = append(me.importers, pkg.ImportPath)
			}
		}
	}
	return me.importers
}

func (me *Pkg) GoFilePaths(inclTests bool) []string {
	l, gofilepaths := len(me.GoFiles), me.goFilePathsNoTest
	if inclTests {
		l, gofilepaths = l+len(me.TestGoFiles), me.goFilePathsAll
	}
	if l != len(gofilepaths) {
		slices := [][]string{me.GoFiles}
		if gofilepaths = make([]string, 0, l); inclTests {
			slices = append(slices, me.TestGoFiles)
		}
		for _, filenames := range slices {
			for _, fname := range filenames {
				gofilepaths = append(gofilepaths, filepath.Join(me.Dir, fname))
			}
		}
		if inclTests {
			me.goFilePathsAll = gofilepaths
		} else {
			me.goFilePathsNoTest = gofilepaths
		}
	}
	return gofilepaths
}

func (me *Pkg) CountLoC() {
	me.ApproxLoC = 0
	for _, gfp := range me.GoFilePaths(false) {
		incomment := false
		for _, ln := range ustr.Split(ufs.ReadTextFile(gfp, false, ""), "\n") {
			if ln = ustr.Trim(ln); len(ln) > 0 { // yeap, will bug for pointlessly unusual multiline-comment "compositions"
				if ustr.Suff(ln, "*/") {
					incomment = false
				} else if ustr.Pref(ln, "/*") {
					incomment = true
				} else if (!incomment) && !ustr.Pref(ln, "//") {
					me.ApproxLoC++
				}
			}
		}
	}
}

func GuruMinimalScopeFor(goFilePath string) (pkgScope string, shouldRefresh bool) {
	var pkgs []*Pkg
	pkgs, shouldRefresh = PkgsForFiles(goFilePath) // we pass only 1 filepath so len(pkgs) will be at-most 1:
	var pkg *Pkg
	if len(pkgs) > 0 {
		pkg = pkgs[0]
	}
	if pkg == nil && PkgsByDir != nil {
		for dp, lastdp := filepath.Dir(filepath.Dir(goFilePath)), ""; dp != "" && dp != lastdp; lastdp, dp = dp, filepath.Dir(dp) {
			if pkg = PkgsByDir[dp]; pkg != nil {
				break
			}
		}
	}
	if pkg != nil {
		check := func(p *Pkg) *Pkg {
			if p.IsCommand() || len(p.TestGoFiles) > 0 {
				return p
			} else if PkgsByDir != nil {
				for _, sub := range PkgsByDir {
					if sub == nil {
						println("WUTUTUT?")
					} else if strings.HasPrefix(sub.Dir, p.Dir+string(filepath.Separator)) && (sub.IsCommand() || len(sub.TestGoFiles) > 0) {
						return sub
					}
				}
			}
			return nil
		}
		for pkg != nil {
			if check(pkg) != nil {
				break
			} else if dp, lastdp := filepath.Dir(pkg.Dir), ""; PkgsByDir != nil {
				for pkg = nil; dp != "" && dp != lastdp; lastdp, dp = dp, filepath.Dir(dp) {
					if pkg = PkgsByDir[dp]; pkg != nil {
						break
					}
				}
			} else {
				pkg = nil
			}
		}
	}
	if pkg != nil {
		pkgScope = pkg.ImportPath
	}
	return
}

func PkgsForFiles(filePaths ...string) (pkgs []*Pkg, shouldRefresh bool) {
	if all := PkgsByDir; all == nil {
		shouldRefresh = true
	} else {
		for _, fp := range filePaths {
			alreadyhave, dp := false, filepath.Dir(fp)
			for i := range pkgs {
				if alreadyhave = (pkgs[i].Dir == dp); alreadyhave {
					break
				}
			}
			if !alreadyhave {
				if pkg := all[dp]; pkg != nil {
					pkgs = append(pkgs, pkg)
				} else if (!shouldRefresh) && strings.ToLower(filepath.Ext(fp)) == ".go" {
					for _, gp := range AllGoPaths() {
						if shouldRefresh = strings.HasPrefix(dp, filepath.Join(gp, "src")); shouldRefresh {
							break
						}
					}
				}
			}
		}
	}
	return
}

// func (me *Pkg) Importers(basedirpath string) (pkgimppaths []string) {
// 	pkgsMutex.Lock()
// 	defer pkgsMutex.Unlock()
// 	for _, pkg := range PkgsByDir {
// 		if uslice.StrHas(pkg.Imports, me.ImportPath) {
// 			if len(basedirpath) == 0 {
// 				pkgimppaths = append(pkgimppaths, pkg.ImportPath)
// 			} else if reldirpath, _ := filepath.Rel(basedirpath, pkg.Dir); len(reldirpath) > 0 {
// 				pkgimppaths = append(pkgimppaths, reldirpath)
// 			}
// 		}
// 	}
// 	return
// }

func RefreshPkgs() error {
	pkgsbydir, pkgsbyimp, pkgserrs := map[string]*Pkg{}, map[string]*Pkg{}, []*Pkg{}

	if cmdout, cmderr, err := urun.CmdExec("go", "list", "-e", "-json", "all"); err != nil {
		return err
	} else if cmderr != "" && !ustr.Pref(strings.ToLower(cmderr), "warning: ") {
		return errors.New(cmderr)
	} else if jsonobjstrs := ustr.Split(ustr.Trim(cmdout), "}\n{"); len(jsonobjstrs) > 0 {
		jsonobjstrs[0] = jsonobjstrs[0][1:]
		idxlast := len(jsonobjstrs) - 1
		jsonobjstrs[idxlast] = jsonobjstrs[idxlast][:len(jsonobjstrs[idxlast])-1]
		for _, jsonobjstr := range jsonobjstrs {
			var pkg Pkg
			if err := json.Unmarshal([]byte("{"+jsonobjstr+"}"), &pkg); err != nil {
				return err
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
		defer func() { pkgsMutex.Unlock(); go pkgAfterRefreshUpdateGuruScopeExcls() }()
		PkgsByDir, PkgsByImP, PkgsErrs = pkgsbydir, pkgsbyimp, pkgserrs
	}
	return nil
}

func pkgAfterRefreshUpdateGuruScopeExcls() {
	pats := make([]string, 0, len(GuruScopeExclPkgs))
	for gsxp, excl := range GuruScopeExclPkgs {
		if excl && strings.HasSuffix(gsxp, "/...") {
			pats = append(pats, gsxp[:len(gsxp)-3])
		}
	}
	guruscopeexclpkgs := make(map[string]bool, len(PkgsByImP))
	for _, pkg := range PkgsByImP {
		if pkg.Error != nil || len(pkg.Errs) > 0 || len(pkg.DepsErrors) > 0 || pkg.Incomplete || len(pkg.InvalidGoFiles) > 0 {
			guruscopeexclpkgs[pkg.ImportPath] = true
			for _, d := range pkg.Dependants() {
				guruscopeexclpkgs[d] = true
			}
		}
	}
	coveredbypat := false
	for pkgimppath, _ := range guruscopeexclpkgs {
		for _, pat := range pats {
			if coveredbypat = strings.HasPrefix(pkgimppath, pat) || pkgimppath == pat[:len(pat)-1]; coveredbypat {
				break
			}
		}
		if !coveredbypat {
			GuruScopeExclPkgs[pkgimppath] = true
		}
	}
}

func PkgImpPathsToNamesInLn(ln string, curPkgDir string) string {
	if PkgsByImP != nil {
		if isla := strings.IndexRune(ln, '/'); isla >= 0 {
			isla1 := isla + 1
			if idot := strings.IndexRune(ln[isla1:], '.'); idot > 0 {
				ipos, imppath := 0, ln[:isla1+idot]
				for ir, r := range imppath {
					if r == '/' {
						break
					} else if r == '*' || r == '[' || r == ']' || r == '(' || r == ')' || r == '{' || r == '}' || unicode.IsSpace(r) {
						ipos = ir + 1
					}
				}
				imppath = imppath[ipos:]
				if pkg := PkgsByImP[imppath]; pkg != nil {
					if pkg.Dir != curPkgDir {
						ln = strings.Replace(ln, imppath+".", pkg.Name+".", -1)
					} else {
						ln = strings.Replace(ln, imppath+".", "", -1)
					}
					ln = PkgImpPathsToNamesInLn(ln, curPkgDir)
				}
			}
		}
	}
	return ln
}

func PkgsByName(name string) (pkgImpPaths []string) {
	if pkgs := PkgsByImP; pkgs != nil {
		for _, pkg := range pkgs {
			if pkg.Name == name {
				pkgImpPaths = append(pkgImpPaths, pkg.ImportPath)
			}
		}
	}
	return
}
