package confunc_test

import (
	"github.com/alperkose/confunc"
	"strconv"
	"testing"
)

func Test_BoolFunc(t *testing.T) {

	expectedValue := true
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: strconv.FormatBool(expectedValue)})).
		Bool(configurationKey)

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}

func Test_BoolFuncWithInterceptor(t *testing.T) {

	aValue := false
	expectedValue := true
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: strconv.FormatBool(aValue)})).
		Bool(configurationKey, func(s confunc.Confunc) confunc.Confunc {
			return func() (string, error) { return strconv.FormatBool(expectedValue), nil }
		})

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}
