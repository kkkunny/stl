package linkedhashset

import (
	"fmt"
	"strings"

	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/linkedhashmap"
)

type _AnyLinkedHashSet[T any] struct {
	data linkedhashmap.LinkedHashMap[T, struct{}]
}

func _NewAnyLinkedHashSet[T any]() LinkedHashSet[T] {
	return &_AnyLinkedHashSet[T]{data: linkedhashmap.AnyWith[T, struct{}]()}
}

func _NewAnyLinkedHashSetWithCapacity[T any](cap uint) LinkedHashSet[T] {
	return &_AnyLinkedHashSet[T]{data: linkedhashmap.AnyWithCap[T, struct{}](cap)}
}

func _NewAnyLinkedHashSetWith[T any](vs ...T) LinkedHashSet[T] {
	self := _NewAnyLinkedHashSetWithCapacity[T](uint(len(vs)))
	for _, v := range vs {
		self.Add(v)
	}
	return self
}

func (self *_AnyLinkedHashSet[T]) Default() LinkedHashSet[T] {
	return _NewAnyLinkedHashSet[T]()
}

func (self *_AnyLinkedHashSet[T]) Clone() LinkedHashSet[T] {
	return &_AnyLinkedHashSet[T]{data: self.data.Clone()}
}

func (self *_AnyLinkedHashSet[T]) Equal(dst LinkedHashSet[T]) bool {
	return self.data.Equal(dst.getData())
}

func (_ *_AnyLinkedHashSet[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	self := _NewAnyLinkedHashSetWithCapacity[T](iter.Length())
	for iter.Next() {
		self.Add(iter.Value())
	}
	return self
}

func (self *_AnyLinkedHashSet[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(self.data.Keys()...)
}

func (self *_AnyLinkedHashSet[T]) Length() uint {
	return self.data.Length()
}

// Add 插入值
func (self *_AnyLinkedHashSet[T]) Add(v T) bool {
	exist := self.data.ContainKey(v)
	if exist {
		return false
	}
	self.data.Set(v, struct{}{})
	return true
}

// Contain 是否包含值
func (self *_AnyLinkedHashSet[T]) Contain(v T) bool {
	return self.data.ContainKey(v)
}

// Remove 移除值
func (self *_AnyLinkedHashSet[T]) Remove(v T) bool {
	exist := self.data.ContainKey(v)
	if !exist {
		return false
	}
	self.data.Remove(v)
	return true
}

// Clear 清空
func (self *_AnyLinkedHashSet[T]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_AnyLinkedHashSet[T]) Empty() bool {
	return self.data.Empty()
}

// ToSlice 转成切片
func (self *_AnyLinkedHashSet[T]) ToSlice() []T {
	return self.data.Keys()
}

func (self *_AnyLinkedHashSet[T]) String() string {
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

func (self *_AnyLinkedHashSet[T]) getData() linkedhashmap.LinkedHashMap[T, struct{}] {
	return self.data
}
