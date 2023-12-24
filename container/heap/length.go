package heap

// Length 长度
func (self Heap[T]) Length() uint {
	self.init()
	return self.data.Length()
}
