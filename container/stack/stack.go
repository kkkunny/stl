package stack

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/iterator"
)

// Stack 栈
type Stack[T any] struct {
	data *[]T
}

func NewStack[T any]() Stack[T] {
	var data []T
	return Stack[T]{data: &data}
}

func NewStackWithCapacity[T any](cap uint) Stack[T] {
	data := make([]T, 0, cap)
	return Stack[T]{data: &data}
}

func NewStackWith[T any](vs ...T) Stack[T] {
	return Stack[T]{data: &vs}
}

func (_ Stack[T]) NewWithIterator(iter iterator.Iterator[T]) Stack[T] {
	// TODO: 先反转再插入
	self := NewStackWithCapacity[T](iter.Length())
	var i uint
	for iter.Next() {
		self.Push(iter.Value())
		i++
	}
	return self
}

func (self Stack[T]) Length() uint {
	return uint(len(*self.data))
}

func (self Stack[T]) Capacity() uint {
	return uint(cap(*self.data))
}

func (self Stack[T]) Equal(dst Stack[T]) bool {
	if self.data == dst.data {
		return true
	}

	if self.Length() != dst.Length() {
		return false
	}

	for i, v := range *self.data {
		if !stlbasic.Equal(v, (*dst.data)[i]) {
			return false
		}
	}
	return true
}

func (self *Stack[T]) Push(v T) {
	*self.data = append(*self.data, v)
}

func (self *Stack[T]) Pop() T {
	v := (*self.data)[len(*self.data)-1]
	*self.data = (*self.data)[:len(*self.data)-1]
	return v
}

func (self Stack[T]) String() string {
	var buf strings.Builder
	buf.WriteString("stack[")
	for i, v := range *self.data {
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < len(*self.data)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func (self Stack[T]) Peek() T {
	return (*self.data)[len(*self.data)-1]
}

func (self Stack[T]) Clone() Stack[T] {
	return Stack[T]{data: stlbasic.Clone(self.data)}
}

func (self *Stack[T]) Clear() {
	*self.data = make([]T, 0)
}

func (self Stack[T]) Empty() bool {
	return self.Length() == 0
}

func (self Stack[T]) Iterator() iterator.Iterator[T] {
	return iterator.NewIterator[T](_NewIterator[T](&self))
}
