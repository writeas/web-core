package memo

import (
	"sync"
	"time"
)

type (
	Memo struct {
		f     Func
		mu    sync.Mutex //guards cache
		cache *entry

		cacheDur  time.Duration
		lastCache time.Time
	}

	Func func() (interface{}, error)
)

type (
	entry struct {
		res   result
		ready chan struct{} // close when res is read
	}

	result struct {
		value interface{}
		err   error
	}
)

func New(f Func, cd time.Duration) *Memo {
	return &Memo{
		f:        f,
		cacheDur: cd,
	}
}

// Invalidate resets the cache to nil if the memo's given cache duration has
// elapsed, and returns true if the cache was actually invalidated.
func (memo *Memo) Invalidate() bool {
	defer memo.mu.Unlock()
	memo.mu.Lock()

	if memo.cache != nil && time.Now().Sub(memo.lastCache) > memo.cacheDur {
		memo.cache = nil
		memo.lastCache = time.Now()
		return true
	}
	return false
}

func (memo *Memo) Get() (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f()

		close(e.ready)
	} else {
		memo.mu.Unlock()

		<-e.ready
	}
	return e.res.value, e.res.err
}

// Reset forcibly resets the cache to nil without checking if the cache
// duration has elapsed.
func (memo *Memo) Reset() {
	defer memo.mu.Unlock()
	memo.mu.Lock()

	memo.cache = nil
	memo.lastCache = time.Now()
}
