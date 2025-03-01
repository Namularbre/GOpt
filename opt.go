package gopt

type Opt[T any] struct {
	Value T
	Valid bool
}

type Res[T any] struct {
	Some T
	Err  error
}
