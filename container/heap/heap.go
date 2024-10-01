package stlheap

import (
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/value"
)

type Heap[T any] interface {
	heapIter[T]
	stlval.Cloneable[Heap[T]]
	stlcmp.Equalable[Heap[T]]
	stliter.IteratorContainer[T]
	stlbasic.Lengthable
	Push(v T, vs ...T)
	Pop() T
	Peek() T
	Clear()
	Empty() bool
	fmt.Stringer
	ToSlice() []T
}

// AnyMinWith 使用自定义cmp函数
func AnyMinWith[T any](vs ...T) Heap[T] {
	if len(vs) == 0 {
		return _NewMinAnyHeap[T]()
	} else {
		return _NewMinAnyHeapWith[T](vs...)
	}
}

// AnyMaxWith 使用自定义cmp函数
func AnyMaxWith[T any](vs ...T) Heap[T] {
	if len(vs) == 0 {
		return _NewMaxAnyHeap[T]()
	} else {
		return _NewMaxAnyHeapWith[T](vs...)
	}
}
