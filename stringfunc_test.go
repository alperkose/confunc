package confunc_test

import (
	"testing"
	"github.com/alperkose/confunc"
)

func Test_FunctionGenerator(t *testing.T){

	expectedValue := "some param"
	configurationKey := "myConfig"
	configUnderTest := confunc.From(confunc.Map(map[string]string{configurationKey:expectedValue})).StringFunc(configurationKey).Build()

	actualValue := configUnderTest()
	if actualValue != expectedValue {
		t.Errorf("expected %v to be %v", actualValue, expectedValue)
	}
}
