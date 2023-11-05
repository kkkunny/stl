package pair

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
