package queue

import (
	"testing"

	"github.com/kkkunny/stl/container/pair"
	stltest "github.com/kkkunny/stl/test"
)

func TestPriorityQueue_PushAndPop(t *testing.T) {
	v := NewPriorityQueue[int]()
	v.Push(1, 1)
	v.Push(2, 2)
	v.Push(3, 3)
	stltest.AssertEq(t, pair.NewPair(v.Pop()), pair.NewPair[uint64, int](3, 3))
	stltest.AssertEq(t, pair.NewPair(v.Pop()), pair.NewPair[uint64, int](2, 2))
	stltest.AssertEq(t, pair.NewPair(v.Pop()), pair.NewPair[uint64, int](1, 1))
}

func TestPriorityQueue_Peek(t *testing.T) {
	v := NewPriorityQueue[int]()
	v.Push(1, 1)
	v.Push(2, 2)
	v.Push(3, 3)
	stltest.AssertEq(t, pair.NewPair(v.Peek()), pair.NewPair[uint64, int](3, 3))
	stltest.AssertEq(t, pair.NewPair(v.Pop()), pair.NewPair[uint64, int](3, 3))
	stltest.AssertEq(t, pair.NewPair(v.Peek()), pair.NewPair[uint64, int](2, 2))
}

func TestPriorityQueue_Clear(t *testing.T) {
	v := NewPriorityQueue[int]()
	v.Push(1, 1)
	stltest.AssertEq(t, v.Length(), 1)
	v.Clear()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestPriorityQueue_Empty(t *testing.T) {
	v := NewPriorityQueue[int]()
	v.Push(1, 1)
	stltest.AssertEq(t, v.Empty(), false)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}
