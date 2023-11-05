package dynarray

import "github.com/kkkunny/stl/container/iterator"

func (_ DynArray[T]) NewWithIterator(iter iterator.Iterator[T]) any {
	self := NewDynArrayWithLength[T](iter.Length())
	var i int
	for iter.Next() {
		(*self.data)[i] = iter.Value()
		i++
	}
	return self
}

// Iterator 迭代器
func (self DynArray[T]) Iterator() iterator.Iterator[T] {
	self.init()
	return newIterator[T](self.data)
}

type _Iterator[T any] struct {
	data *[]T
	next uint
}

func newIterator[T any](data *[]T) *_Iterator[T] {
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
