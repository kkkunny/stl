package stack

// Length 长度
func (self Stack[T]) Length() uint {
	return uint(len(self))
}
