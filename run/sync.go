package urun

import (
	"sync"
)

//	A `sync.Mutex` wrapper for convenient conditional `defer`d un/locking.
//
//	Example: `defer mut.UnlockIf(mut.LockIf(mycondition))`
type MutexIf struct {
	sync.Mutex
}

func (me *MutexIf) Lock() bool {
	me.Mutex.Lock()
	return true
}

//	Calls `me.Lock` if `lock` is `true`, then returns `lock`.
func (me *MutexIf) LockIf(lock bool) bool {
	if lock {
		me.Mutex.Lock()
	}
	return lock
}

//	Calls `me.Unlock` if `unlock` is `true`.
func (me *MutexIf) UnlockIf(unlock bool) {
	if unlock {
		me.Mutex.Unlock()
	}
}

func WaitOn(funcs ...func()) {
	if l := len(funcs); l == 0 {
		return
	} else if l == 1 {
		funcs[0]()
		return
	}
	var wait sync.WaitGroup
	run := func(fn func()) {
		defer wait.Done()
		fn()
	}
	for _, fn := range funcs {
		wait.Add(1)
		go run(fn)
	}
	wait.Wait()
}

func WaitOn_(funcs ...func()) {
	for _, fn := range funcs {
		fn()
	}
}
