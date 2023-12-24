package stack

// Equal 比较相等 O(3N+2log(N))
func (self Heap[T]) Equal(dst Heap[T]) bool {
	self.init()

	if self.Length() != dst.Length() {
		return false
	}
	lv, rv := self.data.Clone(), dst.data.Clone()
	lv.Sort(self.reverse)
	rv.Sort(dst.reverse)
	return lv.Equal(rv)
}
