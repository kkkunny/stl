package linkedhashset

import (
	"fmt"
	"strings"

	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/linkedhashmap"
)

type _StdLinkedHashSet[T comparable] struct {
	data linkedhashmap.LinkedHashMap[T, struct{}]
}

func _NewStdLinkedHashSet[T comparable]() LinkedHashSet[T] {
	return &_StdLinkedHashSet[T]{data: linkedhashmap.AnyWith[T, struct{}]()}
}

func _NewStdLinkedHashSetWithCapacity[T comparable](cap uint) LinkedHashSet[T] {
	return &_StdLinkedHashSet[T]{data: linkedhashmap.AnyWithCap[T, struct{}](cap)}
}

func _NewStdLinkedHashSetWith[T comparable](vs ...T) LinkedHashSet[T] {
	self := _NewStdLinkedHashSetWithCapacity[T](uint(len(vs)))
	for _, v := range vs {
		self.Add(v)
	}
	return self
}

func (self *_StdLinkedHashSet[T]) Clone() LinkedHashSet[T] {
	return &_StdLinkedHashSet[T]{data: self.data.Clone()}
}

func (self *_StdLinkedHashSet[T]) Equal(dst LinkedHashSet[T]) bool {
	return self.data.Equal(dst.getData())
}

func (_ *_StdLinkedHashSet[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	self := _NewStdLinkedHashSetWithCapacity[T](iter.Length())
	for iter.Next() {
		self.Add(iter.Value())
	}
	return self
}

func (self *_StdLinkedHashSet[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(self.data.Keys()...)
}

func (self *_StdLinkedHashSet[T]) Length() uint {
	return self.data.Length()
}

// Add 插入值
func (self *_StdLinkedHashSet[T]) Add(v T) bool {
	exist := self.data.ContainKey(v)
	if exist {
		return false
	}
	self.data.Set(v, struct{}{})
	return true
}

// Contain 是否包含值
func (self *_StdLinkedHashSet[T]) Contain(v T) bool {
	return self.data.ContainKey(v)
}

// Remove 移除值
func (self *_StdLinkedHashSet[T]) Remove(v T) bool {
	exist := self.data.ContainKey(v)
	if !exist {
		return false
	}
	self.data.Remove(v)
	return true
}

// Clear 清空
func (self *_StdLinkedHashSet[T]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_StdLinkedHashSet[T]) Empty() bool {
	return self.data.Empty()
}

// ToSlice 转成切片
func (self *_StdLinkedHashSet[T]) ToSlice() []T {
	return self.data.Keys()
}

func (self *_StdLinkedHashSet[T]) String() string {
	var buf strings.Builder
	buf.WriteString("LinkedHashSet{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_StdLinkedHashSet[T]) getData() linkedhashmap.LinkedHashMap[T, struct{}] {
	return self.data
}
