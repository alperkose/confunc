package confunc

import "os"

type env struct{}

func (s *env) Value(k string) (string, error) {
	return os.Getenv(k), nil
}

func Env() Source {
	return &env{}
}
