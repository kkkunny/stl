package queue

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/heap"
	"github.com/kkkunny/stl/list"
	"github.com/kkkunny/stl/table"
	. "github.com/kkkunny/stl/types"
)

type priorityQueueNode[P Comparator[P], V any] table.Entry[P, V]

func (self *priorityQueueNode[P, V]) Compare(dst *priorityQueueNode[P, V]) int {
	return self.Key.Compare(dst.Key)
}

// 优先级队列
type PriorityQueue[P Comparator[P], V any] struct {
	data *heap.MaxHeap[*priorityQueueNode[P, V]]
}

// 新建队列
func NewPriorityQueue[P Comparator[P], V any]() *PriorityQueue[P, V] {
	return &PriorityQueue[P, V]{
		data: heap.NewMaxHeap[*priorityQueueNode[P, V]](),
	}
}

func (self *PriorityQueue[P, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('[')
	var index Usize
	length := self.data.Length()
	for iter := self.data.Iterator(); iter.HasValue(); iter.Next() {
		node := iter.Value()
		buf.WriteString(fmt.Sprintf("%v: %v", node.Key, node.Value))
		if index < length-1 {
			buf.WriteString(", ")
		}
		index++
	}
	buf.WriteByte(']')
	return buf.String()
}

// 获取长度
func (self *PriorityQueue[P, V]) Length() Usize {
	return self.data.Length()
}

// 是否为空
func (self *PriorityQueue[P, V]) Empty() bool {
	return self.data.Empty()
}

// 压入队列
func (self *PriorityQueue[P, V]) Push(p P, v V) {
	self.data.Push(&priorityQueueNode[P, V]{
		Key:   p,
		Value: v,
	})
}

// 弹出队列
func (self *PriorityQueue[P, V]) Pop() (P, V) {
	node := self.data.Pop()
	return node.Key, node.Value
}

// 提前获取队首
func (self *PriorityQueue[P, V]) Peek() (P, V) {
	node := self.data.Peek()
	return node.Key, node.Value
}

// 清空
func (self *PriorityQueue[P, V]) Clear() {
	self.data.Clear()
}

// 克隆
func (self *PriorityQueue[P, V]) Clone() *PriorityQueue[P, V] {
	return &PriorityQueue[P, V]{
		data: self.data.Clone(),
	}
}

// 获取迭代器
func (self *PriorityQueue[P, V]) Iterator() *PriorityQueueIterator[P, V] {
	return &PriorityQueueIterator[P, V]{iter: self.data.Iterator()}
}

// 迭代器
type PriorityQueueIterator[P Comparator[P], V any] struct {
	iter *list.ArrayListIterator[*priorityQueueNode[P, V]]
}

// 是否存在值
func (self *PriorityQueueIterator[P, V]) HasValue() bool {
	return self.iter.HasValue()
}

// 下一个
func (self *PriorityQueueIterator[P, V]) Next() {
	self.iter.Next()
}

// 获取值
func (self *PriorityQueueIterator[P, V]) Value() (P, V) {
	node := self.iter.Value()
	return node.Key, node.Value
}
