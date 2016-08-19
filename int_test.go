package confunc_test

import (
	"github.com/alperkose/confunc"
	"strconv"
	"testing"
)

func Test_IntFunc(t *testing.T) {

	expectedValue := 17
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: strconv.Itoa(expectedValue)})).
		Int(configurationKey)

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}

func Test_IntFuncWithInterceptor(t *testing.T) {

	aValue := 17
	expectedValue := 25
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: strconv.Itoa(aValue)})).
		Int(configurationKey, func(s confunc.Confunc) confunc.Confunc {
			return func() (string, error) { return strconv.Itoa(expectedValue), nil }
		})

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}
