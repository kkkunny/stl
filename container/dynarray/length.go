package dynarray

// Length 长度
func (self DynArray[T]) Length() uint {
	self.init()
	return uint(len(*self.data))
}
