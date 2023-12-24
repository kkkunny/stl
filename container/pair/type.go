package pair

import stlbasic "github.com/kkkunny/stl/basic"

// Pair 对
type Pair[T, F any] struct {
	First  T
	Second F
}

func NewPair[T, F any](first T, second F) Pair[T, F] {
	return Pair[T, F]{
		First:  first,
		Second: second,
	}
}

func (self Pair[T, F]) Default() Pair[T, F] {
	return NewPair[T, F](stlbasic.Default[T](), stlbasic.Default[F]())
}
