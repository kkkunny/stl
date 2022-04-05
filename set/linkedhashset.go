package set

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/table"
	. "github.com/kkkunny/stl/types"
)

// 有序哈希集合
type LinkedHashSet[T Hasher] struct {
	data *table.LinkedHashMap[T, struct{}]
}

// 新建有序哈希集合
func NewLinkedHashSet[T Hasher](e ...T) *LinkedHashSet[T] {
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
	length := self.data.Length()
	var index Usize
	for iter := self.data.Begin(); iter.HasValue(); iter.Next() {
		buf.WriteString(fmt.Sprintf("%v", iter.Key()))
		if index < length-1 {
			buf.WriteString(", ")
		}
		index++
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度 O(1)
func (self *LinkedHashSet[T]) Length() Usize {
	return self.data.Length()
}

// 是否为空 O(1)
func (self *LinkedHashSet[T]) Empty() bool {
	return self.data.Empty()
}

// 增加元素 O(1)-O(N)
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
	self.data.Remove(e, struct{}{})
	return true
}

// 清空 O(1)
func (self *LinkedHashSet[T]) Clear() {
	self.data.Clear()
}

// 克隆 O(N)
func (self *LinkedHashSet[T]) Clone() *LinkedHashSet[T] {
	return &LinkedHashSet[T]{data: self.data.Clone()}
}

// 过滤 O(N)
func (self *LinkedHashSet[T]) Filter(f func(i Usize, v T) bool) *LinkedHashSet[T] {
	lhs := NewLinkedHashSet[T]()
	var index Usize
	for iter := self.data.Begin(); iter.HasValue(); iter.Next() {
		v := iter.Key()
		if f(index, v) {
			lhs.Add(v)
		}
		index++
	}
	return lhs
}

// 获取起始迭代器
func (self *LinkedHashSet[T]) Begin() *LinkedHashSetIterator[T] {
	return &LinkedHashSetIterator[T]{data: self.data.Begin()}
}

// 获取结束迭代器
func (self *LinkedHashSet[T]) End() *LinkedHashSetIterator[T] {
	return &LinkedHashSetIterator[T]{data: self.data.End()}
}

// 迭代器
type LinkedHashSetIterator[T Hasher] struct {
	data *table.LinkedHashMapIterator[T, struct{}]
}

// 是否存在值
func (self *LinkedHashSetIterator[T]) HasValue() bool {
	return self.data.HasValue()
}

// 上一个
func (self *LinkedHashSetIterator[T]) Prev() {
	self.data.Prev()
}

// 下一个
func (self *LinkedHashSetIterator[T]) Next() {
	self.data.Next()
}

// 获取下标
func (self *LinkedHashSetIterator[T]) Index() Usize {
	return self.data.Index()
}

// 获取值
func (self *LinkedHashSetIterator[T]) Value() T {
	return self.data.Key()
}
