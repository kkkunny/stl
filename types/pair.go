package types

// Pair 组合
type Pair[T, E any] struct {
	First  T
	Second E
}

// NewPair 新建组合
func NewPair[T, E any](f T, s E) Pair[T, E] {
	return Pair[T, E]{First: f, Second: s}
}
