package set

import (
	"cmp"
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/clone"
	stlcmp "github.com/kkkunny/stl/cmp"
	"github.com/kkkunny/stl/container/hashmap"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/linkedhashmap"
	stlmaps "github.com/kkkunny/stl/container/maps"
	"github.com/kkkunny/stl/container/treemap"
)

type Set[T any] interface {
	setIter[T]
	clone.Cloneable[Set[T]]
	stlcmp.Equalable[Set[T]]
	stliter.IteratorContainer[T]
	stlbasic.Lengthable
	Add(v T) bool
	Contain(v T) bool
	Remove(v T) bool
	Clear()
	Empty() bool
	ToSlice() []T
	fmt.Stringer
	getData() stlmaps.MapObj[T, struct{}]
}

// StdHashSetWith 使用go的默认hash函数，相比map，write更慢，read更快
func StdHashSetWith[T comparable](vs ...T) Set[T] {
	set := _NewSet[T](hashmap.StdWith[T, struct{}]())
	for _, v := range vs {
		set.Add(v)
	}
	return set
}

// StdHashSetWithCap 使用go的默认hash函数，相比map，write更慢，read更快
func StdHashSetWithCap[T comparable](cap uint) Set[T] {
	return _NewSet[T](hashmap.StdWithCap[T, struct{}](cap))
}

// AnyHashSetWith 使用自定义hash函数，相比map，全方位慢
func AnyHashSetWith[T any](vs ...T) Set[T] {
	set := _NewSet[T](hashmap.AnyWith[T, struct{}]())
	for _, v := range vs {
		set.Add(v)
	}
	return set
}

// AnyHashSetWithCap 使用自定义hash函数，相比map，全方位慢
func AnyHashSetWithCap[T any](cap uint) Set[T] {
	return _NewSet[T](hashmap.AnyWithCap[T, struct{}](cap))
}

// StdLinkedHashSetWith 使用go的默认hash函数，相比map，write更慢，read更快
func StdLinkedHashSetWith[T comparable](vs ...T) Set[T] {
	set := _NewSet[T](linkedhashmap.StdWith[T, struct{}]())
	for _, v := range vs {
		set.Add(v)
	}
	return set
}

// StdLinkedHashSetWithCap 使用go的默认hash函数，相比map，write更慢，read更快
func StdLinkedHashSetWithCap[T comparable](cap uint) Set[T] {
	return _NewSet[T](linkedhashmap.StdWithCap[T, struct{}](cap))
}

// AnyLinkedHashSetWith 使用自定义hash函数，相比map，全方位慢
func AnyLinkedHashSetWith[T any](vs ...T) Set[T] {
	set := _NewSet[T](linkedhashmap.AnyWith[T, struct{}]())
	for _, v := range vs {
		set.Add(v)
	}
	return set
}

// AnyLinkedHashSetWithCap 使用自定义hash函数，相比map，全方位慢
func AnyLinkedHashSetWithCap[T any](cap uint) Set[T] {
	return _NewSet[T](linkedhashmap.AnyWithCap[T, struct{}](cap))
}

// StdTreeSetWith 使用go的默认cmp函数
func StdTreeSetWith[T cmp.Ordered](vs ...T) Set[T] {
	set := _NewSet[T](treemap.StdWith[T, struct{}]())
	for _, v := range vs {
		set.Add(v)
	}
	return set
}

// AnyTreeSetWith 使用自定义cmp函数，全方位慢
func AnyTreeSetWith[T any](vs ...T) Set[T] {
	set := _NewSet[T](treemap.AnyWith[T, struct{}]())
	for _, v := range vs {
		set.Add(v)
	}
	return set
}

type _Set[T any] struct {
	data stlmaps.MapObj[T, struct{}]
}

func _NewSet[T any](data stlmaps.MapObj[T, struct{}]) Set[T] {
	return &_Set[T]{data: data}
}

func (self *_Set[T]) Clone() Set[T] {
	return &_Set[T]{data: clone.Clone(self.data)}
}

func (self *_Set[T]) Equal(dstObj Set[T]) bool {
	if dstObj == nil && self == nil {
		return true
	} else if dstObj == nil {
		return false
	}

	dst, ok := dstObj.(Set[T])
	if !ok {
		return false
	}

	return self.data.Equal(dst.getData())
}

func (self *_Set[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	data := clone.Clone(self.data)
	data.Clear()
	newSet := _NewSet[T](data)
	for iter.Next() {
		newSet.Add(iter.Value())
	}
	return newSet
}

func (self *_Set[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(self.data.Keys()...)
}

func (self *_Set[T]) Length() uint {
	return self.data.Length()
}

// Add 插入值
func (self *_Set[T]) Add(v T) bool {
	exist := self.data.Contain(v)
	if exist {
		return false
	}
	self.data.Set(v, struct{}{})
	return true
}

// Contain 是否包含值
func (self *_Set[T]) Contain(v T) bool {
	return self.data.Contain(v)
}

// Remove 移除值
func (self *_Set[T]) Remove(v T) bool {
	exist := self.data.Contain(v)
	if !exist {
		return false
	}
	self.data.Remove(v)
	return true
}

// Clear 清空
func (self *_Set[T]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_Set[T]) Empty() bool {
	return self.data.Empty()
}

// ToSlice 转成切片
func (self *_Set[T]) ToSlice() []T {
	return self.data.Keys()
}

func (self *_Set[T]) String() string {
	var buf strings.Builder
	buf.WriteString("Put{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_Set[T]) getData() stlmaps.MapObj[T, struct{}] {
	return self.data
}
