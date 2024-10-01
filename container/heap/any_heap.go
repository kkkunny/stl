package stlheap

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/zyedidia/generic/heap"

	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	stlslices "github.com/kkkunny/stl/container/slices"
)

type anyHeapData[T any] struct {
	data []T
	less func(a, b T) bool
}

func toAnyHeapData[T any](h *heap.Heap[T]) *anyHeapData[T] {
	if h == nil {
		return nil
	}
	return (*anyHeapData[T])(unsafe.Pointer(h))
}

func (self *anyHeapData[T]) toHeap() *heap.Heap[T] {
	if self == nil {
		return nil
	}
	return (*heap.Heap[T])(unsafe.Pointer(self))
}

type _AnyHeap[T any] struct {
	data *heap.Heap[T]
}

// _NewMinAnyHeap 新建小顶堆
func _NewMinAnyHeap[T any]() Heap[T] {
	f := stlcmp.GetCompareFunc[T]()
	return &_AnyHeap[T]{data: heap.New(func(a, b T) bool {
		return f(a, b) < 0
	})}
}

// _NewMaxAnyHeap 新建大顶堆
func _NewMaxAnyHeap[T any]() Heap[T] {
	f := stlcmp.GetCompareFunc[T]()
	return &_AnyHeap[T]{data: heap.New(func(a, b T) bool {
		return f(a, b) > 0
	})}
}

// _NewMinAnyHeapWith 新建小顶堆
func _NewMinAnyHeapWith[T any](vs ...T) Heap[T] {
	h := _NewMinAnyHeap[T]()
	for _, v := range vs {
		h.Push(v)
	}
	return h
}

// _NewMaxAnyHeapWith 新建大顶堆
func _NewMaxAnyHeapWith[T any](vs ...T) Heap[T] {
	h := _NewMaxAnyHeap[T]()
	for _, v := range vs {
		h.Push(v)
	}
	return h
}

// Clone 克隆
func (self *_AnyHeap[T]) Clone() Heap[T] {
	h := toAnyHeapData(self.data)
	newData := &anyHeapData[T]{
		data: stlslices.Clone(h.data),
		less: h.less,
	}
	return &_AnyHeap[T]{data: newData.toHeap()}
}

// Equal 比较相等 O(3N+2log(N))
func (self *_AnyHeap[T]) Equal(dst Heap[T]) bool {
	if self.Length() != dst.Length() {
		return false
	}
	lv, rv := stlslices.Sort(toAnyHeapData(self.data).data), stlslices.Sort(dst.ToSlice())
	return stlslices.Equal(lv, rv)
}

func (self *_AnyHeap[T]) NewWithIterator(iter stliter.Iterator[T]) any {
	data := heap.New(toAnyHeapData(self.data).less)
	for iter.Next() {
		data.Push(iter.Value())
	}
	return &_AnyHeap[T]{data: data}
}

// Iterator 迭代器
func (self *_AnyHeap[T]) Iterator() stliter.Iterator[T] {
	return stliter.NewSliceIterator(toAnyHeapData(self.data).data...)
}

// Length 长度
func (self *_AnyHeap[T]) Length() uint {
	return uint(self.data.Size())
}

// Push 入堆 O(log(N))
func (self *_AnyHeap[T]) Push(v T, vs ...T) {
	for _, v = range append([]T{v}, vs...) {
		self.data.Push(v)
	}
}

// Pop 出堆 O(log(N))
func (self *_AnyHeap[T]) Pop() T {
	v, _ := self.data.Pop()
	return v
}

// Peek 头节点 O(1)
func (self *_AnyHeap[T]) Peek() T {
	v, _ := self.data.Peek()
	return v
}

// Clear 清空
func (self *_AnyHeap[T]) Clear() {
	h := toAnyHeapData(self.data)
	self.data = heap.New(h.less)
}

// Empty 是否为空
func (self *_AnyHeap[T]) Empty() bool {
	_, ok := self.data.Peek()
	return !ok
}

// String 获取字符串
func (self *_AnyHeap[T]) String() string {
	var buf strings.Builder
	buf.WriteString("Heap{")
	for i, v := range toAnyHeapData(self.data).data {
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < self.data.Size()-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_AnyHeap[T]) ToSlice() []T {
	h := toAnyHeapData(self.data)
	return stlslices.Clone(h.data)
}
