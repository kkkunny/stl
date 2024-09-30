package stack

import (
	stlslices "github.com/kkkunny/stl/container/slices"
)

// Equal 比较相等
func (self Stack[T]) Equal(dst Stack[T]) bool {
	return stlslices.Equal(self, dst)
}
