package stliter

type sliceIterator[T any] struct {
	src  *[]T
	next int
}

func NewSliceIterator[T any](src ...T) Iterator[T] {
	return &sliceIterator[T]{
		src:  &src,
		next: 0,
	}
}

func (self *sliceIterator[T]) Length() uint {
	return uint(len(*self.src))
}

func (self *sliceIterator[T]) Next() bool {
	if self.next >= len(*self.src) {
		return false
	}
	self.next++
	return true
}

func (self sliceIterator[T]) HasNext() bool {
	return self.next < len(*self.src)
}

func (self sliceIterator[T]) Value() T {
	return (*self.src)[self.next-1]
}

func (self *sliceIterator[T]) Reset() {
	self.next = 0
}
