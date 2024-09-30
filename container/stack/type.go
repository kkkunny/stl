package stack

// Stack 栈
type Stack[T any] []T

// NewStack 新建栈
func NewStack[T any]() Stack[T] {
	return make([]T, 0)
}

// NewStackWith 新建指定元素的栈
func NewStackWith[T any](vs ...T) Stack[T] {
	return vs
}

func (self Stack[T]) Default() Stack[T] {
	return NewStack[T]()
}
