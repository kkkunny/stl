package stack

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHeap_Equal(t *testing.T) {
	v1 := NewMinHeapWith[int](1, 2, 3)
	v2 := NewMinHeapWith[int](1, 2, 3)
	v3 := NewMinHeapWith[int](2, 1, 0)
	stltest.AssertEq(t, v1.Equal(v2), true)
	stltest.AssertEq(t, v2.Equal(v3), false)
	stltest.AssertEq(t, v1.Equal(v3), false)
}
