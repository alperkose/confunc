package confunc

type String func() string

type Integer func() int

type Boolean func() bool

type Float64 func() float64

type Wrapper func(string) string

type Source interface {
	Value(key string) string
}