package cfconsul_test

import (
	"github.com/alperkose/confunc/cfconsul"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/testutil"
	"testing"
)

func Test_CounsulSource(t *testing.T) {

	consulServer := testutil.NewTestServer(t)
	defer consulServer.Stop()
	configKey := "someKey"
	expectedConfigValue := "someValue"
	consulServer.SetKV(configKey, []byte(expectedConfigValue))

	consulSource := cfconsul.Source(&api.Config{
		Address: consulServer.HTTPAddr,
	})
	actualConfigValue, err := consulSource.Value(configKey)

	if err != nil {
		t.Errorf("error should not have occurred : %v", err.Error())
	}

	if actualConfigValue != expectedConfigValue {
		t.Errorf("expected '%v' to be '%v'", actualConfigValue, expectedConfigValue)
	}
}

func Test_ConsulSourceWhenNoConfigWasSet(t *testing.T) {
	consulServer := testutil.NewTestServer(t)
	defer consulServer.Stop()
	configKey := "someKey"

	consulSource := cfconsul.Source(&api.Config{
		Address: consulServer.HTTPAddr,
	})
	actualConfigValue, err := consulSource.Value(configKey)

	if err == nil {
		t.Errorf("an error should have occurred since there is no configuration")
	}

	if len(actualConfigValue) > 0 {
		t.Errorf("expected '%v' to be empty", actualConfigValue)
	}
}

func Test_ConsulSourceWhenConsulIsNotAccessible(t *testing.T) {
	configKey := "someKey"

	consulSource := cfconsul.Source(api.DefaultConfig())
	actualConfigValue, err := consulSource.Value(configKey)

	if err == nil {
		t.Errorf("an error should have occurred since there is no configuration")
	}

	if len(actualConfigValue) > 0 {
		t.Errorf("expected '%v' to be empty", actualConfigValue)
	}

}
