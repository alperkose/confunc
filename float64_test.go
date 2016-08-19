package confunc_test

import (
	"github.com/alperkose/confunc"
	"strconv"
	"testing"
)

func Test_Float64Func(t *testing.T) {

	expectedValue := 42.42
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: strconv.FormatFloat(expectedValue, 'f', -1, 64)})).
		Float64(configurationKey)

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}

func Test_Float64FuncWithInterceptor(t *testing.T) {

	aValue := 42.42
	expectedValue := 15.96
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: strconv.FormatFloat(aValue, 'f', -1, 64)})).
		Float64(configurationKey, func(s confunc.Confunc) confunc.Confunc {
			return func() (string, error) { return strconv.FormatFloat(expectedValue, 'f', -1, 64), nil }
		})

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}
