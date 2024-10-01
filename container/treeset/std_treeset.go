package treeset

import (
	"fmt"
	"strings"

	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/treemap"
)

type _StdTreeSet[T any] struct {
	data treemap.TreeMap[T, struct{}]
}

func _NewStdTreeSet[T any]() TreeSet[T] {
	return &_StdTreeSet[T]{data: treemap.AnyWith[T, struct{}]()}
}

func _NewStdTreeSetWith[T any](vs ...T) TreeSet[T] {
	self := _NewStdTreeSet[T]()
	for _, v := range vs {
		self.Add(v)
	}
	return self
}

func (self *_StdTreeSet[T]) Clone() TreeSet[T] {
	return &_StdTreeSet[T]{data: self.data.Clone()}
}

func (self *_StdTreeSet[T]) Equal(dst TreeSet[T]) bool {
	return self.data.Equal(dst.getData())
}

func (_ *_StdTreeSet[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	self := _NewStdTreeSet[T]()
	for iter.Next() {
		self.Add(iter.Value())
	}
	return self
}

func (self *_StdTreeSet[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(self.data.Keys()...)
}

func (self *_StdTreeSet[T]) Length() uint {
	return self.data.Length()
}

// Add 插入值
func (self *_StdTreeSet[T]) Add(v T) bool {
	exist := self.data.ContainKey(v)
	if exist {
		return false
	}
	self.data.Set(v, struct{}{})
	return true
}

// Contain 是否包含值
func (self *_StdTreeSet[T]) Contain(v T) bool {
	return self.data.ContainKey(v)
}

// Remove 移除值
func (self *_StdTreeSet[T]) Remove(v T) bool {
	exist := self.data.ContainKey(v)
	if !exist {
		return false
	}
	self.data.Remove(v)
	return true
}

// Clear 清空
func (self *_StdTreeSet[T]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_StdTreeSet[T]) Empty() bool {
	return self.data.Empty()
}

// ToSlice 转成切片
func (self *_StdTreeSet[T]) ToSlice() []T {
	return self.data.Keys()
}

func (self *_StdTreeSet[T]) String() string {
	var buf strings.Builder
	buf.WriteString("TreeSet{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_StdTreeSet[T]) getData() treemap.TreeMap[T, struct{}] {
	return self.data
}
