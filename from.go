package confunc

type Parameterizer struct {
	s Source
}

func (p *Parameterizer) String(sourceKey string, wrappers ...Wrapper) String{
	return func() string {
		value := p.s.Value(sourceKey)
		for _, w := range wrappers {
			value = w(value)
		}
		return value
	}
}

func From(s Source) *Parameterizer{
	return &Parameterizer{s}
}
