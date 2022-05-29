package list

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// 节点
type doubleLinkedListNode[T any] struct {
	elem T
	prev *doubleLinkedListNode[T]
	next *doubleLinkedListNode[T]
}

// 双向链表
type DoubleLinkedList[T any] struct {
	head   *doubleLinkedListNode[T]
	tail   *doubleLinkedListNode[T]
	length int
}

// 新建双向链表
func NewDoubleLinkedList[T any](e ...T) *DoubleLinkedList[T] {
	ll := &DoubleLinkedList[T]{}
	ll.Add(e...)
	return ll
}

// 检查越界
func (self *DoubleLinkedList[T]) checkOut(i int) {
	length := self.Length()
	if i >= length {
		panic(fmt.Sprintf("index out of range [%d] with length %d", i, length))
	}
}

// 根据索引获取节点
func (self *DoubleLinkedList[T]) getNodeByIndex(i int) *doubleLinkedListNode[T] {
	self.checkOut(i)
	if i <= self.length/2 {
		for cursor := self.head; cursor != nil; cursor = cursor.next {
			if i == 0 {
				return cursor
			}
			i--
		}
	} else {
		i = self.length - i - 1
		for cursor := self.tail; cursor != nil; cursor = cursor.prev {
			if i == 0 {
				return cursor
			}
			i--
		}
	}
	return nil
}

// 转成字符串 O(N)
func (self *DoubleLinkedList[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('[')
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		buf.WriteString(fmt.Sprintf("%v", cursor.elem))
		if cursor.next != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

// json序列化
func (self *DoubleLinkedList[T]) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if err := buf.WriteByte('['); err != nil {
		return nil, err
	}
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		data, err := json.Marshal(cursor.elem)
		if err != nil {
			return nil, err
		}
		if _, err = buf.Write(data); err != nil {
			return nil, err
		}
		if cursor.next != nil {
			if err = buf.WriteByte(','); err != nil {
				return nil, err
			}
		}
	}
	if err := buf.WriteByte(']'); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// json反序列化
func (self *DoubleLinkedList[T]) UnmarshalJSON(data []byte) error {
	var tmp []T
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	self.Clear()
	for _, e := range tmp {
		self.PushBack(e)
	}
	return nil
}

// 获取长度 O(1)
func (self *DoubleLinkedList[T]) Length() int {
	return self.length
}

// 是否为空 O(1)
func (self *DoubleLinkedList[T]) Empty() bool {
	return self.head == nil
}

// 增加元素 O(1)
func (self *DoubleLinkedList[T]) Add(e ...T) {
	self.PushBack(e...)
}

// 增加到头部 O(1)
func (self *DoubleLinkedList[T]) PushFront(e ...T) {
	if len(e) == 0 {
		return
	}
	var begin, end *doubleLinkedListNode[T]
	for _, elem := range e {
		node := &doubleLinkedListNode[T]{elem: elem}
		if begin == nil {
			begin, end = node, node
		} else {
			end.next = node
			node.prev = end
			end = node
		}
	}
	if self.head == nil {
		self.head, self.tail = begin, end
	} else {
		end.next = self.head
		self.head.prev = end
		self.head = begin
	}
	self.length += len(e)
}

// 增加到尾部 O(1)
func (self *DoubleLinkedList[T]) PushBack(e ...T) {
	for _, i := range e {
		node := &doubleLinkedListNode[T]{elem: i}
		if self.Empty() {
			self.head, self.tail = node, node
		} else {
			self.tail.next = node
			node.prev = self.tail
			self.tail = node
		}
	}
	self.length += len(e)
}

// 插入元素 O(N)
func (self *DoubleLinkedList[T]) Insert(i int, e ...T) {
	cursor := self.getNodeByIndex(i)
	if cursor.prev == nil { // 头部
		self.PushFront(e...)
	} else {
		for _, i := range e {
			node := &doubleLinkedListNode[T]{elem: i}
			cursor.prev.next = node
			node.prev = cursor.prev
			cursor.prev = node
			node.next = cursor
		}
		self.length += len(e)
	}
}

// 移除节点
func (self *DoubleLinkedList[T]) removeNode(node *doubleLinkedListNode[T]) {
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
	self.length--
}

// 移除元素 O(N)
func (self *DoubleLinkedList[T]) Remove(i int) T {
	cursor := self.getNodeByIndex(i)
	self.removeNode(cursor)
	return cursor.elem
}

// 删除头元素 O(1)
func (self *DoubleLinkedList[T]) PopFront() T {
	self.checkOut(0)
	elem := self.head.elem
	self.removeNode(self.head)
	return elem
}

// 删除尾元素 O(1)
func (self *DoubleLinkedList[T]) PopBack() T {
	self.checkOut(0)
	elem := self.tail.elem
	self.removeNode(self.tail)
	return elem
}

// 获取元素 O(N)
func (self *DoubleLinkedList[T]) Get(i int) T {
	cursor := self.getNodeByIndex(i)
	return cursor.elem
}

// 获取第一个元素 O(1)
func (self *DoubleLinkedList[T]) First() T {
	self.checkOut(0)
	return self.head.elem
}

// 获取最后一个元素 O(1)
func (self *DoubleLinkedList[T]) Last() T {
	self.checkOut(0)
	return self.tail.elem
}

// 设置元素 O(N)
func (self *DoubleLinkedList[T]) Set(i int, e T) T {
	cursor := self.getNodeByIndex(i)
	elem := cursor.elem
	cursor.elem = e
	return elem
}

// 清空 O(1)
func (self *DoubleLinkedList[T]) Clear() {
	if self.Empty() {
		return
	}
	self.head = nil
	self.tail = nil
	self.length = 0
}

// 克隆 O(N)
func (self *DoubleLinkedList[T]) Clone() *DoubleLinkedList[T] {
	dll := NewDoubleLinkedList[T]()
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		dll.PushBack(cursor.elem)
	}
	return dll
}

// 过滤 O(N)
func (self *DoubleLinkedList[T]) Filter(f func(i int, v T) bool) *DoubleLinkedList[T] {
	dll := NewDoubleLinkedList[T]()
	var index int
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		if f(index, cursor.elem) {
			dll.Add(cursor.elem)
		}
		index++
	}
	return dll
}

// 切分[b, e) O(N)
func (self *DoubleLinkedList[T]) Slice(b, e int) *DoubleLinkedList[T] {
	self.checkOut(b)
	dll := NewDoubleLinkedList[T]()
	for cursor := self.getNodeByIndex(b); cursor != nil && b < e; cursor = cursor.next {
		dll.Add(cursor.elem)
		b++
	}
	return dll
}

// 拼接 O(N)
func (self *DoubleLinkedList[T]) Contact(a *DoubleLinkedList[T]) {
	for cursor := a.head; cursor != nil; cursor = cursor.next {
		self.PushBack(cursor.elem)
	}
}

// 是否有任何元素满足条件 O(N)
func (self *DoubleLinkedList[T]) Any(f func(i int, v T) bool) bool {
	var index int
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		if f(index, cursor.elem) {
			return true
		}
		index++
	}
	return false
}

// 是否所有元素都满足条件 O(N)
func (self *DoubleLinkedList[T]) Every(f func(i int, v T) bool) bool {
	var index int
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		if !f(index, cursor.elem) {
			return false
		}
		index++
	}
	return true
}

// 获取起始迭代器
func (self *DoubleLinkedList[T]) Begin() *DoubleLinkedListIterator[T] {
	return &DoubleLinkedListIterator[T]{cursor: self.head}
}

// 获取结束迭代器
func (self *DoubleLinkedList[T]) End() *DoubleLinkedListIterator[T] {
	return &DoubleLinkedListIterator[T]{cursor: self.tail}
}

// 迭代器
type DoubleLinkedListIterator[T any] struct {
	cursor *doubleLinkedListNode[T] // 目前节点
	index  int                      // 索引
}

// 是否存在值
func (self *DoubleLinkedListIterator[T]) HasValue() bool {
	return self.cursor != nil
}

// 是否存在上一个
func (self *DoubleLinkedListIterator[T]) HasPrev() bool {
	return self.cursor.prev != nil
}

// 是否存在下一个
func (self *DoubleLinkedListIterator[T]) HasNext() bool {
	return self.cursor.next != nil
}

// 上一个
func (self *DoubleLinkedListIterator[T]) Prev() {
	self.cursor = self.cursor.prev
	self.index--
}

// 下一个
func (self *DoubleLinkedListIterator[T]) Next() {
	self.cursor = self.cursor.next
	self.index++
}

// 获取索引
func (self *DoubleLinkedListIterator[T]) Index() int {
	return self.index
}

// 获取值
func (self *DoubleLinkedListIterator[T]) Value() T {
	return self.cursor.elem
}
