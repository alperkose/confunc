package confunc_test

import (
	"errors"
	"github.com/alperkose/confunc"
	"testing"
)

func Test_StringFunc(t *testing.T) {

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

func Test_StringFunc_WhenAnInterceptorIsProvided(t *testing.T) {
	configurationValue := "some param"
	postfix := "Wrapped"
	expectedValue := configurationValue + postfix
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(confunc.Map(map[string]string{configurationKey: configurationValue})).
		String(configurationKey, func(v confunc.Confunc) confunc.Confunc {
			return func() (string, error) {
				val, _ := v()
				return val + postfix, nil
			}
		})

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}

func Test_StringFunc_WhenSourceReturnedAnError(t *testing.T) {
	expectedValue := ""
	configurationKey := "myConfig"
	configUnderTest := confunc.
		From(&errorSource{}).
		String(configurationKey)

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected '%v' to be '%v'", actualValue, expectedValue)
	}
}

type errorSource struct{}

func (s *errorSource) Value(k string) (string, error) {
	return "", errors.New("Yassak")
}
