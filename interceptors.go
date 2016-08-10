package confunc

import (
	"sync"
	"time"
)

func CacheOnce() Interceptor {
	var o sync.Once
	var cached string

	return func(v String) string {
		o.Do(func() {
			cached = v()
		})
		return cached
	}
}

func CacheFor(d time.Duration) Interceptor {
	var lastCall time.Time = time.Now().Add(-10*time.Second - d)
	var m sync.Mutex
	var cached string

	return func(v String) string {
		m.Lock()
		if time.Since(lastCall) > d {
			lastCall = time.Now()
			cached = v()
		}
		m.Unlock()
		return cached
	}
}
