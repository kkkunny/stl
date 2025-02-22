package stlheap

import (
	"testing"

	"github.com/kkkunny/stl/clone"
	stltest "github.com/kkkunny/stl/test"
)

func TestAnyHeap_Clone(t *testing.T) {
	v1 := _NewMinAnyHeapWith[int](1, 2, 3)
	v2 := clone.Clone(v1)
	stltest.AssertEq(t, v1.Equal(v2), true)
}

func TestAnyHeap_Equal(t *testing.T) {
	v1 := _NewMinAnyHeapWith[int](1, 2, 3)
	v2 := _NewMinAnyHeapWith[int](1, 2, 3)
	v3 := _NewMinAnyHeapWith[int](2, 1, 0)
	stltest.AssertEq(t, v1.Equal(v2), true)
	stltest.AssertEq(t, v2.Equal(v3), false)
	stltest.AssertEq(t, v1.Equal(v3), false)
}

func TestAnyHeap_Length(t *testing.T) {
	h := _NewMinAnyHeapWith(3, 2, 1)
	stltest.AssertEq(t, h.Length(), 3)
}

func TestAnyHeap_Push(t *testing.T) {
	h := _NewMinAnyHeapWith(3, 2, 1)
	stltest.AssertEq(t, h.Pop(), 1)
	stltest.AssertEq(t, h.Pop(), 2)
	stltest.AssertEq(t, h.Pop(), 3)
}

func TestAnyHeap_Pop(t *testing.T) {
	h := _NewMinAnyHeapWith(3, 2, 1)
	stltest.AssertEq(t, h.Pop(), 1)
	stltest.AssertEq(t, h.Pop(), 2)
	stltest.AssertEq(t, h.Pop(), 3)
}

func TestAnyHeap_Peek(t *testing.T) {
	h := _NewMinAnyHeapWith(3, 2, 1)
	stltest.AssertEq(t, h.Peek(), 1)
	stltest.AssertEq(t, h.Pop(), 1)
	stltest.AssertEq(t, h.Peek(), 2)
	stltest.AssertEq(t, h.Pop(), 2)
	stltest.AssertEq(t, h.Peek(), 3)
	stltest.AssertEq(t, h.Pop(), 3)
}

func TestAnyHeap_Clear(t *testing.T) {
	h := _NewMinAnyHeapWith(3, 2, 1)
	h.Clear()
	stltest.AssertEq(t, h.Length(), 0)
}

func TestAnyHeap_Empty(t *testing.T) {
	h := _NewMinAnyHeapWith(3, 2, 1)
	stltest.AssertEq(t, h.Empty(), false)
	h.Clear()
	stltest.AssertEq(t, h.Empty(), true)
}

func TestAnyHeap_String(t *testing.T) {
	v1 := _NewMinAnyHeapWith[int](1, 3, 2)
	stltest.AssertNotEq(t, v1.String(), "")
}
