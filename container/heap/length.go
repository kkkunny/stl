package heap

// Length 长度
func (self Heap[T]) Length() uint {
	return uint(len(self.data))
}
