package udevgo

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/metaleap/go-util/dev"
	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/run"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
)

var (
	GoVersion string
	GoPaths   []string

	Has_godoc       bool
	Has_gofmt       bool
	Has_goimports   bool
	Has_golint      bool
	Has_guru        bool
	Has_checkvar    bool
	Has_checkalign  bool
	Has_checkstruct bool
	Has_errcheck    bool
	Has_ineffassign bool
	Has_interfacer  bool
	Has_unparam     bool
	Has_unconvert   bool
	Has_maligned    bool
	Has_goconst     bool
	Has_gosimple    bool
	Has_unused      bool
	Has_staticcheck bool
	Has_gorename    bool
	Has_godef       bool
	Has_gocode      bool
	Has_gogetdoc    bool

	SnipImp string
)

func HasGoDevEnv() bool {
	var cmdout string
	var err error

	if len(GoPaths) > 0 && len(GoVersion) > 0 {
		return true
	}

	//  GoVersion
	if cmdout, err = urun.CmdExec("go", "tool", "dist", "version"); err == nil && len(cmdout) > 0 {
		GoVersion = cmdout
	} else if cmdout, err = urun.CmdExec("go", "version"); err == nil && len(cmdout) > 0 {
		GoVersion = strings.TrimPrefix(cmdout, "go version ")
		GoVersion = strings.SplitAfter(GoVersion, " ")[0]
	} else {
		return false
	}
	GoVersion = strings.TrimPrefix(GoVersion, "go")

	//  GoPaths
	if cmdout, err = urun.CmdExec("go", "env", "GOPATH"); err != nil {
		GoVersion = ""
		GoPaths = nil
		return false
	}
	GoPaths = filepath.SplitList(cmdout)
	for i, gopath := range GoPaths {
		if !ufs.DirExists(gopath) {
			GoPaths[i] = ""
		}
	}
	if GoPaths = uslice.StrWithout(GoPaths, true, ""); len(GoPaths) == 0 {
		GoVersion = ""
		GoPaths = nil
		return false
	}
	for _, gopath := range GoPaths {
		if ustr.Pref(udev.SrcDir, gopath) {
			SnipImp = filepath.ToSlash(strings.Trim(udev.SrcDir[len(filepath.Join(gopath, "src")):], "/\\")) + "/"
		}
	}

	//  OKAY! we ran go command and have 1-or-more GOPATHs, the rest is optional
	stdargs := []string{"-help"}
	urun.CmdsTryStart(map[string]*urun.CmdTry{
		"gofmt":     {Ran: &Has_gofmt, Args: stdargs},
		"goimports": {Ran: &Has_goimports, Args: stdargs},

		"golint":      {Ran: &Has_golint, Args: stdargs},
		"ineffassign": {Ran: &Has_ineffassign, Args: stdargs},
		"errcheck":    {Ran: &Has_errcheck, Args: stdargs},
		"aligncheck":  {Ran: &Has_checkalign, Args: stdargs},
		"structcheck": {Ran: &Has_checkstruct, Args: stdargs},
		"varcheck":    {Ran: &Has_checkvar, Args: stdargs},
		"interfacer":  {Ran: &Has_interfacer, Args: stdargs},
		"unparam":     {Ran: &Has_unparam, Args: stdargs},
		"unconvert":   {Ran: &Has_unconvert, Args: stdargs},
		"maligned":    {Ran: &Has_maligned, Args: stdargs},
		"gosimple":    {Ran: &Has_gosimple, Args: stdargs},
		"staticcheck": {Ran: &Has_staticcheck, Args: stdargs},
		"unused":      {Ran: &Has_unused, Args: stdargs},

		"gorename": {Ran: &Has_gorename, Args: stdargs},
		"godef":    {Ran: &Has_godef, Args: stdargs},
		"gocode":   {Ran: &Has_gocode, Args: stdargs},
		"guru":     {Ran: &Has_guru, Args: stdargs},
		"gogetdoc": {Ran: &Has_gogetdoc, Args: stdargs},
		"godoc":    {Ran: &Has_godoc, Args: stdargs},
		"goconst":  {Ran: &Has_goconst, Args: stdargs},
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
