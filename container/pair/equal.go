package pair

import (
	stlbasic "github.com/kkkunny/stl/cmp"
)

func (self Pair[T, F]) Equal(dst Pair[T, F]) bool {
	return stlbasic.Equal(self.First, dst.First) && stlbasic.Equal(self.Second, dst.Second)
}
