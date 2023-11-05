package hashset

// Empty 是否为空
func (self HashSet[T]) Empty() bool {
	return self.data.Empty()
}

// Clear 清空
func (self *HashSet[T]) Clear() {
	self.data.Clear()
}

// Push 增加元素
func (self *HashSet[T]) Push(v T) bool {
	exist := self.data.ContainKey(v)
	self.data.Set(v, struct{}{})
	return !exist
}

// Remove 移除元素
func (self *HashSet[T]) Remove(v T) T {
	self.data.Remove(v)
	return v
}

// ToSlice 转成切片
func (self *HashSet[T]) ToSlice() []T {
	return self.data.Keys().ToSlice()
}
