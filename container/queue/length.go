package queue

import "github.com/kkkunny/stl/container/dynarray"

// Length 长度
func (self Queue[T]) Length() uint {
	return dynarray.DynArray[T](self).Length()
}
