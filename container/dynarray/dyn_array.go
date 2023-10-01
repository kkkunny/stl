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

func NewDynArray[T any]() DynArray[T] {
	var data []T
	return DynArray[T]{data: &data}
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
	return uint(len(*self.data))
}

func (self DynArray[T]) Capacity() uint {
	return uint(cap(*self.data))
}

func (self DynArray[T]) Equal(dst any) bool {
	da, ok := dst.(DynArray[T])
	if !ok {
		return false
	}

	if self.data == da.data {
		return true
	}

	if self.Length() != da.Length() || self.Capacity() != da.Capacity() {
		return false
	}

	for i, v := range *self.data {
		if !stlbasic.Equal(v, (*da.data)[i]) {
			return false
		}
	}
	return true
}

func (self DynArray[T]) Get(i uint) T {
	return (*self.data)[i]
}

func (self *DynArray[T]) Set(i uint, v T) T {
	pv := (*self.data)[i]
	(*self.data)[i] = v
	return pv
}

func (self *DynArray[T]) PushBack(v T) {
	*self.data = append(*self.data, v)
}

func (self *DynArray[T]) PushFront(v T) {
	*self.data = append([]T{v}, *self.data...)
}

func (self *DynArray[T]) Insert(i uint, v T) {
	*self.data = append(*self.data, v)
	copy((*self.data)[i+1:], (*self.data)[i:])
	(*self.data)[i] = v
}

func (self *DynArray[T]) Remove(i uint) T {
	v := (*self.data)[i]
	copy((*self.data)[i:], (*self.data)[i+1:])
	*self.data = (*self.data)[:len(*self.data)-1]
	return v
}

func (self DynArray[T]) String() string {
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
	v := (*self.data)[len(*self.data)-1]
	*self.data = (*self.data)[:len(*self.data)-1]
	return v
}

func (self *DynArray[T]) PopFront() T {
	v := (*self.data)[0]
	*self.data = (*self.data)[1:]
	return v
}

func (self DynArray[T]) Back() T {
	return (*self.data)[len(*self.data)-1]
}

func (self DynArray[T]) Front() T {
	return (*self.data)[0]
}

func (self DynArray[T]) Clone() any {
	return DynArray[T]{data: stlbasic.Clone(self.data)}
}

func (self *DynArray[T]) Clear() {
	*self.data = make([]T, 0)
}

func (self DynArray[T]) Empty() bool {
	return self.Length() == 0
}

func (self DynArray[T]) Iterator() iterator.Iterator[T] {
	return iterator.NewIterator[DynArray[T], T](_NewIterator[T](&self))
}
