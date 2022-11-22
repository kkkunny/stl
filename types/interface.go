package types

// Comparator 比较
type Comparator[T any] interface {
	Compare(T) int
}
