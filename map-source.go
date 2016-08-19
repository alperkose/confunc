package confunc

type mapSource struct {
	ctx map[string]string
}

func (s *mapSource) Value(key string) (string, error) {
	return s.ctx[key], nil
}

func Map(c map[string]string) Source {
	return &mapSource{c}
}
