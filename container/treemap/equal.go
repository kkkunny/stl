package treemap

import (
	stlslices "github.com/kkkunny/stl/container/slices"
)

// Equal 比较
func (self TreeMap[K, V]) Equal(dst TreeMap[K, V]) bool {
	self.init()
	dst.init()

	if self.tree == dst.tree {
		return true
	} else if self.Length() != dst.Length() {
		return false
	}

	return stlslices.Equal(self.KeyValues(), dst.KeyValues())
}
