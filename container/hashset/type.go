package hashset

import "github.com/kkkunny/stl/container/hashmap"

// HashSet 哈希set
type HashSet[T any] hashmap.HashMap[T, struct{}]

func NewHashSet[T any]() HashSet[T] {
	return HashSet[T](hashmap.NewHashMap[T, struct{}]())
}

func NewHashSetWithCapacity[T any](cap uint) HashSet[T] {
	return HashSet[T](hashmap.NewHashMapWithCapacity[T, struct{}](cap))
}

func NewHashSetWith[T any](vs ...T) HashSet[T] {
	self := NewHashSetWithCapacity[T](uint(len(vs)))
	for _, v := range vs {
		self.Add(v)
	}
	return self
}
