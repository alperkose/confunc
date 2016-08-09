package confunc_test

import (
	"github.com/alperkose/confunc"
	"testing"
)

func Test_FunctionGenerator(t *testing.T) {

	expectedValue := "some param"
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: expectedValue})).
		String(configurationKey)

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}

func Test_FunctionGenerator_WhenAnInterceptorIsProvided(t *testing.T) {
	configurationValue := "some param"
	postfix := "Wrapped"
	expectedValue := configurationValue + postfix
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: configurationValue})).
		String(configurationKey, func(v confunc.String) string {
			return v() + postfix
		})

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}
