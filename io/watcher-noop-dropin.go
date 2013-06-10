// +build appengine

package uio

import (
	"path/filepath"

	ustr "github.com/metaleap/go-util/str"
)

//	File-watching is not allowed and not necessary on Google App Engine.
//	So this is a "polyfil" empty struct with no-op methods.
type Watcher struct {
}

//	Returns a new Watcher, err is always nil.
func NewWatcher() (me *Watcher, err error) {
	me = &Watcher{}
	return
}

//	No-op
func (me *Watcher) Close() error {
	return nil
}

//	No-op
func (me *Watcher) Go() {
}

//	Runs handler if runHandlerNow is true, always returns nil.
func (me *Watcher) WatchDir(dirPath string, runHandlerNow bool, handler WatcherHandler) (err error) {
	if runHandlerNow {
		handler(dirPath)
	}
	return
}

//	If runHandlerNow is true, runs handler for all files in dirPath that match fileNamePattern.
func (me *Watcher) WatchFiles(dirPath string, fileNamePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error) {
	if runHandlerNow {
		errs = watchFilesRunHandler(filepath.Clean(dirPath), fileNamePattern, handler)
	}
	return
}
