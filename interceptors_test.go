package confunc_test

import (
	"errors"
	"github.com/alperkose/confunc"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_CacheOnceInterceptor(t *testing.T) {
	interceptorUnderTest := confunc.CacheOnce()
	cfn := interceptorUnderTest(someRandomStuff)

	firstVal, err := cfn()
	if err != nil {
		t.Errorf("error should not have occurred : %v ", err)
	}

	for i := 0; i < 100; i++ {
		cfn = interceptorUnderTest(someRandomStuff)
		aVal, err := cfn()

		if err != nil {
			t.Errorf("error should not have occurred : %v ", err)
		}
		if firstVal != aVal {
			t.Errorf("expected '%v' to be '%v'", firstVal, aVal)
		}
	}
}

func Test_CacheOnce_WhenSourceHasTemporaryError(t *testing.T) {
	interceptorUnderTest := confunc.CacheOnce()
	testCacheInterceptor_WhenSourceHasTemporaryError(t, interceptorUnderTest)
}

func Test_CacheForInterceptor(t *testing.T) {
	var cacheOffset = 250 * time.Millisecond
	interceptorUnderTest := confunc.CacheFor(cacheOffset)

	cfn := interceptorUnderTest(someRandomStuff)
	firstVal, err := cfn()
	if err != nil {
		t.Errorf("error should not have occurred : %v ", err)
	}

	for i := 0; i < 100; i++ {
		cfn = interceptorUnderTest(someRandomStuff)
		aVal, err := cfn()

		if err != nil {
			t.Errorf("error should not have occurred : %v ", err)
		}
		if firstVal != aVal {
			t.Errorf("expected '%v' to be '%v'", firstVal, aVal)
		}
	}

	time.Sleep(cacheOffset)

	cfn = interceptorUnderTest(someRandomStuff)
	lastVal, err := cfn()

	if err != nil {
		t.Errorf("error should not have occurred : %v ", err)
	}
	if firstVal == lastVal {
		t.Errorf("expected '%v' to be different than '%v'", lastVal, firstVal)
	}
	for i := 0; i < 100; i++ {
		cfn = interceptorUnderTest(someRandomStuff)
		aVal, err := cfn()

		if err != nil {
			t.Errorf("error should not have occurred : %v ", err)
		}
		if lastVal != aVal {
			t.Errorf("expected '%v' to be '%v'", firstVal, aVal)
		}
	}
}

func Test_CacheFor_WhenSourceHasTemporaryError(t *testing.T) {
	interceptorUnderTest := confunc.CacheFor(250 * time.Millisecond)
	testCacheInterceptor_WhenSourceHasTemporaryError(t, interceptorUnderTest)
}

func testCacheInterceptor_WhenSourceHasTemporaryError(t *testing.T, interceptorUnderTest confunc.Interceptor) {
	confuncToCover := confuncWithErrorAndSuccessCycle()
	cfn := interceptorUnderTest(confuncToCover)

	_, err := cfn()
	if err == nil {
		t.Errorf("an error should have occurred")
	}

	cfn = interceptorUnderTest(confuncToCover)
	secondVal, err := cfn()
	if err != nil {
		t.Errorf("error should not have occurred : %v ", err)
	}
	if len(secondVal) == 0 {
		t.Errorf("should have returned a value")
	}

	cfn = interceptorUnderTest(confuncToCover)
	thirdVal, err := cfn()
	if err != nil {
		t.Errorf("error should not have occurred : %v ", err)
	}
	if thirdVal != secondVal {
		t.Errorf("should have returned value from cache")
	}
}

func Test_DefaultInterceptor_WhenSourceHasNoValue_ShouldReturnDefaultValue(t *testing.T) {
	defaultValue := "42"
	interceptorUnderTest := confunc.Default(defaultValue)

	cfn := interceptorUnderTest(func() (string, error) { return "", errors.New("Configuration Not found") })
	actualVal, err := cfn()

	if err != nil {
		t.Errorf("error should not have occurred : %v ", err)
	}

	if actualVal != defaultValue {
		t.Errorf("expected '%v' to be '%v'", actualVal, defaultValue)
	}
}

func Test_DefaultInterceptor_WhenSourceHasValue_ShouldReturnSourceValue(t *testing.T) {
	sourceValue := "7"
	defaultValue := "42"
	interceptorUnderTest := confunc.Default(defaultValue)

	cfn := interceptorUnderTest(func() (string, error) { return sourceValue, nil })
	actualVal, err := cfn()

	if err != nil {
		t.Errorf("error should not have occurred : %v ", err)
	}

	if actualVal != sourceValue {
		t.Errorf("expected '%v' to be '%v'", actualVal, sourceValue)
	}
}

var rndSource = rand.New(rand.NewSource(time.Now().Unix()))

func someRandomStuff() (string, error) {
	return strconv.FormatFloat(rndSource.Float64(), 'f', 10, 64), nil
}

func confuncWithErrorAndSuccessCycle() confunc.Confunc {
	var hammerTime = true
	return func() (string, error) {
		v := ""
		var err error
		err = nil
		if hammerTime {
			err = errors.New("Can't touch this!!!")
		} else {
			v, err = someRandomStuff()
		}
		hammerTime = !hammerTime
		return v, err
	}
}

type testSource struct{}

func (ts *testSource) Value(k string) (string, error) {
	v, err := someRandomStuff()

	return v, err
}

func BenchmarkAll(b *testing.B) {
	confuncUnderTest := confunc.From(&testSource{}).Float64("somekey", confunc.CacheOnce())

	firstVal := confuncUnderTest()
	for i := 0; i < b.N; i++ {
		aVal := confuncUnderTest()
		if firstVal != aVal {
			b.Errorf("expected '%v' to be '%v'", firstVal, aVal)
		}
	}

}
