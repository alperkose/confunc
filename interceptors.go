package confunc

import (
	"sync"
	"time"
)

func CacheOnce() Interceptor {
	var m sync.Mutex
	var cached Confunc
	var retry = true

	return func(v Confunc) Confunc {
		m.Lock()
		if retry {
			cachedVal, err := v()
			if err != nil {
				retry = true
				cached = func() (string, error) { return "", err }
			} else {
				retry = false
				cached = func() (string, error) { return cachedVal, nil }
			}
		}
		m.Unlock()
		return cached
	}
}

func CacheFor(cacheDuration time.Duration) Interceptor {
	var lastCall time.Time = time.Now().Add(-10*time.Second - cacheDuration)
	var m sync.Mutex
	var cached Confunc
	var retry = true

	return func(v Confunc) Confunc {

		m.Lock()
		if retry || time.Since(lastCall) > cacheDuration {
			lastCall = time.Now()
			cachedVal, err := v()
			if err != nil {
				retry = true
				cached = func() (string, error) { return "", err }
			} else {
				retry = false
				cached = func() (string, error) { return cachedVal, nil }
			}
		}
		m.Unlock()
		return cached
	}
}

func Default(defaultVal string) Interceptor {
	return func(v Confunc) Confunc {
		actualVal, err := v()
		if err != nil {
			actualVal = defaultVal
		}
		return func() (string, error) { return actualVal, nil }
	}
}
