package linkedlist

type _iterator[T any] struct {
	src  *LinkedList[T]
	cursor *_LinkedListNode[T]
}

func _NewIterator[T any](src *LinkedList[T]) *_iterator[T] {
	return &_iterator[T]{
		src:  src,
		cursor: nil,
	}
}

func (self _iterator[T]) Length() uint {
	return self.src.Length()
}

func (self *_iterator[T]) Next() bool {
	if self.cursor == nil{
		self.cursor = self.src.root
		return self.cursor != nil
	}else{
		self.cursor = self.cursor.Next
		return self.cursor != nil
	}
}

func (self _iterator[T]) HasNext() bool {
	if self.cursor == nil{
		return !self.src.Empty()
	}else{
		return self.cursor.Next != nil
	}
}

func (self _iterator[T]) Value() T {
	return self.cursor.Value
}

func (self *_iterator[T]) Reset() {
	self.cursor = nil
}
