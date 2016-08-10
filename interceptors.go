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

func CacheFor(cacheDuration time.Duration) Interceptor {
	var lastCall time.Time = time.Now().Add(-10*time.Second - cacheDuration)
	var m sync.Mutex
	var cached string

	return func(v String) string {
		m.Lock()
		if time.Since(lastCall) > cacheDuration {
			lastCall = time.Now()
			cached = v()
		}
		m.Unlock()
		return cached
	}
}

func Default(defaultVal string) Interceptor {
	return func(v String) string {
		actualVal := v()
		if len(actualVal) == 0 {
			return defaultVal
		}
		return actualVal
	}
}
