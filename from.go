package confunc

import "strconv"

type Parameterizer struct {
	s Source
}

func (p *Parameterizer) String(sourceKey string, interceptors ...Interceptor) String {

	base := func() string {
		return p.s.Value(sourceKey)
	}

	for _, icp := range interceptors {
		base = convertInterceptor(icp, base)
	}

	return base

}

func (p *Parameterizer) Int(sourceKey string, interceptors ...Interceptor) Integer {
	base := func() int {
		strFunc := p.String(sourceKey, interceptors...)
		val, _ := strconv.Atoi(strFunc())
		return val
	}

	return base
}

func convertInterceptor(i Interceptor, base String) String {
	return func() string {
		return i(base)
	}
}

func From(s Source) *Parameterizer {
	return &Parameterizer{s}
}
