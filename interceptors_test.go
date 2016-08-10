package confunc_test

import (
	"github.com/alperkose/confunc"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_CacheOnceInterceptor(t *testing.T) {
	interceptorUnderTest := confunc.CacheOnceInterceptor()

	firstVal := interceptorUnderTest(someRandomStuff)
	for i := 0; i < 100; i++ {
		aVal := interceptorUnderTest(someRandomStuff)
		if firstVal != aVal {
			t.Errorf("expected '%v' to be '%v'", firstVal, aVal)
		}
	}
}

var rndSource = rand.New(rand.NewSource(time.Now().Unix()))

func someRandomStuff() string {
	return strconv.FormatFloat(rndSource.Float64(), 'f', 10, 64)
}
