package pqueue

import (
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/clone"
	stlcmp "github.com/kkkunny/stl/cmp"
	stlheap "github.com/kkkunny/stl/container/heap"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/tuple"
)

type PQueue[T any] interface {
	pqueueIter[T]
	clone.Cloneable[PQueue[T]]
	stlcmp.Equalable[PQueue[T]]
	stliter.IterContainer[tuple.Tuple2[uint64, T]]
	stlbasic.Lengthable
	Push(prior uint64, value T)
	Pop() (uint64, T)
	Peek() (uint64, T)
	Clear()
	Empty() bool
	fmt.Stringer
	ToSlice() []tuple.Tuple2[uint64, T]
	getData() stlheap.Heap[anyPQueueNode[T]]
}

// AnyWith 使用自定义cmp函数
func AnyWith[T any](vs ...any) PQueue[T] {
	if len(vs) == 0 {
		return _NewAnyPQueue[T]()
	} else {
		return _NewAnyPQueueWith[T](vs...)
	}
}
