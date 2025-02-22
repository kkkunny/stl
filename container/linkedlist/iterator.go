package linkedlist

import (
	"iter"

	stliter "github.com/kkkunny/stl/container/iter"
)

func (_ LinkedList[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	self := NewLinkedList[T]()
	for iter.Next() {
		self.PushBack(iter.Value())
	}
	return self
}

// Iterator 迭代器
func (self LinkedList[T]) Iterator() stliter.Iterator[T] {
	return newIterator[T](&self)
}

func (self LinkedList[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for cursor := self.root; cursor != nil; cursor = cursor.Next {
			if !yield(cursor.Value) {
				return
			}
		}
	}
}

type _Iterator[T any] struct {
	src    *LinkedList[T]
	cursor *node[T]
}

func newIterator[T any](src *LinkedList[T]) *_Iterator[T] {
	return &_Iterator[T]{
		src:    src,
		cursor: nil,
	}
}

func (self _Iterator[T]) Length() uint {
	return self.src.Length()
}

func (self *_Iterator[T]) Next() bool {
	if self.cursor == nil {
		self.cursor = self.src.root
		return self.cursor != nil
	} else {
		self.cursor = self.cursor.Next
		return self.cursor != nil
	}
}

func (self _Iterator[T]) HasNext() bool {
	if self.cursor == nil {
		return !self.src.Empty()
	} else {
		return self.cursor.Next != nil
	}
}

func (self _Iterator[T]) Value() T {
	return self.cursor.Value
}

func (self *_Iterator[T]) Reset() {
	self.cursor = nil
}
