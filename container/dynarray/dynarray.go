package dynarray

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/iterator"
)

// DynArray 动态数组
type DynArray[T any] struct {
	data *[]T
}

func (self *DynArray[T]) init() {
	if self.data != nil {
		return
	}
	self.data = new([]T)
}

func NewDynArray[T any]() DynArray[T] {
	return DynArray[T]{data: new([]T)}
}

func NewDynArrayWithCapacity[T any](cap uint) DynArray[T] {
	data := make([]T, 0, cap)
	return DynArray[T]{data: &data}
}

func NewDynArrayWithLength[T any](l uint) DynArray[T] {
	data := make([]T, l)
	return DynArray[T]{data: &data}
}

func NewDynArrayWith[T any](vs ...T) DynArray[T] {
	return DynArray[T]{data: &vs}
}

func (_ DynArray[T]) NewWithIterator(iter iterator.Iterator[T]) DynArray[T] {
	self := NewDynArrayWithLength[T](iter.Length())
	var i uint
	for iter.Next() {
		self.Set(i, iter.Value())
		i++
	}
	return self
}

func (self DynArray[T]) Length() uint {
	self.init()
	return uint(len(*self.data))
}

func (self DynArray[T]) Capacity() uint {
	self.init()
	return uint(cap(*self.data))
}

func (self DynArray[T]) Equal(dst DynArray[T]) bool {
	self.init()

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

func (self DynArray[T]) Get(i uint) T {
	self.init()
	return (*self.data)[i]
}

func (self *DynArray[T]) Set(i uint, v T) T {
	self.init()
	pv := (*self.data)[i]
	(*self.data)[i] = v
	return pv
}

func (self *DynArray[T]) PushBack(v T) {
	self.init()
	*self.data = append(*self.data, v)
}

func (self *DynArray[T]) PushFront(v T) {
	self.init()
	*self.data = append([]T{v}, *self.data...)
}

func (self *DynArray[T]) Insert(i uint, v T) {
	self.init()
	*self.data = append(*self.data, v)
	copy((*self.data)[i+1:], (*self.data)[i:])
	(*self.data)[i] = v
}

func (self *DynArray[T]) Remove(i uint) T {
	self.init()
	v := (*self.data)[i]
	copy((*self.data)[i:], (*self.data)[i+1:])
	*self.data = (*self.data)[:len(*self.data)-1]
	return v
}

func (self DynArray[T]) String() string {
	self.init()
	var buf strings.Builder
	buf.WriteByte('[')
	for i, v := range *self.data {
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < len(*self.data)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func (self *DynArray[T]) PopBack() T {
	self.init()
	v := (*self.data)[len(*self.data)-1]
	*self.data = (*self.data)[:len(*self.data)-1]
	return v
}

func (self *DynArray[T]) PopFront() T {
	self.init()
	v := (*self.data)[0]
	*self.data = (*self.data)[1:]
	return v
}

func (self DynArray[T]) Back() T {
	self.init()
	return (*self.data)[len(*self.data)-1]
}

func (self DynArray[T]) Front() T {
	self.init()
	return (*self.data)[0]
}

func (self DynArray[T]) Clone() DynArray[T] {
	self.init()
	return DynArray[T]{data: stlbasic.Clone(self.data)}
}

func (self *DynArray[T]) Clear() {
	self.init()
	*self.data = make([]T, 0)
}

func (self DynArray[T]) Empty() bool {
	self.init()
	return self.Length() == 0
}

func (self DynArray[T]) Iterator() iterator.Iterator[T] {
	self.init()
	return iterator.NewIterator[T](_NewIterator[T](&self))
}
