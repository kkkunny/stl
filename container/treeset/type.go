package treeset

import (
	"github.com/kkkunny/stl/container/treemap"
)

// TreeSet 有序set
type TreeSet[T any] treemap.TreeMap[T, struct{}]

func NewTreeSet[T any]() TreeSet[T] {
	return TreeSet[T](treemap.NewTreeMap[T, struct{}]())
}

func NewTreeSetWith[T any](vs ...T) TreeSet[T] {
	self := NewTreeSet[T]()
	for _, v := range vs {
		self.Add(v)
	}
	return self
}

func (self TreeSet[T]) Default() TreeSet[T] {
	return NewTreeSet[T]()
}
