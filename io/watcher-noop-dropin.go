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
func (me *Watcher) Close() (err error) {
	return
}

//	No-op
func (me *Watcher) Go() {
}

//	If runHandlerNow is true, runs handler for all dirs/files in dirPath that match namePattern.
func (me *Watcher) WatchIn(dirPath string, namePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error) {
	if runHandlerNow {
		errs = watchRunHandler(filepath.Clean(dirPath), namePattern, handler)
	}
	return
}
