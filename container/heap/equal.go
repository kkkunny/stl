package heap

import stlslices "github.com/kkkunny/stl/container/slices"

// Equal 比较相等 O(3N+2log(N))
func (self Heap[T]) Equal(dst Heap[T]) bool {
	self.init()

	if self.Length() != dst.Length() {
		return false
	}
	lv, rv := stlslices.Sort(self.data), stlslices.Sort(dst.data)
	return stlslices.Equal(lv, rv)
}
