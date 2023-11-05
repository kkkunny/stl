package hashset

import (
	"github.com/kkkunny/stl/container/hashmap"
)

// HashSet 哈希set
type HashSet[T any] struct {
	data hashmap.HashMap[T, struct{}]
}

func NewHashSet[T any]() HashSet[T] {
	return HashSet[T]{data: hashmap.NewHashMap[T, struct{}]()}
}

func NewHashSetWithCapacity[T any](cap uint) HashSet[T] {
	return HashSet[T]{data: hashmap.NewHashMapWithCapacity[T, struct{}](cap)}
}

func NewHashSetWith[T any](vs ...T) HashSet[T] {
	self := NewHashSetWithCapacity[T](uint(len(vs)))
	for _, v := range vs {
		self.Push(v)
	}
	return self
}
