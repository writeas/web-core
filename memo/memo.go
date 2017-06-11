package memo

import (
	"sync"
)

type (
	Memo struct {
		f     Func
		mu    sync.Mutex //guards cache
		cache *entry
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

func New(f Func) *Memo {
	return &Memo{
		f: f,
	}
}

func (memo *Memo) Invalidate() {
	memo.mu.Lock()
	memo.cache = nil
	memo.mu.Unlock()
}

func (memo *Memo) Get() (value interface{}, err error) {
	memo.mu.Lock()
	if memo.cache == nil {
		memo.cache = &entry{ready: make(chan struct{})}
		memo.mu.Unlock()

		memo.cache.res.value, memo.cache.res.err = memo.f()

		close(memo.cache.ready)
	} else {
		memo.mu.Unlock()

		<-memo.cache.ready
	}
	return memo.cache.res.value, memo.cache.res.err
}
