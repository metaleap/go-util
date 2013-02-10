package uio

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//	Used for DirWalker.DirVisitor and DirWalker.FileVisitor. Always check fileInfo for nil!
//	Always return keepWalking as true unless you want to immediately terminate a Walk() early.
type WalkerVisitor func(walker *DirWalker, fullPath string, fileInfo os.FileInfo) (keepWalking bool)

//	An empty WalkerVisitor used in place of a nil DirWalker.DirVisitor or a nil DirWalker.FileVisitor during a DirWalker.Walk(). Returns true.
func WalkerVisitorNoop(_ *DirWalker, _ string, _ os.FileInfo) bool {
	return true
}

//	Returns a new WalkerVisitor that during a DirWalker.Walk() tracks FileInfo.ModTime() for all visited files
//	and/or directories, always storing the newest one in fileTime, and terminating early as soon as fileTime
//	records a value higher than the specified testTime.
func NewWalkerVisitor_IsNewerThan(testTime time.Time, fileTime *time.Time) WalkerVisitor {
	var tmpTime time.Time
	return func(_ *DirWalker, _ string, fileInfo os.FileInfo) bool {
		if fileInfo != nil {
			if tmpTime = fileInfo.ModTime(); tmpTime.UnixNano() > fileTime.UnixNano() {
				*fileTime = tmpTime
			}
		}
		return fileTime.UnixNano() <= testTime.UnixNano()
	}
}

//	Provides recursive directory walking with a variety of options.
type DirWalker struct {
	//	Walk() returns a slice of all errors encountered but
	//	to cancel walking upon the first error, set this to true.
	BreakOnError bool

	//	After invoking DirVisitor on the specified directory, by default
	//	its files get visited first before visiting its sub-directories.
	//	If VisitDirsFirst is true, then files get visited last, after
	//	having visited all sub-directories.
	VisitDirsFirst bool

	//	Defaults to true. If false, only the files in the specified directory
	//	(and the directory itself) get visited, but no sub-directories.
	VisitSubDirs bool

	//	Called for every directory being visited during Walk().
	DirVisitor WalkerVisitor

	//	Called for every file being visited during Walk().
	FileVisitor WalkerVisitor
}

//	Initializes and returns a new DirWalker with the specified (optional) visitors.
func NewDirWalker(dirVisitor, fileVisitor WalkerVisitor) (me *DirWalker) {
	me = &DirWalker{DirVisitor: dirVisitor, FileVisitor: fileVisitor, VisitSubDirs: true}
	return
}

//	Initiates me walking through the specified directory.
func (me *DirWalker) Walk(dirPath string) (errs []error) {
	me.walk(dirPath, &errs)
	return
}

func (me *DirWalker) walk(dirPath string, errs *[]error) {
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		if *errs = append(*errs, err); me.BreakOnError {
			return
		}
	}
	dirVisitor, fileVisitor := me.DirVisitor, me.FileVisitor
	if dirVisitor == nil {
		dirVisitor = WalkerVisitorNoop
	}
	if fileVisitor == nil {
		fileVisitor = WalkerVisitorNoop
	}
	if dirVisitor(me, dirPath, dirInfo) {
		if fileInfos, err := ioutil.ReadDir(dirPath); err == nil {
			if me.VisitDirsFirst {
				if !me.walkInfos(dirPath, fileInfos, true, dirVisitor, errs) {
					return
				}
			}
			if !me.walkInfos(dirPath, fileInfos, false, fileVisitor, errs) {
				return
			}
			if !me.VisitDirsFirst {
				if !me.walkInfos(dirPath, fileInfos, true, dirVisitor, errs) {
					return
				}
			}
		} else if *errs = append(*errs, err); me.BreakOnError {
			return
		}
	}
}

func (me *DirWalker) walkInfos(dirPath string, fileInfos []os.FileInfo, isDir bool, visitor WalkerVisitor, errs *[]error) (keepWalking bool) {
	var fullPath string
	for _, fi := range fileInfos {
		if fullPath = filepath.Join(dirPath, fi.Name()); fi.IsDir() == isDir {
			if keepWalking = visitor(me, fullPath, fi); !keepWalking {
				break
			} else if isDir && me.VisitSubDirs {
				me.walk(fullPath, errs)
			}
		}
	}
	return
}

/*

#	Recursively walks along a directory hierarchy, calling the specified callback function for each file encountered.
#	dirPath: the path of the directory in which to start walking
#	fileSuffix: optional; if specified, fileFunc is only called for files whose name has this suffix
#	fileFunc: callback function called per file. Returns true to keep recursing into sub-dirs. Arguments: full file path and current recurseSubDirs value
#	recurseSubDirs: true to recurse into sub-directories.
func WalkDirectory(dirPath, fileSuffix string, fileFunc func(string, bool) bool, recurseSubDirs bool) error {
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err == nil {
		for _, fi := range fileInfos {
			if !fi.IsDir() {
				if (len(fileSuffix) == 0) || strings.HasSuffix(fi.Name(), fileSuffix) {
					recurseSubDirs = fileFunc(filepath.Join(dirPath, fi.Name()), recurseSubDirs)
				}
			}
		}
		if recurseSubDirs {
			for _, fi := range fileInfos {
				if fi.IsDir() {
					if err = WalkDirectory(filepath.Join(dirPath, fi.Name()), fileSuffix, fileFunc, recurseSubDirs); err != nil {
						break
					}
				}
			}
		}
	}
	return err
}

*/
