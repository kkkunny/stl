package queue

import (
	stlslices "github.com/kkkunny/stl/container/slices"
)

// Clone 克隆
func (self Queue[T]) Clone() Queue[T] {
	return stlslices.Clone(self)
}
