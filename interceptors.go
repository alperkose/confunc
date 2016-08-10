package confunc

import "sync"

func CacheOnceInterceptor() Interceptor {
	var o sync.Once
	var cached string

	return func(v String) string {
		o.Do(func() {
			cached = v()
		})
		return cached
	}
}
