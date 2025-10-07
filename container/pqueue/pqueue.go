package pqueue

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/clone"
	stlcmp "github.com/kkkunny/stl/cmp"
	stlheap "github.com/kkkunny/stl/container/heap"
	stliter "github.com/kkkunny/stl/container/iter"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/container/tuple"
)

type PQueue[Prior cmp.Ordered, Elem any] interface {
	Iter2() iter.Seq2[Prior, Elem]
	clone.Cloneable[PQueue[Prior, Elem]]
	stlcmp.Equalable[PQueue[Prior, Elem]]
	stliter.IteratorContainer[tuple.Tuple2[Prior, Elem]]
	stlbasic.Lengthable
	Push(prior Prior, value Elem)
	Pop() (Prior, Elem)
	Peek() (Prior, Elem)
	Clear()
	Empty() bool
	fmt.Stringer
	ToSlice() []tuple.Tuple2[Prior, Elem]
	getData() stlheap.Heap[anyPQueueNode[Prior, Elem]]
}

// AnyWith 使用自定义cmp函数
func AnyWith[Prior cmp.Ordered, Elem any](vs ...any) PQueue[Prior, Elem] {
	if len(vs) == 0 {
		return _NewAnyPQueue[Prior, Elem]()
	} else {
		return _NewAnyPQueueWith[Prior, Elem](vs...)
	}
}

type anyPQueueNode[Prior cmp.Ordered, Elem any] struct {
	priority Prior
	value    Elem
}

func (self anyPQueueNode[Prior, Elem]) Equal(dst anyPQueueNode[Prior, Elem]) bool {
	return self.priority == dst.priority
}

func (self anyPQueueNode[Prior, Elem]) Compare(dst anyPQueueNode[Prior, Elem]) int {
	return cmp.Compare(self.priority, dst.priority)
}

type _AnyPQueue[Prior cmp.Ordered, Elem any] struct {
	data stlheap.Heap[anyPQueueNode[Prior, Elem]]
}

func _NewAnyPQueue[Prior cmp.Ordered, Elem any]() PQueue[Prior, Elem] {
	return &_AnyPQueue[Prior, Elem]{data: stlheap.AnyMaxWith[anyPQueueNode[Prior, Elem]]()}
}

func _NewAnyPQueueWith[Prior cmp.Ordered, Elem any](vs ...any) PQueue[Prior, Elem] {
	self := _NewAnyPQueue[Prior, Elem]()
	for i := 0; i < len(vs); i += 2 {
		self.Push(vs[i].(Prior), vs[i+1].(Elem))
	}
	return self
}

// Clone 克隆
func (self *_AnyPQueue[Prior, Elem]) Clone() PQueue[Prior, Elem] {
	return &_AnyPQueue[Prior, Elem]{data: clone.Clone(self.data)}
}

// Equal 比较相等
func (self *_AnyPQueue[Prior, Elem]) Equal(dst PQueue[Prior, Elem]) bool {
	return self.data.Equal(dst.getData())
}

func (self *_AnyPQueue[Prior, Elem]) NewWithIterator(iter stliter.Iterator[tuple.Tuple2[Prior, Elem]]) any {
	pq := _NewAnyPQueue[Prior, Elem]()
	stliter.IteratorForeach(iter, func(v tuple.Tuple2[Prior, Elem]) bool {
		pq.Push(v.Unpack())
		return true
	})
	return pq
}

// Iterator 迭代器
func (self *_AnyPQueue[Prior, Elem]) Iterator() stliter.Iterator[tuple.Tuple2[Prior, Elem]] {
	return stliter.NewSliceIterator(self.ToSlice()...)
}

func (self *_AnyPQueue[Prior, Elem]) Iter2() iter.Seq2[Prior, Elem] {
	f := self.data.Iter()
	return func(yield func(Prior, Elem) bool) {
		f(func(node anyPQueueNode[Prior, Elem]) bool {
			return yield(node.priority, node.value)
		})
	}
}

// Length 长度
func (self *_AnyPQueue[Prior, Elem]) Length() uint {
	return self.data.Length()
}

// Push 入队
func (self *_AnyPQueue[Prior, Elem]) Push(prior Prior, value Elem) {
	self.data.Push(anyPQueueNode[Prior, Elem]{
		priority: prior,
		value:    value,
	})
}

// Pop 出队
func (self *_AnyPQueue[Prior, Elem]) Pop() (Prior, Elem) {
	node := self.data.Pop()
	return node.priority, node.value
}

// Peek 头元素
func (self *_AnyPQueue[Prior, Elem]) Peek() (Prior, Elem) {
	node := self.data.Peek()
	return node.priority, node.value
}

// Clear 清空
func (self *_AnyPQueue[Prior, Elem]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_AnyPQueue[Prior, Elem]) Empty() bool {
	return self.data.Empty()
}

// String 获取字符串
func (self *_AnyPQueue[Prior, Elem]) String() string {
	var buf strings.Builder
	buf.WriteString("PQueue{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value().E1()))
		buf.WriteByte(':')
		buf.WriteString(fmt.Sprintf("%v", iter.Value().E2()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_AnyPQueue[Prior, Elem]) ToSlice() []tuple.Tuple2[Prior, Elem] {
	data := self.data.ToSlice()
	slices.SortFunc(data, func(a anyPQueueNode[Prior, Elem], b anyPQueueNode[Prior, Elem]) int {
		return -a.Compare(b)
	})
	return stlslices.Map(data, func(_ int, node anyPQueueNode[Prior, Elem]) tuple.Tuple2[Prior, Elem] {
		return tuple.Pack2(node.priority, node.value)
	})
}

func (self *_AnyPQueue[Prior, Elem]) getData() stlheap.Heap[anyPQueueNode[Prior, Elem]] {
	return self.data
}
