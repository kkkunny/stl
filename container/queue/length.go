package queue

// Length 长度
func (self Queue[T]) Length() uint {
	return uint(len(self))
}
