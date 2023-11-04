package dynarray

// Capacity 容量
func (self DynArray[T]) Capacity() uint {
	self.init()
	return uint(cap(*self.data))
}
