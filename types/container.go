package types

// 组合
type Pair[T any, E any] struct {
	First  T
	Second E
}

// 新建组合
func NewPair[T any, E any](f T, s E) Pair[T, E] {
	return Pair[T, E]{First: f, Second: s}
}
