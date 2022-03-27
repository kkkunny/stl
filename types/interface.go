package types

// 哈希
type Hasher interface {
	Hash() Usize
}

// 比较
type Comparator[T any] interface {
	Compare(T) int
}
