package confunc_test

import (
	"github.com/alperkose/confunc"
	"os"
	"testing"
)

func Test_EnvSource(t *testing.T) {
	expectedValue := "TEST_VALUE"
	configurationKey := "TEST_PARAM"
	os.Setenv(configurationKey, expectedValue)

	var sut confunc.Source

	sut = confunc.Env()
	actualValue, err := sut.Value(configurationKey)

	if err != nil {
		t.Errorf("error should not have occurred : %v", err.Error())
	}
	if actualValue != expectedValue {
		t.Errorf("expected %v to be %v", actualValue, expectedValue)
	}

}
