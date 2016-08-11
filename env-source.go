package confunc

import "os"

type env struct{}

func (s *env) Value(k string) string {
	return os.Getenv(k)
}

func Env() Source {
	return &env{}
}
