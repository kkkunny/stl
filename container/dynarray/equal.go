package dynarray

import stlbasic "github.com/kkkunny/stl/basic"

// Equal 比较相等
func (self DynArray[T]) Equal(dst DynArray[T]) bool {
	self.init()

	if self.data == dst.data {
		return true
	}

	if self.Length() != dst.Length() {
		return false
	}

	for i, v := range *self.data {
		if !stlbasic.Equal(v, (*dst.data)[i]) {
			return false
		}
	}
	return true
}
