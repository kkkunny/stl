package pqueue

import (
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
	stlcmp "github.com/kkkunny/stl/cmp"
	stlheap "github.com/kkkunny/stl/container/heap"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
	"github.com/kkkunny/stl/value"
)

type PQueue[T any] interface {
	pqueueIter[T]
	stlval.Cloneable[PQueue[T]]
	stlcmp.Equalable[PQueue[T]]
	stliter.IteratorContainer[pair.Pair[uint64, T]]
	stlbasic.Lengthable
	Push(prior uint64, value T)
	Pop() (uint64, T)
	Peek() (uint64, T)
	Clear()
	Empty() bool
	fmt.Stringer
	ToSlice() []pair.Pair[uint64, T]
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
