package confunc

import (
	"errors"
	"os"
)

type env struct{}

func (s *env) Value(k string) (string, error) {
	v, ok := os.LookupEnv(k)
	if !ok {
		return v, errors.New("No value is retrieved")
	}
	return v, nil
}

func Env() Source {
	return &env{}
}
