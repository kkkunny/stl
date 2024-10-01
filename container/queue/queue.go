package queue

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	stlslices "github.com/kkkunny/stl/container/slices"
)

type Queue[T any] interface {
	queueIter[T]
	stlbasic.Cloneable[Queue[T]]
	stlcmp.Equalable[Queue[T]]
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

func New[T any](vs ...T) Queue[T] {
	if len(vs) == 0 {
		return _NewQueue[T]()
	} else {
		return _NewQueueWith[T](vs...)
	}
}

type _Queue[T any] struct {
	data []T
}

func _NewQueue[T any]() Queue[T] {
	return &_Queue[T]{data: make([]T, 0)}
}

func _NewQueueWith[T any](vs ...T) Queue[T] {
	return &_Queue[T]{data: vs}
}

// String 获取字符串
func (self *_Queue[T]) String() string {
	var buf strings.Builder
	buf.WriteString("Queue{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Push 入队
func (self *_Queue[T]) Push(v T, vs ...T) {
	self.data = append(self.data, v)
	self.data = append(self.data, vs...)
}

// Pop 出队
func (self *_Queue[T]) Pop() T {
	v := stlslices.First(self.data)
	self.data = self.data[1:]
	return v
}

// Peek 头元素
func (self *_Queue[T]) Peek() T {
	return stlslices.First(self.data)
}

// Clear 清空
func (self *_Queue[T]) Clear() {
	self.data = make([]T, 0)
}

// Empty 是否为空
func (self *_Queue[T]) Empty() bool {
	return len(self.data) == 0
}

// ToSlice 转成切片
func (self *_Queue[T]) ToSlice() []T {
	return stlslices.Clone(self.data)
}

// Length 长度
func (self *_Queue[T]) Length() uint {
	return uint(len(self.data))
}

func (self *_Queue[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	data := make([]T, iter.Length())
	var i int
	for iter.Next() {
		data[i] = iter.Value()
		i++
	}
	return &_Queue[T]{data: data}
}

// Iterator 迭代器
func (self *_Queue[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(self.data...)
}

// Equal 比较相等
func (self *_Queue[T]) Equal(dst Queue[T]) bool {
	return stlslices.Equal(self.data, dst.getData())
}

// Clone 克隆
func (self *_Queue[T]) Clone() Queue[T] {
	return &_Queue[T]{data: stlslices.Clone(self.data)}
}

func (self *_Queue[T]) getData() []T {
	return self.data
}
