package list

import (
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

// 转成字符串
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

// 获取长度
func (self *LinkedList[T]) Length() Usize {
	return self.length
}

// 是否为空
func (self *LinkedList[T]) Empty() bool {
	return self.head == nil
}

// 增加元素
func (self *LinkedList[T]) Add(e ...T) {
	self.PushBack(e...)
}

// 增加到头部
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

// 增加到尾部
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

// 插入元素
func (self *LinkedList[T]) Insert(i Usize, e ...T) {
	self.checkOut(i)
	cursor := self.head
	for index := Usize(0); cursor != nil; cursor = cursor.next {
		if index == i {
			break
		}
		index++
	}
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

// 移除元素
func (self *LinkedList[T]) Remove(i Usize) T {
	self.checkOut(i)
	cursor := self.head
	for index := Usize(0); cursor != nil; cursor = cursor.next {
		if index == i {
			break
		}
		index++
	}
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

// 删除头节点
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

// 删除尾节点
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

// 获取元素
func (self *LinkedList[T]) Get(i Usize) T {
	self.checkOut(i)
	cursor := self.head
	for index := Usize(0); cursor != nil; cursor = cursor.next {
		if index == i {
			break
		}
		index++
	}
	return cursor.elem
}

// 获取第一个节点
func (self *LinkedList[T]) First() T {
	self.checkOut(0)
	return self.head.elem
}

// 获取最后一个节点
func (self *LinkedList[T]) Last() T {
	self.checkOut(0)
	return self.tail.elem
}

// 设置元素
func (self *LinkedList[T]) Set(i Usize, e T) T {
	self.checkOut(i)
	cursor := self.head
	for index := Usize(0); cursor != nil; cursor = cursor.next {
		if index == i {
			break
		}
		index++
	}
	elem := cursor.elem
	cursor.elem = e
	return elem
}

// 清空
func (self *LinkedList[T]) Clear() {
	self.head = nil
	self.tail = nil
	self.length = 0
}

// 克隆
func (self *LinkedList[T]) Clone() *LinkedList[T] {
	newList := NewLinkedList[T]()
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		newList.PushBack(cursor.elem)
	}
	return newList
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
}

// 是否存在值
func (self *LinkedListIterator[T]) HasValue() bool {
	return self.cursor != nil
}

// 上一个
func (self *LinkedListIterator[T]) Prev() {
	self.cursor = self.cursor.prev
}

// 下一个
func (self *LinkedListIterator[T]) Next() {
	self.cursor = self.cursor.next
}

// 获取值
func (self *LinkedListIterator[T]) Value() T {
	return self.cursor.elem
}
