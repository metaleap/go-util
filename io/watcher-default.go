// +build !appengine

package uio

import (
	"path/filepath"
	"time"

	"github.com/goforks/fsnotify"

	ustr "github.com/metaleap/go-util/str"
)

//	A convenience wrapper around fsnotify.Watcher.
//	Usage: var w uio.Watcher; w.WatchDir(..); w.WatchFiles(..); go w.Go()
type Watcher struct {
	//	Embedded fsnotify.Watcher
	*fsnotify.Watcher

	//	Defaults to a time.Duration of 250 milliseconds
	DebounceNano int64

	//	A collection of custom fsnotify.FileEvent handlers.
	//	Not related to the handlers specified via the WatchDir(..) and WatchFiles(..) methods.
	OnEvent []func(evt *fsnotify.FileEvent)

	//	A collection of custom error handlers.
	OnError []func(err error)

	dirHandlers map[string][]WatcherHandler

	fileHandlers map[string][]WatcherHandler
}

//	Always returns a new Watcher, even if err is not nil.
func NewWatcher() (me *Watcher, err error) {
	me = &Watcher{dirHandlers: map[string][]WatcherHandler{}, fileHandlers: map[string][]WatcherHandler{}}
	me.DebounceNano = time.Duration(250 * time.Millisecond).Nanoseconds()
	me.Watcher, err = fsnotify.NewWatcher()
	return
}

//	Starts watching. A never-ending loop designed to be called in a new go-routine.
func (me *Watcher) Go() {
	var (
		evt      *fsnotify.FileEvent
		err      error
		hasLast  bool
		now, dif int64
	)
	lastEvt := map[string]int64{}
	for {
		select {
		case evt = <-me.Event:
			if evt != nil {
				_, hasLast = lastEvt[evt.Name]
				now = time.Now().UnixNano()
				if dif = now - lastEvt[evt.Name]; dif > me.DebounceNano || !hasLast {
					lastEvt[evt.Name] = now
					for _, on := range me.OnEvent {
						on(evt)
					}
					dirPath := filepath.Dir(evt.Name)
					for _, on := range me.dirHandlers[dirPath] {
						on(dirPath)
					}
					for filePathPattern, handlers := range me.fileHandlers {
						if filepath.Dir(filePathPattern) == dirPath && ustr.MatchesAny(filepath.Base(evt.Name), filepath.Base(filePathPattern)) {
							for _, on := range handlers {
								on(evt.Name)
							}
						}
					}
				}
			}
		case err = <-me.Error:
			if err != nil {
				for _, on := range me.OnError {
					on(err)
				}
			}
		}
	}
}

//	Watches the specified directory (but not sub-directories) for change events.
//
//	handler is invoked whenever a change event is observed.
//
//	runHandlerNow allows immediate one-time invokation of handler.
//	This is for the use-case pattern "load this dir now, then reload whenever it is modified"
func (me *Watcher) WatchDir(dirPath string, runHandlerNow bool, handler WatcherHandler) (err error) {
	dirPath = filepath.Clean(dirPath)
	all, ok := me.dirHandlers[dirPath]
	if !ok {
		err = me.Watch(dirPath)
	}
	if err == nil {
		me.dirHandlers[dirPath] = append(all, handler)
		if runHandlerNow {
			handler(dirPath)
		}
	}
	return
}

//	Watches files (whose name matches the specified pattern) in the specified directory for change events.
//
//	handler is invoked whenever a change event is observed, providing the full file path.
//
//	runHandlerNow allows immediate one-time invokation of handler.
//	This is for the use-case pattern "load those files now, then reload whenever they are modified"
func (me *Watcher) WatchFiles(dirPath string, fileNamePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error) {
	dirPath = filepath.Clean(dirPath)
	filePath := filepath.Join(dirPath, string(fileNamePattern))
	if _, ok := me.dirHandlers[dirPath]; !ok {
		if err := me.Watch(dirPath); err != nil {
			errs = append(errs, err)
		} else {
			me.dirHandlers[dirPath] = []WatcherHandler{}
		}
	}
	if len(errs) == 0 {
		me.fileHandlers[filePath] = append(me.fileHandlers[filePath], handler)
		if runHandlerNow {
			errs = append(errs, watchFilesRunHandler(dirPath, fileNamePattern, handler)...)
		}
	}
	return
}
