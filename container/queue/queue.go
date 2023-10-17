package queue

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/iterator"
)

// Queue 队列
type Queue[T any] struct {
	data *[]T
}

func NewQueue[T any]() Queue[T] {
	var data []T
	return Queue[T]{data: &data}
}

func NewQueueWithCapacity[T any](cap uint) Queue[T] {
	data := make([]T, 0, cap)
	return Queue[T]{data: &data}
}

func NewQueueWith[T any](vs ...T) Queue[T] {
	return Queue[T]{data: &vs}
}

func (_ Queue[T]) NewWithIterator(iter iterator.Iterator[T]) Queue[T] {
	self := NewQueueWithCapacity[T](iter.Length())
	var i uint
	for iter.Next() {
		self.Push(iter.Value())
		i++
	}
	return self
}

func (self Queue[T]) Length() uint {
	return uint(len(*self.data))
}

func (self Queue[T]) Capacity() uint {
	return uint(cap(*self.data))
}

func (self Queue[T]) Equal(dst Queue[T]) bool {
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

func (self *Queue[T]) Push(v T) {
	*self.data = append(*self.data, v)
}

func (self *Queue[T]) Pop() T {
	v := (*self.data)[0]
	*self.data = (*self.data)[1:]
	return v
}

func (self Queue[T]) String() string {
	var buf strings.Builder
	buf.WriteString("queue[")
	for i, v := range *self.data {
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < len(*self.data)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func (self Queue[T]) Debug(preifx uint) string {
	var buf strings.Builder
	buf.WriteString("queue{")
	for i, v := range *self.data {
		buf.WriteString(stlbasic.Debug(v, preifx))
		if i < len(*self.data)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self Queue[T]) Peek() T {
	return (*self.data)[0]
}

func (self Queue[T]) Clone() Queue[T] {
	return Queue[T]{data: stlbasic.Clone(self.data)}
}

func (self *Queue[T]) Clear() {
	*self.data = make([]T, 0)
}

func (self Queue[T]) Empty() bool {
	return self.Length() == 0
}

func (self Queue[T]) Iterator() iterator.Iterator[T] {
	return iterator.NewIterator[T](_NewIterator[T](&self))
}
