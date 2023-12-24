package heap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHeap_Push(t *testing.T) {
	h := NewMinHeapWith(3, 2, 1)
	stltest.AssertEq(t, h.Pop(), 1)
	stltest.AssertEq(t, h.Pop(), 2)
	stltest.AssertEq(t, h.Pop(), 3)
}

func TestHeap_Pop(t *testing.T) {
	h := NewMinHeapWith(3, 2, 1)
	stltest.AssertEq(t, h.Pop(), 1)
	stltest.AssertEq(t, h.Pop(), 2)
	stltest.AssertEq(t, h.Pop(), 3)
}

func TestHeap_Peek(t *testing.T) {
	h := NewMinHeapWith(3, 2, 1)
	stltest.AssertEq(t, h.Peek(), 1)
	stltest.AssertEq(t, h.Pop(), 1)
	stltest.AssertEq(t, h.Peek(), 2)
	stltest.AssertEq(t, h.Pop(), 2)
	stltest.AssertEq(t, h.Peek(), 3)
	stltest.AssertEq(t, h.Pop(), 3)
}

func TestHeap_Reverse(t *testing.T) {
	h := NewMinHeapWith(3, 2, 1)
	h.Reverse()
	stltest.AssertEq(t, h, NewMaxHeapWith(3, 2, 1))
}

func TestHeap_Clear(t *testing.T) {
	h := NewMinHeapWith(3, 2, 1)
	h.Clear()
	stltest.AssertEq(t, h.Length(), 0)
}

func TestHeap_Empty(t *testing.T) {
	h := NewMinHeapWith(3, 2, 1)
	stltest.AssertEq(t, h.Empty(), false)
	h.Clear()
	stltest.AssertEq(t, h.Empty(), true)
}
