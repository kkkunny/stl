package stack

import (
	"github.com/kkkunny/stl/container/dynarray"
)

// Equal 比较相等
func (self Stack[T]) Equal(dst Stack[T]) bool {
	return dynarray.DynArray[T](self).Equal(dynarray.DynArray[T](dst))
}
