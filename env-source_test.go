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
	actualValue := sut.Value(configurationKey)
	if actualValue != expectedValue {
		t.Errorf("expected %v to be %v", actualValue, expectedValue)
	}
}
