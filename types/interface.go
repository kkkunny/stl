package types

// 哈希
type Hasher interface {
	Hash() int32
}

// 比较
type Comparator[T any] interface {
	Compare(T) int
}
