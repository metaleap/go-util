// +build appengine

package uio

import (
	"os"
	"path/filepath"

	ustr "github.com/metaleap/go-util/str"
)

type Watcher struct {
}

func NewWatcher() (me *Watcher, err error) {
	me = &Watcher{}
	return
}

func (me *Watcher) Close() error {
	return nil
}

func (me *Watcher) Go() {
}

func (me *Watcher) WatchDir(dirPath string, runHandlerNow bool, handler func()) {
	if runHandlerNow {
		handler()
	}
}

func (me *Watcher) WatchFiles(dirPath, fileNamePattern string, runHandlerNow bool, handler func(string)) {
	dirPath = filepath.Clean(dirPath)
	if runHandlerNow {
		var m ustr.Matcher
		m.AddPatterns(fileNamePattern)
		NewDirWalker(false, nil, func(_ *DirWalker, fullPath string, _ os.FileInfo) bool {
			if m.IsMatch(filepath.Base(fullPath)) {
				handler(fullPath)
			}
			return true
		}).Walk(dirPath)
	}
}
