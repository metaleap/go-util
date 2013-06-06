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
	if runHandlerNow {
		watchFilesRunHandler(filepath.Clean(dirPath), fileNamePattern, handler)
	}
}
