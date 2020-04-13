package concurrency

import (
	"runtime"
	"sync/atomic"
)

// use in place of mutex
type SpinLock struct {
	state *int32
}

const free = int32(0)

func (l *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(l.state, free, 42) {
		runtime.Gosched()
	}
}

func (l *SpinLock) Unlock() {
	atomic.StoreInt32(l.state, free)
}
