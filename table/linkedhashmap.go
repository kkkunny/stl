package table

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/list"
	. "github.com/kkkunny/stl/types"
)

// 节点
type linkedHashMapNode[K Hasher, V any] struct {
	key  K
	prev *linkedHashMapNode[K, V]
	next *linkedHashMapNode[K, V]
}

// 节点
type linkedHashMapEntry[K Hasher, V any] struct {
	value V
	node  *linkedHashMapNode[K, V]
}

// 有序哈希表
type LinkedHashMap[K Hasher, V any] struct {
	head *linkedHashMapNode[K, V]
	tail *linkedHashMapNode[K, V]
	data *HashMap[K, *linkedHashMapEntry[K, V]]
}

// 新建有序哈希表
func NewLinkedHashMap[K Hasher, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		data: NewHashMap[K, *linkedHashMapEntry[K, V]](),
	}
}

// 转成字符串 O(N)
func (self *LinkedHashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var index Usize
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		buf.WriteString(fmt.Sprintf("%v: %v", iter.Key(), iter.Value()))
		if index < self.data.length-1 {
			buf.WriteString(", ")
		}
		index++
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度 O(1)
func (self *LinkedHashMap[K, V]) Length() Usize {
	return self.data.length
}

// 是否为空 O(1)
func (self *LinkedHashMap[K, V]) Empty() bool {
	return self.data.length == 0
}

// 设置键值对 O(1)-O(N)
func (self *LinkedHashMap[K, V]) Set(k K, v V) {
	if !self.ContainKey(k) {
		value := &linkedHashMapEntry[K, V]{
			value: v,
			node:  &linkedHashMapNode[K, V]{key: k},
		}
		if self.head == nil {
			self.head, self.tail = value.node, value.node
		} else {
			self.tail.next = value.node
			value.node.prev = self.tail
			self.tail = value.node
		}
		self.data.Set(k, value)
	} else {
		value := self.data.Get(k, nil)
		value.value = v
	}
}

// 获取值 O(1)
func (self *LinkedHashMap[K, V]) Get(k K, v ...V) V {
	value := self.data.Get(k, nil)
	if value != nil {
		return value.value
	}
	if len(v) == 0 {
		var v V
		return v
	}
	return v[0]
}

// 检查越界
func (self *LinkedHashMap[K, V]) checkOut(i Usize) {
	length := self.Length()
	if i >= length {
		panic(fmt.Sprintf("index out of range [%d] with length %d", i, length))
	}
}

// 根据下标获取值 O(N)
func (self *LinkedHashMap[K, V]) GetByIndex(i Usize) (K, V) {
	self.checkOut(i)
	cursor := self.head
	for index := Usize(0); index < i; index++ {
		cursor = cursor.next
	}
	key := cursor.key
	return key, self.data.Get(key, nil).value
}

// 移除链表节点
func (self *LinkedHashMap[K, V]) removeNode(node *linkedHashMapNode[K, V]) {
	if node.prev == nil && node.next == nil {
		self.head, self.tail = nil, nil
	} else if node.prev == nil {
		node.next.prev = nil
		self.head = node.next
	} else if node.next == nil {
		node.prev.next = nil
		self.tail = node.prev
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.prev, node.next = nil, nil
}

// 移除键值对 O(1)
func (self *LinkedHashMap[K, V]) Remove(k K, v ...V) V {
	value := self.data.Remove(k, nil)
	if value != nil {
		self.removeNode(value.node)
		return value.value
	}
	if len(v) == 0 {
		var v V
		return v
	}
	return v[0]
}

// 根据下标移除键值对 O(N)
func (self *LinkedHashMap[K, V]) RemoveById(i Usize) (K, V) {
	self.checkOut(i)
	cursor := self.head
	for index := Usize(0); index < i; index++ {
		cursor = cursor.next
	}
	key := cursor.key
	self.removeNode(cursor)
	return key, self.data.Remove(key, nil).value
}

// 是否存在键 O(1)
func (self *LinkedHashMap[K, V]) ContainKey(k K) bool {
	return self.data.ContainKey(k)
}

// 获取键 O(N)
func (self *LinkedHashMap[K, V]) Keys() *list.ArrayList[K] {
	keys := list.NewArrayList[K](self.data.length, self.data.length)
	var index Usize
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		keys.Set(index, cursor.key)
		index++
	}
	return keys
}

// 获取值 O(N)
func (self *LinkedHashMap[K, V]) Values() *list.ArrayList[V] {
	values := list.NewArrayList[V](self.data.length, self.data.length)
	var index Usize
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		values.Set(index, self.data.Get(cursor.key, nil).value)
		index++
	}
	return values
}

// 清空 O(1)
func (self *LinkedHashMap[K, V]) Clear() {
	self.head, self.tail = nil, nil
	self.data.Clear()
}

// 克隆 O(N)
func (self *LinkedHashMap[K, V]) Clone() *LinkedHashMap[K, V] {
	newlhm := NewLinkedHashMap[K, V]()
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		newlhm.Set(cursor.key, self.data.Get(cursor.key, nil).value)
	}
	return newlhm
}

// 过滤 O(N)
func (self *LinkedHashMap[K, V]) Filter(f func(i Usize, k K, v V) bool) *LinkedHashMap[K, V] {
	lhm := NewLinkedHashMap[K, V]()
	var index Usize
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		key := cursor.key
		value := self.data.Get(cursor.key, nil).value
		if f(index, key, value) {
			lhm.Set(key, value)
		}
		index++
	}
	return lhm
}

// 获取起始迭代器
func (self *LinkedHashMap[K, V]) Begin() *LinkedHashMapIterator[K, V] {
	return &LinkedHashMapIterator[K, V]{
		data:   self,
		cursor: self.head,
	}
}

// 获取结束迭代器
func (self *LinkedHashMap[K, V]) End() *LinkedHashMapIterator[K, V] {
	return &LinkedHashMapIterator[K, V]{
		data:   self,
		cursor: self.tail,
	}
}

// 迭代器
type LinkedHashMapIterator[K Hasher, V any] struct {
	data   *LinkedHashMap[K, V]
	cursor *linkedHashMapNode[K, V] // 目前节点
	index  Usize                    // 下标
}

// 是否存在下一个
func (self *LinkedHashMapIterator[K, V]) HasValue() bool {
	return self.cursor != nil
}

// 上一个
func (self *LinkedHashMapIterator[K, V]) Prev() {
	self.cursor = self.cursor.prev
}

// 下一个
func (self *LinkedHashMapIterator[K, V]) Next() {
	self.cursor = self.cursor.next
	self.index++
}

// 获取下标
func (self *LinkedHashMapIterator[K, V]) Index() Usize {
	return self.index
}

// 获取键
func (self *LinkedHashMapIterator[K, V]) Key() K {
	return self.cursor.key
}

// 获取值
func (self *LinkedHashMapIterator[K, V]) Value() V {
	return self.data.data.Get(self.cursor.key, nil).value
}
