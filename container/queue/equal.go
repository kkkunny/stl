package queue

import (
	"github.com/kkkunny/stl/container/dynarray"
)

// Equal 比较相等
func (self Queue[T]) Equal(dst Queue[T]) bool {
	return dynarray.DynArray[T](self).Equal(dynarray.DynArray[T](dst))
}
