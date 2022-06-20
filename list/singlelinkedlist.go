package list

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// 节点
type singleLinkedListNode[T any] struct {
	elem T
	next *singleLinkedListNode[T]
}

// 单向链表
type SingleLinkedList[T any] struct {
	head   *singleLinkedListNode[T]
	tail   *singleLinkedListNode[T]
	length int
}

// 新建单向链表
func NewSingleLinkedList[T any](e ...T) *SingleLinkedList[T] {
	ll := &SingleLinkedList[T]{}
	ll.Add(e...)
	return ll
}

// 检查越界
func (self *SingleLinkedList[T]) checkOut(i int) {
	length := self.Length()
	if i >= length {
		panic(fmt.Sprintf("index out of range [%d] with length %d", i, length))
	}
}

// 转成字符串 O(N)
func (self *SingleLinkedList[T]) String() string {
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
func (self *SingleLinkedList[T]) MarshalJSON() ([]byte, error) {
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
func (self *SingleLinkedList[T]) UnmarshalJSON(data []byte) error {
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
func (self *SingleLinkedList[T]) Length() int {
	return self.length
}

// 是否为空 O(1)
func (self *SingleLinkedList[T]) Empty() bool {
	return self.head == nil
}

// 增加元素 O(1)
func (self *SingleLinkedList[T]) Add(e ...T) {
	self.PushBack(e...)
}

// 增加到头部 O(1)
func (self *SingleLinkedList[T]) PushFront(e ...T) {
	if len(e) == 0 {
		return
	}
	var begin, end *singleLinkedListNode[T]
	for _, elem := range e {
		node := &singleLinkedListNode[T]{elem: elem}
		if begin == nil {
			begin, end = node, node
		} else {
			end.next = node
			end = node
		}
	}
	if self.head == nil {
		self.head, self.tail = begin, end
	} else {
		end.next = self.head
		self.head = begin
	}
	self.length += len(e)
}

// 增加到尾部 O(1)
func (self *SingleLinkedList[T]) PushBack(e ...T) {
	for _, i := range e {
		node := &singleLinkedListNode[T]{elem: i}
		if self.head == nil {
			self.head, self.tail = node, node
		} else {
			self.tail.next = node
			self.tail = node
		}
	}
	self.length += len(e)
}

// 插入元素 O(N)
func (self *SingleLinkedList[T]) Insert(i int, e ...T) {
	if len(e) == 0 {
		return
	}
	self.checkOut(i)
	if i == 0 { // 插入头部
		self.PushFront(e...)
	} else {
		var begin, end *singleLinkedListNode[T]
		for _, elem := range e {
			node := &singleLinkedListNode[T]{elem: elem}
			if begin == nil {
				begin, end = node, node
			} else {
				end.next = node
				end = node
			}
		}
		cursor := self.head
		for ; cursor != nil; cursor = cursor.next {
			if i == 1 {
				break
			}
			i--
		}
		end.next = cursor.next
		cursor.next = begin
		self.length += len(e)
	}
}

// 移除元素 O(N)
func (self *SingleLinkedList[T]) Remove(i int) T {
	if i == 0 {
		return self.PopFront()
	} else if i == self.length-1 {
		return self.PopBack()
	} else {
		self.checkOut(i)
		cursor := self.head
		for ; cursor != nil; cursor = cursor.next {
			if i == 1 {
				break
			}
			i--
		}
		elem := cursor.next.elem
		cursor.next = cursor.next.next
		self.length--
		return elem
	}
}

// 删除头元素 O(1)
func (self *SingleLinkedList[T]) PopFront() T {
	self.checkOut(0)
	elem := self.head.elem
	self.head = self.head.next
	self.length--
	return elem
}

// 删除尾元素 O(N)
func (self *SingleLinkedList[T]) PopBack() T {
	if self.length == 1 {
		return self.PopFront()
	}
	self.checkOut(0)
	elem := self.tail.elem
	cursor := self.head
	for ; cursor.next.next != nil; cursor = cursor.next {
	}
	cursor.next = nil
	self.tail = cursor
	self.length--
	return elem
}

// 获取元素 O(N)
func (self *SingleLinkedList[T]) Get(i int) T {
	self.checkOut(i)
	cursor := self.head
	for ; cursor != nil; cursor = cursor.next {
		if i == 0 {
			break
		}
		i--
	}
	return cursor.elem
}

// 获取第一个元素 O(1)
func (self *SingleLinkedList[T]) First() T {
	self.checkOut(0)
	return self.head.elem
}

// 获取最后一个元素 O(1)
func (self *SingleLinkedList[T]) Last() T {
	self.checkOut(0)
	return self.tail.elem
}

// 设置元素 O(N)
func (self *SingleLinkedList[T]) Set(i int, e T) T {
	self.checkOut(i)
	cursor := self.head
	for ; cursor != nil; cursor = cursor.next {
		if i == 0 {
			break
		}
		i--
	}
	elem := cursor.elem
	cursor.elem = e
	return elem
}

// 清空 O(1)
func (self *SingleLinkedList[T]) Clear() {
	if self.Empty() {
		return
	}
	self.head = nil
	self.tail = nil
	self.length = 0
}

// 克隆 O(N)
func (self *SingleLinkedList[T]) Clone() *SingleLinkedList[T] {
	sll := NewSingleLinkedList[T]()
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		sll.PushBack(cursor.elem)
	}
	return sll
}

// 过滤 O(N)
func (self *SingleLinkedList[T]) Filter(f func(i int, v T) bool) *SingleLinkedList[T] {
	sll := NewSingleLinkedList[T]()
	var index int
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		if f(index, cursor.elem) {
			sll.Add(cursor.elem)
		}
		index++
	}
	return sll
}

// 切分[b, e) O(N)
func (self *SingleLinkedList[T]) Slice(b, e int) *SingleLinkedList[T] {
	self.checkOut(b)
	sll := NewSingleLinkedList[T]()
	cursor := self.head
	for i := b; i != 0; cursor = cursor.next {
		i--
	}
	for ; cursor != nil && b < e; cursor = cursor.next {
		sll.Add(cursor.elem)
		b++
	}
	return sll
}

// 拼接 O(N)
func (self *SingleLinkedList[T]) Contact(a *SingleLinkedList[T]) {
	for cursor := a.head; cursor != nil; cursor = cursor.next {
		self.PushBack(cursor.elem)
	}
}

// 是否有任何元素满足条件 O(N)
func (self *SingleLinkedList[T]) Any(f func(i int, v T) bool) bool {
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
func (self *SingleLinkedList[T]) Every(f func(i int, v T) bool) bool {
	var index int
	for cursor := self.head; cursor != nil; cursor = cursor.next {
		if !f(index, cursor.elem) {
			return false
		}
		index++
	}
	return true
}

// SingleLinkedListMap 映射
func SingleLinkedListMap[T any, E any](src *SingleLinkedList[T], f func(i int, v T) E) *SingleLinkedList[E] {
	newList := NewSingleLinkedList[E]()
	var i int
	for cursor := src.head; cursor != nil; cursor = cursor.next {
		newList.PushBack(f(i, cursor.elem))
		i++
	}
	return newList
}

// 获取迭代器
func (self *SingleLinkedList[T]) Iterator() *SingleLinkedListIterator[T] {
	return &SingleLinkedListIterator[T]{cursor: self.head}
}

// 迭代器
type SingleLinkedListIterator[T any] struct {
	cursor *singleLinkedListNode[T] // 目前节点
	index  int                      // 索引
}

// 是否存在值
func (self *SingleLinkedListIterator[T]) HasValue() bool {
	return self.cursor != nil
}

// 是否存在下一个
func (self *SingleLinkedListIterator[T]) HasNext() bool {
	return self.cursor.next != nil
}

// 下一个
func (self *SingleLinkedListIterator[T]) Next() {
	self.cursor = self.cursor.next
	self.index++
}

// 获取索引
func (self *SingleLinkedListIterator[T]) Index() int {
	return self.index
}

// 获取值
func (self *SingleLinkedListIterator[T]) Value() T {
	return self.cursor.elem
}
