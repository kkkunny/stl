package pair

import (
	stlbasic "github.com/kkkunny/stl/cmp"
)

func (self Pair[T, F]) Equal(dstObj any) bool {
	dst, ok := dstObj.(Pair[T, F])
	if !ok {
		return false
	}
	return stlbasic.Equal(self.First, dst.First) && stlbasic.Equal(self.Second, dst.Second)
}
