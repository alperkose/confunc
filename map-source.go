package confunc

import "errors"

type mapSource struct {
	ctx map[string]string
}

func (s *mapSource) Value(key string) (string, error) {
	v, ok := s.ctx[key]
	if !ok {
		return "", errors.New("No value is retrieved")
	}
	return v, nil
}

func Map(c map[string]string) Source {
	return &mapSource{c}
}
