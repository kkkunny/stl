package treemap

import stlbasic "github.com/kkkunny/stl/basic"

// Equal 比较
func (self TreeMap[K, V]) Equal(dst TreeMap[K, V]) bool {
	self.init()
	dst.init()

	if self.tree == dst.tree {
		return true
	} else if self.Length() != dst.Length() {
		return false
	}

	var i uint
	for selfIter, dstIter := self.KeyValues().Iterator(), dst.KeyValues().Iterator(); selfIter.Next() && dstIter.Next(); i++ {
		if !stlbasic.Equal(selfIter.Value(), dstIter.Value()) {
			return false
		}
	}
	return true
}
