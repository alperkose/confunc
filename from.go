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
	strFunc := p.String(sourceKey, interceptors...)
	base := func() int {
		val, _ := strconv.Atoi(strFunc())
		return val
	}

	return base
}

func (p *Parameterizer) Bool(sourceKey string, interceptors ...Interceptor) Boolean {
	strFunc := p.String(sourceKey, interceptors...)
	base := func() bool {
		val, _ := strconv.ParseBool(strFunc())
		return val
	}

	return base
}

func (p *Parameterizer) Float64(sourceKey string, interceptors ...Interceptor) Float64 {
	strFunc := p.String(sourceKey, interceptors...)
	base := func() float64 {
		val, _ := strconv.ParseFloat(strFunc(), 64)
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
