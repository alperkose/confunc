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

func Test_IntFuncWithWrapper(t *testing.T) {

	aValue := 17
	expectedValue := 25
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: strconv.Itoa(aValue)})).
		Int(configurationKey, func(s confunc.String) string {
			return strconv.Itoa(expectedValue)
		})

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}
