package set

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/list"
)

// 哈希集合
type HashSet[T comparable] struct {
	data map[T]struct{}
}

// 新建哈希集合
func NewHashSet[T comparable](e ...T) *HashSet[T] {
	hs := &HashSet[T]{
		data: make(map[T]struct{}),
	}
	for _, v := range e {
		hs.data[v] = struct{}{}
	}
	return hs
}

// 转成字符串 O(N)
func (self *HashSet[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var index int
	for k := range self.data {
		buf.WriteString(fmt.Sprintf("%v", k))
		if index < len(self.data)-1 {
			buf.WriteString(", ")
		}
		index++
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度 O(1)
func (self *HashSet[T]) Length() int {
	return len(self.data)
}

// 是否为空 O(1)
func (self *HashSet[T]) Empty() bool {
	return len(self.data) == 0
}

// 增加元素 O(1)
func (self *HashSet[T]) Add(e T) bool {
	if _, ok := self.data[e]; ok {
		return false
	}
	self.data[e] = struct{}{}
	return true
}

// 是否包含元素 O(1)
func (self *HashSet[T]) Contain(e T) bool {
	_, ok := self.data[e]
	return ok
}

// 删除元素 O(1)
func (self *HashSet[T]) Remove(e T) bool {
	if _, ok := self.data[e]; ok {
		delete(self.data, e)
		return true
	}
	return false
}

// 清空 O(1)
func (self *HashSet[T]) Clear() {
	if !self.Empty() {
		self.data = make(map[T]struct{})
	}
}

// 克隆 O(N)
func (self *HashSet[T]) Clone() *HashSet[T] {
	hs := NewHashSet[T]()
	for k := range self.data {
		hs.data[k] = struct{}{}
	}
	return hs
}

// 过滤 O(N)
func (self *HashSet[T]) Filter(f func(v T) bool) *HashSet[T] {
	hs := NewHashSet[T]()
	for k := range self.data {
		if f(k) {
			hs.Add(k)
		}
	}
	return hs
}

// 任意一个满足条件 O(N)
func (self *HashSet[T]) Any(f func(v T) bool) bool {
	for k := range self.data {
		if f(k) {
			return true
		}
	}
	return false
}

// 每一个满足条件 O(N)
func (self *HashSet[T]) Every(f func(v T) bool) bool {
	for k := range self.data {
		if !f(k) {
			return false
		}
	}
	return true
}

// 获取迭代器
func (self *HashSet[T]) Iterator() *list.ArrayListIterator[T] {
	al := list.NewArrayList[T](len(self.data), len(self.data))
	var index int
	for k := range self.data {
		al.Set(index, k)
		index++
	}
	return al.Begin()
}
