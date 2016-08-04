package confunc_test

import (
	"testing"

	"github.com/alperkose/confunc"
)


func Test_MapSource(t *testing.T){
	expectedValue := "some param"
	configurationKey := "myConfig"

	var sut confunc.Source

	sut = confunc.Map(map[string]string{configurationKey:expectedValue})
	actualValue := sut.Value(configurationKey)
	if actualValue != expectedValue {
		t.Errorf("expected %v to be %v", actualValue, expectedValue)
	}

}

func Test_MapSourceWhenThereIsNoValueForGivenKey(t *testing.T){
	configurationKey := "nonExistantKey"
	var sut confunc.Source

	sut = confunc.Map(map[string]string{"key1":"value1"})
	actualValue := sut.Value(configurationKey)
	if len(actualValue) > 0{
		t.Errorf("expected %v to be empty", actualValue)
	}

}