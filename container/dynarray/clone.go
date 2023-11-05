package dynarray

// Clone 克隆
func (self DynArray[T]) Clone() DynArray[T] {
	self.init()
	newData := make([]T, len(*self.data), cap(*self.data))
	copy(newData, *self.data)
	return DynArray[T]{data: &newData}
}
