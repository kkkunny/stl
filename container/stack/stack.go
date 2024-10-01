package stack

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/internal/slices"
	"github.com/kkkunny/stl/value"
)

type Stack[T any] interface {
	stackIter[T]
	stlval.Cloneable[Stack[T]]
	stlcmp.Equalable[Stack[T]]
	stliter.IteratorContainer[T]
	stlbasic.Lengthable
	Push(v T, vs ...T)
	Pop() T
	Peek() T
	Clear()
	Empty() bool
	fmt.Stringer
	ToSlice() []T
	getData() []T
}

func New[T any](vs ...T) Stack[T] {
	if len(vs) == 0 {
		return _NewStack[T]()
	} else {
		return _NewStackWith[T](vs...)
	}
}

type _Stack[T any] struct {
	data []T
}

func _NewStack[T any]() Stack[T] {
	return &_Stack[T]{data: make([]T, 0)}
}

func _NewStackWith[T any](vs ...T) Stack[T] {
	return &_Stack[T]{data: vs}
}

// String 获取字符串
func (self *_Stack[T]) String() string {
	var buf strings.Builder
	buf.WriteString("Stack{")
	for i, v := range self.data {
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < len(self.data)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Push 入队
func (self *_Stack[T]) Push(v T, vs ...T) {
	self.data = append(self.data, v)
	self.data = append(self.data, vs...)
}

// Pop 出队
func (self *_Stack[T]) Pop() T {
	v := stlslices.Last(self.data)
	self.data = self.data[:len(self.data)-1]
	return v
}

// Peek 头元素
func (self *_Stack[T]) Peek() T {
	return stlslices.Last(self.data)
}

// Clear 清空
func (self *_Stack[T]) Clear() {
	self.data = make([]T, 0)
}

// Empty 是否为空
func (self *_Stack[T]) Empty() bool {
	return len(self.data) == 0
}

// ToSlice 转成切片
func (self *_Stack[T]) ToSlice() []T {
	return stlslices.Clone(self.data)
}

// Length 长度
func (self *_Stack[T]) Length() uint {
	return uint(len(self.data))
}

func (self *_Stack[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	data := make([]T, iter.Length())
	var i int
	for iter.Next() {
		data[i] = iter.Value()
		i++
	}
	return &_Stack[T]{data: data}
}

// Iterator 迭代器
func (self *_Stack[T]) Iterator() stliter.Iterator[T] {
	reverse := stlslices.Clone(self.data)
	slices.Reverse(reverse)
	return stliter.NewSliceIterator(reverse...)
}

// Equal 比较相等
func (self *_Stack[T]) Equal(dst Stack[T]) bool {
	return stlslices.Equal(self.data, dst.getData())
}

// Clone 克隆
func (self *_Stack[T]) Clone() Stack[T] {
	return &_Stack[T]{data: stlslices.Clone(self.data)}
}

func (self *_Stack[T]) getData() []T {
	return self.data
}
