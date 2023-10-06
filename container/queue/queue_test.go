package queue

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewQueue(t *testing.T) {
	v := NewQueue[int]()
	stltest.AssertEq(t, v.Length(), 0)
	stltest.AssertEq(t, v.Capacity(), 0)
}

func TestNewQueueWithCapacity(t *testing.T) {
	v := NewQueueWithCapacity[int](10)
	stltest.AssertEq(t, v.Capacity(), 10)
}

func TestNewQueueWith(t *testing.T) {
	v := NewQueueWith(1, 2, 3)
	stltest.AssertEq(t, v.Length(), 3)
	stltest.AssertEq(t, v.Capacity(), 3)
}

func TestQueue_Push(t *testing.T) {
	v := NewQueueWith(1, 2, 3)
	v.Push(4)
	stltest.AssertEq(t, v.Peek(), 1)
}

func TestQueue_Peek(t *testing.T) {
	v := NewQueueWith(1, 2, 3)
	stltest.AssertEq(t, v.Peek(), 1)
}

func TestQueue_Pop(t *testing.T) {
	v := NewQueueWith(1, 2, 3)
	stltest.AssertEq(t, v.Pop(), 1)
	stltest.AssertEq(t, v.Pop(), 2)
	stltest.AssertEq(t, v.Pop(), 3)
}

func TestQueue_Clear(t *testing.T) {
	v := NewQueueWith(1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, int(v.Length()), 0)
}

func TestQueue_Empty(t *testing.T) {
	v := NewQueueWith(1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}
