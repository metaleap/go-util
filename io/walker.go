package uio

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type WalkerVisitor func(walker *DirWalker, fullPath string, isDir bool) (keepWalking bool)

type DirWalker struct {
	BreakOnError, VisitDirsFirst, VisitSubDirs bool
	DirVisitor, FileVisitor                    WalkerVisitor
}

func NewDirWalker(dirVisitor, fileVisitor WalkerVisitor) (me *DirWalker) {
	me = &DirWalker{DirVisitor: dirVisitor, FileVisitor: fileVisitor, VisitSubDirs: true}
	if me.DirVisitor == nil {
		me.DirVisitor = func(_ *DirWalker, _ string, _ bool) bool { return true }
	}
	if me.FileVisitor == nil {
		me.FileVisitor = func(_ *DirWalker, _ string, _ bool) bool { return true }
	}
	return
}

func (me *DirWalker) Walk(dirPath string) (errs []error) {
	me.walk(dirPath, &errs)
	return
}

func (me *DirWalker) walk(dirPath string, errs *[]error) {
	if me.DirVisitor(me, dirPath, true) {
		if fileInfos, err := ioutil.ReadDir(dirPath); err == nil {
			if me.VisitDirsFirst {
				if !me.walkInfos(dirPath, fileInfos, true, me.DirVisitor, errs) {
					return
				}
			}
			if !me.walkInfos(dirPath, fileInfos, false, me.FileVisitor, errs) {
				return
			}
			if !me.VisitDirsFirst {
				if !me.walkInfos(dirPath, fileInfos, true, me.DirVisitor, errs) {
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
			if keepWalking = visitor(me, fullPath, isDir); !keepWalking {
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
