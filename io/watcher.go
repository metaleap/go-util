package uio

import (
	"path/filepath"

	"github.com/goforks/fsnotify"
)

type WatchEventHandler func(evt *fsnotify.FileEvent)

type Watcher struct {
	*fsnotify.Watcher
	OnEvent      []WatchEventHandler
	OnError      []func(err error)
	DirHandlers  map[string][]WatchEventHandler
	FileHandlers map[string][]WatchEventHandler
}

func NewWatcher() (me *Watcher, err error) {
	me = &Watcher{DirHandlers: map[string][]WatchEventHandler{}, FileHandlers: map[string][]WatchEventHandler{}}
	if me.Watcher, err = fsnotify.NewWatcher(); err != nil {
		me = nil
	}
	return
}

func (me *Watcher) Go() {
	onAll := func(all []WatchEventHandler, evt *fsnotify.FileEvent) {
		for _, on := range all {
			on(evt)
		}
	}
	var (
		evt *fsnotify.FileEvent
		err error
	)
	for {
		select {
		case evt = <-me.Event:
			if evt != nil {
				onAll(me.OnEvent, evt)
				onAll(me.DirHandlers[filepath.Dir(evt.Name)], evt)
				onAll(me.FileHandlers[evt.Name], evt)
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

func (me *Watcher) WatchDir(dirPath string, runHandlerNow bool, handler WatchEventHandler) {
	all, ok := me.DirHandlers[dirPath]
	if !ok {
		me.Watch(dirPath)
	}
	me.DirHandlers[dirPath] = append(all, handler)
	if runHandlerNow {
		handler(nil)
	}
}

func (me *Watcher) WatchFile(filePath string, runHandlerNow bool, handler WatchEventHandler) {
	dirPath := filepath.Dir(filePath)
	if _, ok := me.DirHandlers[dirPath]; !ok {
		me.Watch(dirPath)
		me.DirHandlers[dirPath] = []WatchEventHandler{}
	}
	me.FileHandlers[filePath] = append(me.FileHandlers[filePath], handler)
	if runHandlerNow {
		handler(nil)
	}
}
