package util

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//	Returns the *filepath*-joined full directory path for a specified $GOPATH/src/github.com sub-directory.
//	Example: util.BaseCodePathGo("metaleap", "go-util") = "c:\gd\src\github.com\metaleap\go-util" if $GOPATH is c:\gd.
func BaseCodePathGithub(gitHubName string, subDirNames ...string) string {
	return filepath.Join(append([]string{os.Getenv("GOPATH"), "src", "github.com", gitHubName}, subDirNames...)...)
}

//	Returns the *filepath*-joined full directory path for a specified $GOPATH/src sub-directory.
//	Example: util.BaseCodePathGo("tools", "importers", "sql") = "c:\gd\src\tools\importers\sql" if $GOPATH is c:\gd.
func BaseCodePathGo(subDirNames ...string) string {
	return filepath.Join(append([]string{os.Getenv("GOPATH"), "src"}, subDirNames...)...)
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

//	Attempts to extract major and minor version components from a string that begins with a version number.
//	Example: returns []int{3, 2} and float64(3.2) for a verstr that is "3.2.0 - Build 8.15.10.2761"
func ParseVersion(verstr string) (majorMinor []int, both float64) {
	for _, p := range strings.Split(verstr, ".") {
		if pos := strings.Index(p, " "); pos > 0 {
			p = p[:pos]
		}
		if i, err := strconv.ParseUint(p, 10, 8); err == nil {
			if majorMinor = append(majorMinor, int(i)); len(majorMinor) >= 2 {
				break
			}
		} else {
			break
		}
	}
	if len(majorMinor) > 0 {
		both = float64(majorMinor[0])
	}
	if len(majorMinor) > 1 {
		both += (float64(majorMinor[1]) * 0.1)
	}
	return
}
