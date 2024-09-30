package stack

import (
	stlslices "github.com/kkkunny/stl/container/slices"
)

// Clone 克隆
func (self Stack[T]) Clone() Stack[T] {
	return stlslices.Clone(self)
}
