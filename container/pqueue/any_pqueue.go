package pqueue

import (
	"fmt"
	"strings"

	stlheap "github.com/kkkunny/stl/container/heap"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/internal/slices"
	stlval "github.com/kkkunny/stl/value"
)

type anyPQueueNode[T any] struct {
	priority uint64
	value    T
}

func (self anyPQueueNode[T]) Equal(dst anyPQueueNode[T]) bool {
	return self.priority == dst.priority
}

func (self anyPQueueNode[T]) Compare(dst anyPQueueNode[T]) int {
	if self.priority < dst.priority {
		return -1
	} else if self.priority == dst.priority {
		return 0
	} else {
		return 1
	}
}

type _AnyPQueue[T any] struct {
	data stlheap.Heap[anyPQueueNode[T]]
}

func _NewAnyPQueue[T any]() PQueue[T] {
	return &_AnyPQueue[T]{data: stlheap.AnyMaxWith[anyPQueueNode[T]]()}
}

func _NewAnyPQueueWith[T any](vs ...any) PQueue[T] {
	self := _NewAnyPQueue[T]()
	for i := 0; i < len(vs); i += 2 {
		self.Push(vs[i].(uint64), vs[i+1].(T))
	}
	return self
}

// Clone 克隆
func (self *_AnyPQueue[T]) Clone() PQueue[T] {
	return &_AnyPQueue[T]{data: stlval.Clone(self.data)}
}

// Equal 比较相等
func (self *_AnyPQueue[T]) Equal(dst PQueue[T]) bool {
	return self.data.Equal(dst.getData())
}

func (self *_AnyPQueue[T]) NewWithIterator(iter stliter.Iterator[pair.Pair[uint64, T]]) any {
	pq := _NewAnyPQueue[T]()
	stliter.IteratorForeach(iter, func(v pair.Pair[uint64, T]) bool {
		pq.Push(v.First, v.Second)
		return true
	})
	return pq
}

// Iterator 迭代器
func (self *_AnyPQueue[T]) Iterator() stliter.Iterator[pair.Pair[uint64, T]] {
	return stliter.NewSliceIterator(self.ToSlice()...)
}

// Length 长度
func (self *_AnyPQueue[T]) Length() uint {
	return self.data.Length()
}

// Push 入队
func (self *_AnyPQueue[T]) Push(prior uint64, value T) {
	self.data.Push(anyPQueueNode[T]{
		priority: prior,
		value:    value,
	})
}

// Pop 出队
func (self *_AnyPQueue[T]) Pop() (uint64, T) {
	node := self.data.Pop()
	return node.priority, node.value
}

// Peek 头元素
func (self *_AnyPQueue[T]) Peek() (uint64, T) {
	node := self.data.Peek()
	return node.priority, node.value
}

// Clear 清空
func (self *_AnyPQueue[T]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_AnyPQueue[T]) Empty() bool {
	return self.data.Empty()
}

// String 获取字符串
func (self *_AnyPQueue[T]) String() string {
	var buf strings.Builder
	buf.WriteString("PQueue{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%d", iter.Value().First))
		buf.WriteByte(':')
		buf.WriteString(fmt.Sprintf("%v", iter.Value().Second))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_AnyPQueue[T]) ToSlice() []pair.Pair[uint64, T] {
	data := self.data.ToSlice()
	slices.SortFunc(data, func(a anyPQueueNode[T], b anyPQueueNode[T]) int {
		return -a.Compare(b)
	})
	return stlslices.Map(data, func(_ int, node anyPQueueNode[T]) pair.Pair[uint64, T] {
		return pair.NewPair(node.priority, node.value)
	})
}

func (self *_AnyPQueue[T]) getData() stlheap.Heap[anyPQueueNode[T]] {
	return self.data
}
