package set

import (
	"fmt"
	"github.com/kkkunny/stl/table"
	. "github.com/kkkunny/stl/types"
	"strings"
)

// 哈希集合
type HashSet[T Hasher] struct {
	data *table.HashMap[T, struct{}]
}

// 新建哈希集合
func NewHashSet[T Hasher](e ...T) *HashSet[T] {
	set := &HashSet[T]{data: table.NewHashMap[T, struct{}]()}
	for _, i := range e {
		set.Add(i)
	}
	return set
}

// 转成字符串
func (self *HashSet[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	length := self.data.Length()
	var index Usize
	for iter := self.data.Iterator(); iter.HasValue(); iter.Next() {
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
func (self *HashSet[T]) Length() Usize {
	return self.data.Length()
}

// 是否为空
func (self *HashSet[T]) Empty() bool {
	return self.data.Empty()
}

// 增加元素
func (self *HashSet[T]) Add(e T) bool {
	if self.data.ContainKey(e) {
		return false
	}
	self.data.Set(e, struct{}{})
	return true
}

// 是否包含元素
func (self *HashSet[T]) Contain(e T) bool {
	return self.data.ContainKey(e)
}

// 删除元素
func (self *HashSet[T]) Remove(e T) bool {
	if !self.data.ContainKey(e) {
		return false
	}
	self.data.Remove(e, struct{}{})
	return true
}

// 清空
func (self *HashSet[T]) Clear() {
	self.data.Clear()
}

// 克隆
func (self *HashSet[T]) Clone() *HashSet[T] {
	return &HashSet[T]{data: self.data.Clone()}
}

// 获取迭代器
func (self *HashSet[T]) Iterator() *HashSetIterator[T] {
	return &HashSetIterator[T]{data: self.data.Iterator()}
}

// 迭代器
type HashSetIterator[T Hasher] struct {
	data *table.HashMapIterator[T, struct{}]
}

// 是否存在下一个
func (self *HashSetIterator[T]) HasNext() bool {
	return self.data.HasValue()
}

// 下一个
func (self *HashSetIterator[T]) Next() {
	self.data.Next()
}

// 获取值
func (self *HashSetIterator[T]) Value() T {
	k, _ := self.data.Value()
	return k
}
