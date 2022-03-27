package set

import (
	"fmt"
	"stl/table"
	. "stl/types"
	"strings"
)

// 有序集合
type TreeSet[T Comparator[T]] struct {
	data *table.TreeMap[T, any]
}

// 新建有序集合
func NewTreeSet[T Comparator[T]](e ...T) *TreeSet[T] {
	set := &TreeSet[T]{data: table.NewTreeMap[T, any]()}
	for _, i := range e {
		set.Add(i)
	}
	return set
}

// 转成字符串
func (self *TreeSet[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	length := self.data.Length()
	var index Usize
	for iter := self.data.Begin(); iter.HasValue(); iter.Next() {
		k, _ := iter.Value()
		buf.WriteString(fmt.Sprintf("%v", k))
		if index < length-1 {
			buf.WriteString(", ")
		}
		index++
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度
func (self *TreeSet[T]) Length() Usize {
	return self.data.Length()
}

// 是否为空
func (self *TreeSet[T]) Empty() bool {
	return self.data.Empty()
}

// 增加元素
func (self *TreeSet[T]) Add(e T) bool {
	if self.data.ContainKey(e) {
		return false
	}
	self.data.Set(e, struct{}{})
	return true
}

// 是否包含元素
func (self *TreeSet[T]) Contain(e T) bool {
	return self.data.ContainKey(e)
}

// 删除元素
func (self *TreeSet[T]) Remove(e T) bool {
	if !self.data.ContainKey(e) {
		return false
	}
	self.data.Remove(e, struct{}{})
	return true
}

// 清空
func (self *TreeSet[T]) Clear() {
	self.data.Clear()
}

// 克隆
func (self *TreeSet[T]) Clone() *TreeSet[T] {
	return &TreeSet[T]{data: self.data.Clone()}
}

// 获取起始迭代器
func (self *TreeSet[T]) Begin() *TreeSetIterator[T] {
	return &TreeSetIterator[T]{data: self.data.Begin()}
}

// 获取结束迭代器
func (self *TreeSet[T]) End() *TreeSetIterator[T] {
	return &TreeSetIterator[T]{data: self.data.End()}
}

// 迭代器
type TreeSetIterator[T Comparator[T]] struct {
	data *table.TreeMapIterator[T, any]
}

// 是否存在值
func (self *TreeSetIterator[T]) HasValue() bool {
	return self.data.HasValue()
}

// 上一个
func (self *TreeSetIterator[T]) Prev() {
	self.data.Prev()
}

// 下一个
func (self *TreeSetIterator[T]) Next() {
	self.data.Next()
}

// 获取值
func (self *TreeSetIterator[T]) Value() T {
	k, _ := self.data.Value()
	return k
}
