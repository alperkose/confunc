package cfconsul

import (
	"github.com/alperkose/confunc"
	"github.com/hashicorp/consul/api"
)

type consul struct {
	cfg *api.Config
}

func (s *consul) Value(k string) string {
	consulAPI, err := api.NewClient(s.cfg)
	if err != nil {
		return ""
	}
	pair, q, err := consulAPI.KV().Get(k, nil)
	if err != nil || pair == nil {
		return ""
	}

	return string(pair.Value)

}

func Source(config *api.Config) confunc.Source {
	return &consul{config}
}
