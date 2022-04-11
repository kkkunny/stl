package list

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	. "github.com/kkkunny/stl/types"
)

// 节点
type node[T any] struct {
	elem T
	prev *node[T]
	next *node[T]
}

// 链表
type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length Usize
}

// 新建动态数组
func NewLinkedList[T any](e ...T) *LinkedList[T] {
	ll := &LinkedList[T]{}
	ll.Add(e...)
	return ll
}

// 检查越界
func (self *LinkedList[T]) checkOut(i Usize) {
	length := self.Length()
	if i >= length {
		panic(fmt.Sprintf("index out of range [%d] with length %d", i, length))
	}
}

// 根据索引获取节点
func (self *LinkedList[T]) getNodeByIndex(i Usize) *node[T] {
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
func (self *LinkedList[T]) String() string {
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
func (self *LinkedList[T]) MarshalJSON() ([]byte, error) {
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
func (self *LinkedList[T]) UnmarshalJSON(data []byte) error {
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
func (self *LinkedList[T]) Length() Usize {
	return self.length
}

// 是否为空 O(1)
func (self *LinkedList[T]) Empty() bool {
	return self.head == nil
}

// 增加元素 O(1)
func (self *LinkedList[T]) Add(e ...T) {
	self.PushBack(e...)
}

// 增加到头部 O(1)
func (self *LinkedList[T]) PushFront(e ...T) {
	for i, elem := range e {
		if i == 0 {
			node := &node[T]{elem: elem}
			if self.Empty() {
				self.head, self.tail = node, node
			} else {
				node.next = self.head
				self.head.prev = node
				self.head = node
			}
			self.length++
		} else if self.length == Usize(i) {
			self.PushBack(e[i:]...)
			break
		} else {
			self.Insert(1, e[i:]...)
			break
		}
	}
}

// 增加到尾部 O(1)
func (self *LinkedList[T]) PushBack(e ...T) {
	for _, i := range e {
		node := &node[T]{elem: i}
		if self.Empty() {
			self.head, self.tail = node, node
		} else {
			self.tail.next = node
			node.prev = self.tail
			self.tail = node
		}
	}
	self.length += Usize(len(e))
}

// 插入元素 O(N)
func (self *LinkedList[T]) Insert(i Usize, e ...T) {
	cursor := self.getNodeByIndex(i)
	if cursor.prev == nil { // 插入头部
		self.PushFront(e...)
	} else {
		for _, i := range e {
			node := &node[T]{elem: i}
			cursor.prev.next = node
			node.prev = cursor.prev
			cursor.prev = node
			node.next = cursor
		}
		self.length += Usize(len(e))
	}
}

// 移除元素 O(N)
func (self *LinkedList[T]) Remove(i Usize) T {
	cursor := self.getNodeByIndex(i)
	if cursor.prev == nil && cursor.next == nil {
		self.head, self.tail = nil, nil
	} else if cursor.prev == nil { // 头部
		self.head = cursor.next
		self.head.prev = nil
	} else if cursor.next == nil { // 尾部
		self.tail = cursor.prev
		self.tail.next = nil
	} else { // 中间
		cursor.prev.next = cursor.next
		cursor.next.prev = cursor.prev
	}
	cursor.prev, cursor.next = nil, nil
	self.length--
	return cursor.elem
}

// 删除头元素 O(1)
func (self *LinkedList[T]) PopFront() T {
	self.checkOut(0)
	elem := self.head.elem
	if self.length == 1 {
		self.head, self.tail = nil, nil
	} else {
		self.head = self.head.next
		self.head.prev = nil
	}
	self.length--
	return elem
}

// 删除尾元素 O(1)
func (self *LinkedList[T]) PopBack() T {
	self.checkOut(0)
	elem := self.tail.elem
	if self.length == 1 {
		self.head, self.tail = nil, nil
	} else {
		self.tail = self.tail.prev
		self.tail.next = nil
	}
	self.length--
	return elem
}

// 获取元素 O(N)
func (self *LinkedList[T]) Get(i Usize) T {
	cursor := self.getNodeByIndex(i)
	return cursor.elem
}

// 获取第一个元素 O(1)
func (self *LinkedList[T]) First() T {
	self.checkOut(0)
	return self.head.elem
}

// 获取最后一个元素 O(1)
func (self *LinkedList[T]) Last() T {
	self.checkOut(0)
	return self.tail.elem
}

// 设置元素 O(N)
func (self *LinkedList[T]) Set(i Usize, e T) T {
	cursor := self.getNodeByIndex(i)
	elem := cursor.elem
	cursor.elem = e
	return elem
}

// 清空 O(1)
func (self *LinkedList[T]) Clear() {
	self.head = nil
	self.tail = nil
	self.length = 0
}

// 克隆 O(N)
func (self *LinkedList[T]) Clone() *LinkedList[T] {
	newList := NewLinkedList[T]()
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		newList.PushBack(cursor.elem)
	}
	return newList
}

// 过滤 O(N)
func (self *LinkedList[T]) Filter(f func(i Usize, v T) bool) *LinkedList[T] {
	ll := NewLinkedList[T]()
	var index Usize
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		if f(index, cursor.elem) {
			ll.Add(cursor.elem)
		}
		index++
	}
	return ll
}

// 切分[b, e) O(N)
func (self *LinkedList[T]) Slice(b, e Usize) *LinkedList[T] {
	self.checkOut(b)
	tmp := NewLinkedList[T]()
	for cursor := self.getNodeByIndex(b); cursor != nil && b < e; cursor = cursor.next {
		tmp.Add(cursor.elem)
		b++
	}
	return tmp
}

// 获取起始迭代器
func (self *LinkedList[T]) Begin() *LinkedListIterator[T] {
	return &LinkedListIterator[T]{
		data:   self,
		cursor: self.head,
	}
}

// 获取结束迭代器
func (self *LinkedList[T]) End() *LinkedListIterator[T] {
	return &LinkedListIterator[T]{
		data:   self,
		cursor: self.tail,
	}
}

// 迭代器
type LinkedListIterator[T any] struct {
	data   *LinkedList[T] // 列表
	cursor *node[T]       // 目前节点
	index  Usize          // 索引
}

// 是否存在值
func (self *LinkedListIterator[T]) HasValue() bool {
	return self.cursor != nil
}

// 是否存在下一个
func (self *LinkedListIterator[T]) HasNext() bool {
	return self.cursor.next != nil
}

// 上一个
func (self *LinkedListIterator[T]) Prev() {
	self.cursor = self.cursor.prev
}

// 下一个
func (self *LinkedListIterator[T]) Next() {
	self.cursor = self.cursor.next
	self.index++
}

// 获取索引
func (self *LinkedListIterator[T]) Index() Usize {
	return self.index
}

// 获取值
func (self *LinkedListIterator[T]) Value() T {
	return self.cursor.elem
}
