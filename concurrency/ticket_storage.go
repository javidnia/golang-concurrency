package concurrency

import (
	"runtime"
	"sync/atomic"
)

type TicketStorage struct {
	ticket *uint64
	done   *uint64
	slots  []string
}

func (ts *TicketStorage) Put(s string) {
	t := atomic.AddUint64(ts.ticket, 1) - 1             // draw a ticket
	ts.slots[t] = s                                     // store your data
	for !atomic.CompareAndSwapUint64(ts.done, t, t+1) { // increase done, preventing race condition
		runtime.Gosched()
	}
}

func (ts *TicketStorage) GetDone() []string {
	return ts.slots[:atomic.LoadUint64(ts.done)+1] // this is wait-free
}
