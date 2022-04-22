package table

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/list"
)

// 节点
type linkedHashMapEntry[K comparable, V any] struct {
	Entry[K, V]
	prev *linkedHashMapEntry[K, V]
	next *linkedHashMapEntry[K, V]
}

// 有序哈希表
type LinkedHashMap[K comparable, V any] struct {
	data map[K]*linkedHashMapEntry[K, V]
	head *linkedHashMapEntry[K, V]
	tail *linkedHashMapEntry[K, V]
}

// 新建有序哈希表
func NewLinkedHashMap[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{data: make(map[K]*linkedHashMapEntry[K, V])}
}

// 转成字符串 O(N)
func (self *LinkedHashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		buf.WriteString(fmt.Sprintf("%v: %v", cursor.Key, cursor.Value))
		if cursor.next != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度 O(1)
func (self *LinkedHashMap[K, V]) Length() int {
	return len(self.data)
}

// 是否为空 O(1)
func (self *LinkedHashMap[K, V]) Empty() bool {
	return len(self.data) == 0
}

// 设置键值对 O(1)
func (self *LinkedHashMap[K, V]) Set(k K, v V) {
	if _, ok := self.data[k]; ok {
		return
	}
	node := &linkedHashMapEntry[K, V]{
		Entry: Entry[K, V]{
			Key:   k,
			Value: v,
		},
	}
	self.data[k] = node
	if self.tail == nil {
		self.head, self.tail = node, node
	} else {
		self.tail.next = node
		node.prev = self.tail
		self.tail = node
	}
}

// 获取值 O(1)
func (self *LinkedHashMap[K, V]) Get(k K, v ...V) V {
	if node, ok := self.data[k]; ok {
		return node.Value
	} else if len(v) == 0 {
		var vv V
		return vv
	} else {
		return v[0]
	}
}

// 检查越界
func (self *LinkedHashMap[K, V]) checkOut(i int) {
	length := self.Length()
	if i < 0 || i >= length {
		panic(fmt.Sprintf("index out of range [%d] with length %d", i, length))
	}
}

// 根据下标获取值 O(N)
func (self *LinkedHashMap[K, V]) GetByIndex(i int) (K, V) {
	self.checkOut(i)
	cursor := self.head
	for index := 0; cursor != nil; cursor = cursor.next {
		if index == i {
			break
		}
		index++
	}
	return cursor.Key, cursor.Value
}

// 移除链表节点
func (self *LinkedHashMap[K, V]) removeNode(node *linkedHashMapEntry[K, V]) {
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
	if node, ok := self.data[k]; ok {
		delete(self.data, k)
		self.removeNode(node)
		return node.Value
	} else if len(v) == 0 {
		var vv V
		return vv
	} else {
		return v[0]
	}
}

// 根据下标移除键值对 O(N)
func (self *LinkedHashMap[K, V]) RemoveByIndex(i int) (K, V) {
	self.checkOut(i)
	cursor := self.head
	for index := 0; cursor != nil; cursor = cursor.next {
		if index == i {
			break
		}
		index++
	}
	delete(self.data, cursor.Key)
	self.removeNode(cursor)
	return cursor.Key, cursor.Value
}

// 是否存在键 O(1)
func (self *LinkedHashMap[K, V]) ContainKey(k K) bool {
	_, ok := self.data[k]
	return ok
}

// 获取键 O(N)
func (self *LinkedHashMap[K, V]) Keys() *list.ArrayList[K] {
	keys := list.NewArrayList[K](len(self.data), len(self.data))
	var index int
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		keys.Set(index, cursor.Key)
		index++
	}
	return keys
}

// 获取值 O(N)
func (self *LinkedHashMap[K, V]) Values() *list.ArrayList[V] {
	values := list.NewArrayList[V](len(self.data), len(self.data))
	var index int
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		values.Set(index, cursor.Value)
		index++
	}
	return values
}

// 清空 O(1)
func (self *LinkedHashMap[K, V]) Clear() {
	if self.Empty() {
		return
	}
	self.head, self.tail = nil, nil
	self.data = make(map[K]*linkedHashMapEntry[K, V])
}

// 克隆 O(N)
func (self *LinkedHashMap[K, V]) Clone() *LinkedHashMap[K, V] {
	lhm := NewLinkedHashMap[K, V]()
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		lhm.Set(cursor.Key, cursor.Value)
	}
	return lhm
}

// 过滤 O(N)
func (self *LinkedHashMap[K, V]) Filter(f func(i int, k K, v V) bool) *LinkedHashMap[K, V] {
	lhm := NewLinkedHashMap[K, V]()
	var index int
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		if f(index, cursor.Key, cursor.Value) {
			lhm.Set(cursor.Key, cursor.Value)
		}
		index++
	}
	return lhm
}

// 获取起始迭代器
func (self *LinkedHashMap[K, V]) Begin() *LinkedHashMapIterator[K, V] {
	return &LinkedHashMapIterator[K, V]{cursor: self.head}
}

// 获取结束迭代器
func (self *LinkedHashMap[K, V]) End() *LinkedHashMapIterator[K, V] {
	return &LinkedHashMapIterator[K, V]{
		cursor: self.tail,
		index:  len(self.data) - 1,
	}
}

// 迭代器
type LinkedHashMapIterator[K comparable, V any] struct {
	cursor *linkedHashMapEntry[K, V] // 目前节点
	index  int                       // 下标
}

// 是否存在值
func (self *LinkedHashMapIterator[K, V]) HasValue() bool {
	return self.cursor != nil
}

// 是否存在上一个
func (self *LinkedHashMapIterator[K, V]) HasPrev() bool {
	return self.cursor.prev != nil
}

// 是否存在下一个
func (self *LinkedHashMapIterator[K, V]) HasNext() bool {
	return self.cursor.next != nil
}

// 上一个
func (self *LinkedHashMapIterator[K, V]) Prev() {
	if self.HasPrev() {
		self.cursor = self.cursor.prev
		self.index--
	}
}

// 下一个
func (self *LinkedHashMapIterator[K, V]) Next() {
	if self.HasNext() {
		self.cursor = self.cursor.next
		self.index++
	}
}

// 获取下标
func (self *LinkedHashMapIterator[K, V]) Index() int {
	return self.index
}

// 获取键
func (self *LinkedHashMapIterator[K, V]) Key() K {
	return self.cursor.Key
}

// 获取值
func (self *LinkedHashMapIterator[K, V]) Value() V {
	return self.cursor.Value
}
