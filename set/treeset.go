package set

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/list"
	"golang.org/x/exp/constraints"

	"github.com/kkkunny/stl/table"
)

// 有序集合
type TreeSet[T constraints.Ordered] struct {
	data *table.TreeMap[T, struct{}]
}

// 新建有序集合
func NewTreeSet[T constraints.Ordered](e ...T) *TreeSet[T] {
	ts := &TreeSet[T]{data: table.NewTreeMap[T, struct{}]()}
	for _, i := range e {
		ts.Add(i)
	}
	return ts
}

// 转成字符串 O(N)
func (self *TreeSet[T]) String() string {
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
func (self *TreeSet[T]) Length() int {
	return self.data.Length()
}

// 是否为空 O(1)
func (self *TreeSet[T]) Empty() bool {
	return self.data.Empty()
}

// 增加元素 O(logN)
func (self *TreeSet[T]) Add(e T) bool {
	if self.data.ContainKey(e) {
		return false
	}
	self.data.Set(e, struct{}{})
	return true
}

// 是否包含元素 O(logN)
func (self *TreeSet[T]) Contain(e T) bool {
	return self.data.ContainKey(e)
}

// 删除元素 O(logN)
func (self *TreeSet[T]) Remove(e T) bool {
	if !self.data.ContainKey(e) {
		return false
	}
	self.data.Remove(e)
	return true
}

// 清空 O(1)
func (self *TreeSet[T]) Clear() {
	self.data.Clear()
}

// 克隆 O(NlogN)
func (self *TreeSet[T]) Clone() *TreeSet[T] {
	return &TreeSet[T]{data: self.data.Clone()}
}

// 过滤 O(N)
func (self *TreeSet[T]) Filter(f func(v T) bool) *TreeSet[T] {
	return &TreeSet[T]{data: self.data.Filter(func(k T, _ struct{}) bool {
		return f(k)
	})}
}

// 获取起始迭代器
func (self *TreeSet[T]) Begin() *list.ArrayListIterator[T] {
	al := list.NewArrayList[T](self.data.Length(), self.data.Length())
	var index int
	for iter := self.data.Begin(); iter.HasValue(); iter.Next() {
		al.Set(index, iter.Key())
		index++
	}
	return al.Begin()
}

// 获取结束迭代器
func (self *TreeSet[T]) End() *list.ArrayListIterator[T] {
	al := list.NewArrayList[T](self.data.Length(), self.data.Length())
	var index int
	for iter := self.data.Begin(); iter.HasValue(); iter.Next() {
		al.Set(index, iter.Key())
		index++
	}
	return al.End()
}
