// +build appengine

package uio

import (
	"path/filepath"

	ustr "github.com/metaleap/go-util/str"
)

//	A convenient wrapper around `goforks/fsnotify.Watcher`.
//
//	Usage:
//		var w uio.Watcher
//		w.WatchIn(dir, pattern, runNow, handler)
//		go w.Go()
//		otherCode(laterOn...)
//		w.WatchIn(anotherDir...)
type Watcher struct {
}

//	Returns a new `Watcher`, `err` is always nil.
func NewWatcher() (me *Watcher, err error) {
	me = &Watcher{}
	return
}

//	Closes the underlying `me.Watcher`.
func (me *Watcher) Close() (err error) {
	return
}

func (me *Watcher) Go() {
}

func (me *Watcher) WatchIn(dirPath string, namePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error) {
	if runHandlerNow {
		errs = watchRunHandler(filepath.Clean(dirPath), namePattern, handler)
	}
	return
}
