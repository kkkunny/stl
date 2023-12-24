package linkedhashset

import (
	"github.com/kkkunny/stl/container/linkedhashmap"
)

// LinkedHashSet 插入顺序哈希set
type LinkedHashSet[T any] linkedhashmap.LinkedHashMap[T, struct{}]

func NewLinkedHashSet[T any]() LinkedHashSet[T] {
	return LinkedHashSet[T](linkedhashmap.NewLinkedHashMap[T, struct{}]())
}

func NewLinkedHashSetWithCapacity[T any](cap uint) LinkedHashSet[T] {
	return LinkedHashSet[T](linkedhashmap.NewLinkedHashMapWithCapacity[T, struct{}](cap))
}

func NewLinkedHashSetWith[T any](vs ...T) LinkedHashSet[T] {
	self := NewLinkedHashSetWithCapacity[T](uint(len(vs)))
	for _, v := range vs {
		self.Add(v)
	}
	return self
}

func (self LinkedHashSet[T]) Default() LinkedHashSet[T] {
	return NewLinkedHashSet[T]()
}
