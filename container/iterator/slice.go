package iterator

type _SliceIterator[T any] struct {
	src  *[]T
	next int
}

func _NewSliceIterator[T any](src ...T) *_SliceIterator[T] {
	return &_SliceIterator[T]{
		src:  &src,
		next: 0,
	}
}

func (self *_SliceIterator[T]) Length() uint {
	return uint(len(*self.src))
}

func (self *_SliceIterator[T]) Next() bool {
	if self.next >= len(*self.src) {
		return false
	}
	self.next++
	return true
}

func (self _SliceIterator[T]) HasNext() bool {
	return self.next < len(*self.src)
}

func (self _SliceIterator[T]) Value() T {
	return (*self.src)[self.next-1]
}

func (self *_SliceIterator[T]) Reset() {
	self.next = 0
}
