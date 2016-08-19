package cfconsul

import (
	"errors"
	"github.com/alperkose/confunc"
	"github.com/hashicorp/consul/api"
)

type consul struct {
	cfg *api.Config
}

func (s *consul) Value(k string) (string, error) {
	consulAPI, err := api.NewClient(s.cfg)
	if err != nil {
		return "", err
	}
	pair, _, err := consulAPI.KV().Get(k, nil)
	if err != nil {
		return "", err
	}
	if pair == nil {
		return "", errors.New("Configuration not found")
	}

	return string(pair.Value), nil

}

func Source(config *api.Config) confunc.Source {
	return &consul{config}
}
