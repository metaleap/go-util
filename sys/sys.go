package usys

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/metaleap/go-util/fs"
)

var (
	_userHomeDirPath  string
	_userDataDirPaths = make(map[bool]string, 2)

	//	Look-up hash-table for the `OSName` function.
	OSNames = map[string]string{
		"windows":   "Windows",
		"darwin":    "Mac OS X",
		"linux":     "Linux",
		"freebsd":   "FreeBSD",
		"appengine": "Google App Engine",
	}
)

//	Short-hand for: `runtime.GOMAXPROCS(2 * runtime.NumCPU())`.
func MaxProcs() {
	runtime.GOMAXPROCS(2 * runtime.NumCPU())
}

//	Returns the human-readable operating system name represented by the specified
//	`goOS` name, by looking up the corresponding entry in `OSNames`.
func OSName(goOS string) (name string) {
	if name = OSNames[goOS]; len(name) == 0 {
		name = strings.ToTitle(goOS)
	}
	return
}

func UserDataDirPath(preferCacheOverConfig bool) string {
	dirpath := _userDataDirPaths[preferCacheOverConfig]
	if len(dirpath) == 0 {
		probeenvvars := []string{"XDG_CONFIG_HOME", "XDG_CACHE_HOME", "LOCALAPPDATA", "APPDATA"}
		if preferCacheOverConfig {
			probeenvvars[0], probeenvvars[1] = probeenvvars[1], probeenvvars[0]
		}
		for _, envvar := range probeenvvars {
			if maybedirpath := os.Getenv(envvar); len(maybedirpath) > 0 && ufs.DirExists(maybedirpath) {
				dirpath = maybedirpath
				break
			}
		}
		if len(dirpath) == 0 {
			probehomesubdirs := []string{".config", ".cache", "Library/Caches", "Library/Application Support"}
			if preferCacheOverConfig {
				probehomesubdirs[0], probehomesubdirs[1] = probehomesubdirs[1], probehomesubdirs[0]
			}
			for _, homesubdir := range probehomesubdirs {
				if maybedirpath := filepath.Join(UserHomeDirPath(), homesubdir); ufs.DirExists(maybedirpath) {
					dirpath = maybedirpath
					break
				}
			}
			if len(dirpath) == 0 {
				dirpath = UserHomeDirPath()
			}
		}
		_userDataDirPaths[preferCacheOverConfig] = dirpath
	}
	return dirpath
}

//	Returns the path to the current user's home directory.
func UserHomeDirPath() string {
	dirpath := _userHomeDirPath
	if len(dirpath) == 0 {
		if user, err := user.Current(); err == nil && len(user.HomeDir) > 0 && ufs.DirExists(user.HomeDir) {
			dirpath = user.HomeDir
		} else if dirpath = os.Getenv("USERPROFILE"); len(dirpath) == 0 || !ufs.DirExists(dirpath) {
			dirpath = os.Getenv("HOME")
		}
		_userHomeDirPath = dirpath
	}
	return dirpath
}
