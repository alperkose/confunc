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

func Test_BoolFuncWithWrapper(t *testing.T) {

	aValue := false
	expectedValue := true
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: strconv.FormatBool(aValue)})).
		Bool(configurationKey, func(s confunc.String) string {
			return strconv.FormatBool(expectedValue)
		})

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}
