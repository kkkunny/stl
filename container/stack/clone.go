package stack

import "github.com/kkkunny/stl/container/dynarray"

// Clone 克隆
func (self Stack[T]) Clone() Stack[T] {
	return Stack[T](dynarray.DynArray[T](self).Clone())
}
