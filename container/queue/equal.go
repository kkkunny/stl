package queue

import (
	stlslices "github.com/kkkunny/stl/container/slices"
)

// Equal 比较相等
func (self Queue[T]) Equal(dst Queue[T]) bool {
	return stlslices.Equal(self, dst)
}
