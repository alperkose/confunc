package confunc

type Parameterizer struct {
	s Source
}

func (p *Parameterizer) StringFunc(sourceKey string) *StringBuilder{
	return &StringBuilder{p.s, sourceKey}
}

func From(s Source) *Parameterizer{
	return &Parameterizer{s}
}
