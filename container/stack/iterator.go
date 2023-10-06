package stack

type _iterator[T any] struct {
	src  *Stack[T]
	next uint
}

func _NewIterator[T any](src *Stack[T]) *_iterator[T] {
	return &_iterator[T]{
		src:  src,
		next: 0,
	}
}

func (self _iterator[T]) Length() uint {
	return self.src.Length()
}

func (self *_iterator[T]) Next() bool {
	if self.next >= self.Length() {
		return false
	}
	self.next++
	return true
}

func (self _iterator[T]) HasNext() bool {
	return self.next < self.Length()
}

func (self _iterator[T]) Value() T {
	return (*self.src.data)[uint(len(*self.src.data))-self.next]
}

func (self *_iterator[T]) Reset() {
	self.next = 0
}
