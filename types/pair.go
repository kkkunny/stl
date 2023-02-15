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

// ThreePair 三组合
type ThreePair[T1, T2, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

// NewThreePair 新建组合
func NewThreePair[T1, T2, T3 any](f T1, s T2, t T3) ThreePair[T1, T2, T3] {
	return ThreePair[T1, T2, T3]{First: f, Second: s, Third: t}
}
