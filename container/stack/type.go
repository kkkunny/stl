package stack

import "github.com/kkkunny/stl/container/dynarray"

// Stack 栈
type Stack[T any] dynarray.DynArray[T]

// NewStack 新建栈
func NewStack[T any]() Stack[T] {
	return Stack[T](dynarray.NewDynArray[T]())
}

// NewStackWith 新建指定元素的栈
func NewStackWith[T any](vs ...T) Stack[T] {
	return Stack[T](dynarray.NewDynArrayWith[T](vs...))
}
