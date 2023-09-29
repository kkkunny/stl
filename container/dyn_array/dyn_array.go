package dynarray

import (
	"fmt"
	"reflect"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
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

	var tmp T
	if _, ok := any(tmp).(stlbasic.Comparable); ok {
		for i, v := range *self.data {
			if eqv, ok := any(v).(stlbasic.Comparable); ok {
				if !eqv.Equal((*self.data)[i]) {
					return false
				}
			}
		}
	} else if tp := reflect.TypeOf(tmp); tp.Comparable() {
		for i, v := range *self.data {
			if !reflect.ValueOf(v).Equal(reflect.ValueOf((*self.data)[i])) {
				return false
			}
		}
	} else {
		panic(fmt.Errorf("type `%s` cannot be compared", tp))
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

func (self *DynArray[T]) String() string {
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
