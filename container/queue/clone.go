package queue

import "github.com/kkkunny/stl/container/dynarray"

// Clone 克隆
func (self Queue[T]) Clone() Queue[T] {
	return Queue[T](dynarray.DynArray[T](self).Clone())
}
