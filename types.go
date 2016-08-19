package confunc

type String func() string

type Integer func() int

type Boolean func() bool

type Float64 func() float64

type Confunc func() (string, error)

type Interceptor func(Confunc) Confunc

type Source interface {
	Value(v string) (string, error)
}
