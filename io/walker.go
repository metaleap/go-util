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

	//	If false, only the files in the specified directory
	//	(and the directory itself) get visited, but no sub-directories.
	VisitSubDirs bool

	//	Called for every directory being visited during Walk().
	DirVisitor WalkerVisitor

	//	Called for every file being visited during Walk().
	FileVisitor WalkerVisitor
}

//	Initializes and returns a new DirWalker with the specified (optional) visitors.
//	The deep argument sets the VisitSubDirs field.
func NewDirWalker(deep bool, dirVisitor, fileVisitor WalkerVisitor) (me *DirWalker) {
	me = &DirWalker{DirVisitor: dirVisitor, FileVisitor: fileVisitor, VisitSubDirs: deep}
	return
}

//	Initiates me walking through the specified directory.
func (me *DirWalker) Walk(dirPath string) (errs []error) {
	me.walk(true, dirPath, &errs)
	return
}

func (me *DirWalker) walk(walkSelf bool, dirPath string, errs *[]error) {
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
	if walkSelf {
		walkSelf = dirVisitor(me, dirPath, dirInfo)
	} else {
		walkSelf = true
	}
	if walkSelf {
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
	keepWalking = true
	for _, fi := range fileInfos {
		if fullPath = filepath.Join(dirPath, fi.Name()); fi.IsDir() == isDir {
			if keepWalking = visitor(me, fullPath, fi); !keepWalking {
				break
			} else if isDir && me.VisitSubDirs {
				me.walk(false, fullPath, errs)
			}
		}
	}
	return
}
