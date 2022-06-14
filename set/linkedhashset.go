package set

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/list"

	"github.com/kkkunny/stl/table"
)

// 有序哈希集合
type LinkedHashSet[T comparable] struct {
	data *table.LinkedHashMap[T, struct{}]
}

// 新建有序哈希集合
func NewLinkedHashSet[T comparable](e ...T) *LinkedHashSet[T] {
	set := &LinkedHashSet[T]{data: table.NewLinkedHashMap[T, struct{}]()}
	for _, i := range e {
		set.Add(i)
	}
	return set
}

// 转成字符串 O(N)
func (self *LinkedHashSet[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	for iter := self.data.Begin(); iter.HasValue(); iter.Next() {
		buf.WriteString(fmt.Sprintf("%v", iter.Key()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度 O(1)
func (self *LinkedHashSet[T]) Length() int {
	return self.data.Length()
}

// 是否为空 O(1)
func (self *LinkedHashSet[T]) Empty() bool {
	return self.data.Empty()
}

// 增加元素 O(1)
func (self *LinkedHashSet[T]) Add(e T) bool {
	if self.data.ContainKey(e) {
		return false
	}
	self.data.Set(e, struct{}{})
	return true
}

// 是否包含元素 O(1)
func (self *LinkedHashSet[T]) Contain(e T) bool {
	return self.data.ContainKey(e)
}

// 删除元素 O(1)
func (self *LinkedHashSet[T]) Remove(e T) bool {
	if !self.data.ContainKey(e) {
		return false
	}
	self.data.Remove(e)
	return true
}

// 清空 O(1)
func (self *LinkedHashSet[T]) Clear() {
	if self.Empty() {
		return
	}
	self.data.Clear()
}

// 克隆 O(N)
func (self *LinkedHashSet[T]) Clone() *LinkedHashSet[T] {
	return &LinkedHashSet[T]{data: self.data.Clone()}
}

// 过滤 O(N)
func (self *LinkedHashSet[T]) Filter(f func(i int, v T) bool) *LinkedHashSet[T] {
	return &LinkedHashSet[T]{data: self.data.Filter(func(i int, k T, _ struct{}) bool {
		return f(i, k)
	})}
}

// 获取起始迭代器
func (self *LinkedHashSet[T]) Begin() *list.ArrayListIterator[T] {
	al := list.NewArrayList[T](self.data.Length(), self.data.Length())
	for iter := self.data.Begin(); iter.HasValue(); iter.Next() {
		al.Set(iter.Index(), iter.Key())
	}
	return al.Begin()
}

// 获取结束迭代器
func (self *LinkedHashSet[T]) End() *list.ArrayListIterator[T] {
	al := list.NewArrayList[T](self.data.Length(), self.data.Length())
	for iter := self.data.Begin(); iter.HasValue(); iter.Next() {
		al.Set(iter.Index(), iter.Key())
	}
	return al.End()
}

// First 获取第一个值 O(1)
func (self *LinkedHashSet[T]) First() T {
	k, _ := self.data.First()
	return k
}

// Last 获取最后一个值 O(1)
func (self *LinkedHashSet[T]) Last() T {
	k, _ := self.data.Last()
	return k
}
