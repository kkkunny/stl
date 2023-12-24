package heap

import (
	"testing"

	"github.com/kkkunny/stl/container/iterator"
	stltest "github.com/kkkunny/stl/test"
)

func TestHeap_Iterator(t *testing.T) {
	v1 := NewMinHeapWith[int](3, 2, 1)
	v2 := iterator.Map[int, int, Heap[int]](v1, func(v int) int {
		return v - 1
	})
	stltest.AssertEq(t, v2, NewMinHeapWith[int](2, 1, 0))
}
