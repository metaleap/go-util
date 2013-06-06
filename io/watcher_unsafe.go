// +build !appengine

package uio

import (
	"path/filepath"

	"github.com/goforks/fsnotify"

	ustr "github.com/metaleap/go-util/str"
)

type Watcher struct {
	*fsnotify.Watcher
	OnEvent      []func(evt *fsnotify.FileEvent)
	OnError      []func(err error)
	DirHandlers  map[string][]func()
	FileHandlers map[string][]func(filePath string)
}

func NewWatcher() (me *Watcher, err error) {
	me = &Watcher{DirHandlers: map[string][]func(){}, FileHandlers: map[string][]func(string){}}
	me.Watcher, err = fsnotify.NewWatcher()
	return
}

func (me *Watcher) Close() error {
	return me.Watcher.Close()
}

func (me *Watcher) Go() {
	var (
		evt *fsnotify.FileEvent
		err error
	)
	for {
		select {
		case evt = <-me.Event:
			if evt != nil {
				for _, on := range me.OnEvent {
					on(evt)
				}
				dirPath := filepath.Dir(evt.Name)
				for _, on := range me.DirHandlers[dirPath] {
					on()
				}
				for filePathPattern, handlers := range me.FileHandlers {
					if filepath.Dir(evt.Name) == dirPath && ustr.MatchesAny(filepath.Base(evt.Name), filepath.Base(filePathPattern)) {
						for _, on := range handlers {
							on(evt.Name)
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

func (me *Watcher) WatchDir(dirPath string, runHandlerNow bool, handler func()) {
	dirPath = filepath.Clean(dirPath)
	all, ok := me.DirHandlers[dirPath]
	if !ok {
		me.Watch(dirPath)
	}
	me.DirHandlers[dirPath] = append(all, handler)
	if runHandlerNow {
		handler()
	}
}

func (me *Watcher) WatchFiles(dirPath, fileNamePattern string, runHandlerNow bool, handler func(string)) {
	dirPath = filepath.Clean(dirPath)
	filePath := filepath.Join(dirPath, fileNamePattern)
	if _, ok := me.DirHandlers[dirPath]; !ok {
		me.Watch(dirPath)
		me.DirHandlers[dirPath] = []func(){}
	}
	me.FileHandlers[filePath] = append(me.FileHandlers[filePath], handler)
	if runHandlerNow {
		watchFilesRunHandler(dirPath, fileNamePattern, handler)
	}
}
