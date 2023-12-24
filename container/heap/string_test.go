package heap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHeap_String(t *testing.T) {
	v1 := NewMinHeapWith[int](1, 3, 2)
	stltest.AssertEq(t, v1.String(), "Heap{1, 2, 3}")
}
