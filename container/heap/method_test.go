package stack

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
