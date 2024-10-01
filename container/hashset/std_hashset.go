package hashset

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/container/hashmap"
	stliter "github.com/kkkunny/stl/container/iter"
)

type _StdHashSet[T comparable] struct {
	data hashmap.HashMap[T, struct{}]
}

func _NewStdHashSet[T comparable]() HashSet[T] {
	return &_StdHashSet[T]{data: hashmap.StdWith[T, struct{}]()}
}

func _NewStdHashSetWithCapacity[T comparable](cap uint) HashSet[T] {
	return &_StdHashSet[T]{data: hashmap.StdWithCap[T, struct{}](cap)}
}

func _NewStdHashSetWith[T comparable](vs ...T) HashSet[T] {
	self := _NewStdHashSetWithCapacity[T](uint(len(vs)))
	for _, v := range vs {
		self.Add(v)
	}
	return self
}

func (self *_StdHashSet[T]) Clone() HashSet[T] {
	return &_StdHashSet[T]{data: self.data.Clone()}
}

func (self *_StdHashSet[T]) Equal(dst HashSet[T]) bool {
	return self.data.Equal(dst.getData())
}

func (_ *_StdHashSet[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	self := _NewAnyHashSetWithCapacity[T](iter.Length())
	for iter.Next() {
		self.Add(iter.Value())
	}
	return self
}

func (self *_StdHashSet[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(self.data.Keys()...)
}

func (self *_StdHashSet[T]) Length() uint {
	return self.data.Length()
}

// Add 插入值
func (self *_StdHashSet[T]) Add(v T) bool {
	exist := self.data.ContainKey(v)
	if exist {
		return false
	}
	self.data.Set(v, struct{}{})
	return true
}

// Contain 是否包含值
func (self *_StdHashSet[T]) Contain(v T) bool {
	return self.data.ContainKey(v)
}

// Remove 移除值
func (self *_StdHashSet[T]) Remove(v T) bool {
	exist := self.data.ContainKey(v)
	if !exist {
		return false
	}
	self.data.Remove(v)
	return true
}

// Clear 清空
func (self *_StdHashSet[T]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_StdHashSet[T]) Empty() bool {
	return self.data.Empty()
}

// ToSlice 转成切片
func (self *_StdHashSet[T]) ToSlice() []T {
	return self.data.Keys()
}

func (self *_StdHashSet[T]) String() string {
	var buf strings.Builder
	buf.WriteString("HashSet{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_StdHashSet[T]) getData() hashmap.HashMap[T, struct{}] {
	return self.data
}
