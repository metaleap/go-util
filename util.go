package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	goPaths       [][]string
	goPathsLenIs1 bool
)

func init() {
	for _, gp := range strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator)) {
		goPaths = append(goPaths, []string{gp, "src"})
	}
	goPathsLenIs1 = (len(goPaths) == 1)
}

func dirExists(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return stat.IsDir()
	}
	return false
}

//	Returns the path/filepath.Join()ed full directory path for a specified $GOPATH/src sub-directory.
//	Example: util.GopathSrc("tools", "importers", "sql") = "c:\gd\src\tools\importers\sql" if $GOPATH is c:\gd.
func GopathSrc(subDirNames ...string) (gps string) {
	var gp []string
	for _, gp = range goPaths {
		if gps = filepath.Join(append(gp, subDirNames...)...); goPathsLenIs1 || dirExists(gps) {
			break
		}
	}
	return
}

//	Returns the path/filepath.Join()ed full directory path for a specified $GOPATH/src/github.com sub-directory.
//	Example: util.GopathSrcGithub("metaleap", "go-util", "num") = "c:\gd\src\github.com\metaleap\go-util\num" if $GOPATH is c:\gd.
func GopathSrcGithub(gitHubName string, subDirNames ...string) string {
	return GopathSrc(append([]string{"github.com", gitHubName}, subDirNames...)...)
}

//	If err isn't nil, short-hand for log.Println(err.Error())
func LogError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

//	Returns the path to the current user's home directory.
//	Might be C:\Users\Kitty under Windows, /home/Kitty under Linux or /Users/Kitty under OSX.
//	Specifically, returns the value of either the $userprofile or the $HOME environment variable, whichever one is set.
func UserHomeDirPath() (dirPath string) {
	if dirPath = os.ExpandEnv("$userprofile"); len(dirPath) == 0 {
		dirPath = os.ExpandEnv("$HOME")
	}
	return
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifb(cond, ifTrue, ifFalse bool) bool {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifd(cond bool, ifTrue, ifFalse float64) float64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifi(cond bool, ifTrue, ifFalse int) int {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifi16(cond bool, ifTrue, ifFalse int16) int16 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifi32(cond bool, ifTrue, ifFalse int32) int32 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifi64(cond bool, ifTrue, ifFalse int64) int64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifs(cond bool, ifTrue string, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifu32(cond bool, ifTrue, ifFalse uint32) uint32 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifu64(cond bool, ifTrue, ifFalse uint64) uint64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}
