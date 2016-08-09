package confunc

type String func() string

type Integer func() int

type Boolean func() bool

type Float64 func() float64

type Interceptor func(String) string

type Source interface {
	Value(v string) string
}
