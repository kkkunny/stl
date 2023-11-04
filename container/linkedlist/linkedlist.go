package linkedlist

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/iterator"
)

type _LinkedListNode[T any] struct {
	Prev, Next *_LinkedListNode[T]
	Value      T
}

// LinkedList 链表
type LinkedList[T any] struct {
	length     uint
	root, tail *_LinkedListNode[T]
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

func NewLinkedListWith[T any](vs ...T) LinkedList[T] {
	self := NewLinkedList[T]()
	for _, v := range vs {
		self.PushBack(v)
	}
	return self
}

func (_ LinkedList[T]) NewWithIterator(iter iterator.Iterator[T]) LinkedList[T] {
	self := NewLinkedList[T]()
	for iter.Next() {
		self.PushBack(iter.Value())
	}
	return self
}

func (self LinkedList[T]) Length() uint {
	return self.length
}

func (self LinkedList[T]) Equal(dst LinkedList[T]) bool {
	if self.root == dst.root {
		return true
	}

	if self.Length() != dst.Length() {
		return false
	}

	for c1, c2 := self.root, dst.root; c1 != nil && c2 != nil; c1, c2 = c1.Next, c2.Next {
		if !stlbasic.Equal(c1.Value, c2.Value) {
			return false
		}
	}
	return true
}

func (self *LinkedList[T]) PushBack(v T) {
	node := &_LinkedListNode[T]{Value: v}
	if self.root == nil {
		self.root = node
		self.tail = node
	} else {
		self.tail.Next = node
		node.Prev = self.tail
		self.tail = node
	}
	self.length++
}

func (self *LinkedList[T]) PushFront(v T) {
	node := &_LinkedListNode[T]{Value: v}
	if self.root == nil {
		self.root = node
		self.tail = node
	} else {
		self.root.Prev = node
		node.Next = self.root
		self.root = node
	}
	self.length++
}

func (self LinkedList[T]) String() string {
	var buf strings.Builder
	buf.WriteByte('[')
	for cursor := self.root; cursor != nil; cursor = cursor.Next {
		buf.WriteString(fmt.Sprintf("%v", cursor.Value))
		if cursor.Next != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func (self *LinkedList[T]) PopBack() T {
	var v T
	if self.root == self.tail {
		v = self.root.Value
		self.root = nil
		self.tail = nil
	} else {
		v = self.tail.Value
		self.tail.Prev.Next = nil
		self.tail = self.tail.Prev
	}
	self.length--
	return v
}

func (self *LinkedList[T]) PopFront() T {
	var v T
	if self.root == self.tail {
		v = self.root.Value
		self.root = nil
		self.tail = nil
	} else {
		v = self.root.Value
		self.root.Next.Prev = nil
		self.root = self.root.Next
	}
	self.length--
	return v
}

func (self LinkedList[T]) Back() T {
	return self.tail.Value
}

func (self LinkedList[T]) Front() T {
	return self.root.Value
}

func (self LinkedList[T]) Clone() LinkedList[T] {
	list := NewLinkedList[T]()
	for cursor := self.root; cursor != nil; cursor = cursor.Next {
		list.PushBack(cursor.Value)
	}
	return list
}

func (self *LinkedList[T]) Clear() {
	self.root = nil
	self.tail = nil
	self.length = 0
}

func (self LinkedList[T]) Empty() bool {
	return self.length == 0
}

func (self LinkedList[T]) Iterator() iterator.Iterator[T] {
	return iterator.NewIterator[T](_NewIterator[T](&self))
}
