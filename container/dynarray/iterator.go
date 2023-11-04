package dynarray

import "github.com/kkkunny/stl/container/iterator"

// AppendWithIterator 通过迭代器追加
func (self *DynArray[T]) AppendWithIterator(iter iterator.Iterator[T]) {
	self.init()
	for iter.Next() {
		self.PushBack(iter.Value())
	}
}

// Append 追加
func (self *DynArray[T]) Append(ctr iterator.IteratorContainer[T]) {
	self.AppendWithIterator(ctr.Iterator())
}

// Iterator 迭代器
func (self DynArray[T]) Iterator() iterator.Iterator[T] {
	self.init()
	return iterator.NewIterator[T](_NewIterator[T](self.data))
}

type _Iterator[T any] struct {
	data *[]T
	next uint
}

func _NewIterator[T any](data *[]T) *_Iterator[T] {
	return &_Iterator[T]{
		data: data,
		next: 0,
	}
}

func (self _Iterator[T]) Length() uint {
	return uint(len(*self.data))
}

func (self *_Iterator[T]) Next() bool {
	if self.next >= self.Length() {
		return false
	}
	self.next++
	return true
}

func (self _Iterator[T]) HasNext() bool {
	return self.next < self.Length()
}

func (self _Iterator[T]) Value() T {
	return (*self.data)[self.next-1]
}

func (self *_Iterator[T]) Reset() {
	self.next = 0
}
