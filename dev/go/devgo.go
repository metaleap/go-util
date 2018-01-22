package udevgo

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
)

var (
	GoVersion      string
	GoVersionShort string
	GoPaths        []string

	Has_godoc        bool
	Has_gofmt        bool
	Has_goimports    bool
	Has_goreturns    bool
	Has_guru         bool
	Has_gorename     bool
	Has_godef        bool
	Has_gogetdoc     bool
	Has_gocode       bool
	Has_structlayout bool
	Has_godocdown    bool

	Has_golint      bool
	Has_checkvar    bool
	Has_checkalign  bool
	Has_checkstruct bool
	Has_errcheck    bool
	Has_ineffassign bool
	Has_interfacer  bool
	Has_unparam     bool
	Has_unindent    bool
	Has_unconvert   bool
	Has_maligned    bool
	Has_goconst     bool
	Has_gosimple    bool
	Has_unused      bool
	Has_staticcheck bool
	Has_deadcode    bool
)

func HasGoDevEnv() bool {
	var cmdout, cmderr string
	var err error

	if len(GoPaths) > 0 && GoVersion != "" {
		return true
	}

	//  GoVersion
	if cmdout, cmderr, err = urun.CmdExec("go", "tool", "dist", "version"); err == nil && cmderr == "" && cmdout != "" {
		GoVersion = strings.TrimSpace(cmdout)
	}
	if GoVersion == "" {
		if cmdout, cmderr, err = urun.CmdExec("go", "version"); err == nil && cmderr == "" && cmdout != "" {
			if GoVersion = strings.TrimPrefix(strings.TrimSpace(cmdout), "go version "); GoVersion != "" {
				GoVersion = strings.SplitAfter(GoVersion, " ")[0]
			}
		}
	}
	if GoVersion = strings.TrimPrefix(GoVersion, "go"); GoVersion == "" {
		return false
	}

	//  GoPaths
	if cmdout, cmderr, err = urun.CmdExec("go", "env", "GOPATH"); err != nil || cmderr != "" {
		GoVersion = ""
		GoPaths = nil
		return false
	}
	GoPaths = filepath.SplitList(strings.TrimSpace(cmdout))
	for i, gopath := range GoPaths {
		if !ufs.DirExists(gopath) {
			GoPaths[i] = ""
		}
	}
	if GoPaths = uslice.StrWithout(GoPaths, true, ""); len(GoPaths) == 0 {
		GoVersion = ""
		GoPaths = nil
		return false
	} else if cmdout, cmderr, err = urun.CmdExec("go", "env", "GOROOT"); err == nil && cmderr == "" && cmdout != "" {
		if gorootdirpath := strings.TrimSpace(cmdout); gorootdirpath != "" && ufs.DirExists(gorootdirpath) && !uslice.StrHas(GoPaths, gorootdirpath) {
			GoPaths = append(GoPaths, gorootdirpath)
		}
	}

	i, l := strings.IndexRune(GoVersion, '.'), strings.LastIndex(GoVersion, ".")
	for GoVersionShort = GoVersion; l > i; l = strings.LastIndex(GoVersionShort, ".") {
		GoVersionShort = GoVersionShort[:l]
	}

	//  OKAY! we ran go command and have 1-or-more GOPATHs, the rest is optional
	stdargs := []string{"-help"}
	urun.CmdsTryStart(map[string]*urun.CmdTry{
		"gofmt":     {Ran: &Has_gofmt, Args: stdargs},
		"goimports": {Ran: &Has_goimports, Args: stdargs},
		"goreturns": {Ran: &Has_goreturns, Args: stdargs},

		"golint":      {Ran: &Has_golint, Args: stdargs},
		"ineffassign": {Ran: &Has_ineffassign, Args: stdargs},
		"errcheck":    {Ran: &Has_errcheck, Args: stdargs},
		"aligncheck":  {Ran: &Has_checkalign, Args: stdargs},
		"structcheck": {Ran: &Has_checkstruct, Args: stdargs},
		"varcheck":    {Ran: &Has_checkvar, Args: stdargs},
		"interfacer":  {Ran: &Has_interfacer, Args: stdargs},
		"unparam":     {Ran: &Has_unparam, Args: stdargs},
		"unindent":    {Ran: &Has_unindent, Args: stdargs},
		"unconvert":   {Ran: &Has_unconvert, Args: stdargs},
		"maligned":    {Ran: &Has_maligned, Args: stdargs},
		"gosimple":    {Ran: &Has_gosimple, Args: stdargs},
		"staticcheck": {Ran: &Has_staticcheck, Args: stdargs},
		"unused":      {Ran: &Has_unused, Args: stdargs},
		"deadcode":    {Ran: &Has_deadcode, Args: stdargs},

		"structlayout": {Ran: &Has_structlayout, Args: stdargs},
		"gorename":     {Ran: &Has_gorename, Args: stdargs},
		"godef":        {Ran: &Has_godef, Args: stdargs},
		"gocode":       {Ran: &Has_gocode, Args: stdargs},
		"guru":         {Ran: &Has_guru, Args: stdargs},
		"gogetdoc":     {Ran: &Has_gogetdoc, Args: stdargs},
		"godocdown":    {Ran: &Has_godocdown, Args: stdargs},
		"godoc":        {Ran: &Has_godoc, Args: stdargs},
		"goconst":      {Ran: &Has_goconst, Args: stdargs},
	})
	return true
}

//	Returns all paths listed in the `GOPATH` environment variable, for users who don't care about calling HasGoDevEnv.
func AllGoPaths() []string {
	if len(GoPaths) == 0 {
		GoPaths = filepath.SplitList(os.Getenv("GOPATH"))
	}
	return GoPaths
}

func DirPathToImportPath(dirpath string) string {
	for _, gopath := range AllGoPaths() {
		if strings.HasPrefix(dirpath, gopath) {
			return dirpath[len(filepath.Join(gopath, "src"))+1:]
		}
	}
	return ""
}

//	Returns the `path/filepath.Join`-ed full directory path for a specified `$GOPATH/src` sub-directory.
//	Example: `util.GopathSrc("tools", "importers", "sql")` yields `c:\gd\src\tools\importers\sql` if `$GOPATH` is `c:\gd`.
func GopathSrc(subDirNames ...string) (gps string) {
	gp := []string{"", "src"}
	for _, goPath := range AllGoPaths() { // in 99% of setups there's only 1 GOPATH, but hey..
		gp[0] = goPath
		if gps = filepath.Join(append(gp, subDirNames...)...); ufs.DirExists(gps) {
			break
		}
	}
	return
}

//	Returns the `path/filepath.Join`-ed full directory path for a specified `$GOPATH/src/github.com` sub-directory.
//	Example: `util.GopathSrcGithub("go-util", "num")` yields `c:\gd\src\github.com\go-util\num` if `$GOPATH` is `c:\gd`.
func GopathSrcGithub(gitHubName string, subDirNames ...string) string {
	return GopathSrc(append([]string{"github.com", gitHubName}, subDirNames...)...)
}
