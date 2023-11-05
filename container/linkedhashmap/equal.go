package linkedhashmap

import stlbasic "github.com/kkkunny/stl/basic"

func (self LinkedHashMap[K, V]) Equal(dst LinkedHashMap[K, V]) bool {
	self.init()
	if self.Length() != dst.Length() {
		return false
	} else if self.list == dst.list {
		return true
	}

	for c1, c2 := self.list.Front(), dst.list.Front(); c1 != nil && c2 != nil; c1, c2 = c1.Next(), c2.Next() {
		v1, v2 := c1.Value, c2.Value
		if !stlbasic.Equal(v1.First, v2.First) || !stlbasic.Equal(v1.Second, v2.Second) {
			return false
		}
	}
	return true
}
