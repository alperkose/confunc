package confunc_test

import (
	"github.com/alperkose/confunc"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_CacheOnceInterceptor(t *testing.T) {
	interceptorUnderTest := confunc.CacheOnce()

	firstVal := interceptorUnderTest(someRandomStuff)
	for i := 0; i < 100; i++ {
		aVal := interceptorUnderTest(someRandomStuff)
		if firstVal != aVal {
			t.Errorf("expected '%v' to be '%v'", firstVal, aVal)
		}
	}
}

func Test_CacheForInterceptor(t *testing.T) {
	var cacheOffset = 250 * time.Millisecond
	interceptorUnderTest := confunc.CacheFor(cacheOffset)

	firstVal := interceptorUnderTest(someRandomStuff)
	for i := 0; i < 100; i++ {
		aVal := interceptorUnderTest(someRandomStuff)
		if firstVal != aVal {
			t.Errorf("expected '%v' to be '%v'", firstVal, aVal)
		}
	}

	time.Sleep(cacheOffset)

	lastVal := interceptorUnderTest(someRandomStuff)
	if firstVal == lastVal {
		t.Errorf("expected '%v' to be different than '%v'", lastVal, firstVal)
	}
	for i := 0; i < 100; i++ {
		aVal := interceptorUnderTest(someRandomStuff)
		if lastVal != aVal {
			t.Errorf("expected '%v' to be '%v'", firstVal, aVal)
		}
	}
}

var rndSource = rand.New(rand.NewSource(time.Now().Unix()))

func someRandomStuff() string {
	return strconv.FormatFloat(rndSource.Float64(), 'f', 10, 64)
}
