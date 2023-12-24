package stack

import (
	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/dynarray"
)

// Heap 堆
type Heap[T any] struct {
	reverse bool
	data    *dynarray.DynArray[T]
}

// NewMinHeap 新建小顶堆
func NewMinHeap[T any]() Heap[T] {
	return Heap[T]{
		reverse: false,
		data:    stlbasic.Ptr(dynarray.NewDynArray[T]()),
	}
}

// NewMaxHeap 新建大顶堆
func NewMaxHeap[T any]() Heap[T] {
	return Heap[T]{
		reverse: true,
		data:    stlbasic.Ptr(dynarray.NewDynArray[T]()),
	}
}

// NewMinHeapWith 新建小顶堆
func NewMinHeapWith[T any](vs ...T) Heap[T] {
	h := NewMinHeap[T]()
	for _, v := range vs {
		h.Push(v)
	}
	return h
}

// NewMaxHeapWith 新建大顶堆
func NewMaxHeapWith[T any](vs ...T) Heap[T] {
	h := NewMaxHeap[T]()
	for _, v := range vs {
		h.Push(v)
	}
	return h
}
