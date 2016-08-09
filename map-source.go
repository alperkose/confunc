package confunc

type mapSource struct {
	ctx map[string]string
}

func (s *mapSource) Value(key string) string {
	return s.ctx[key]
}

func Map(c map[string]string) Source {
	return &mapSource{c}
}
