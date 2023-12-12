package queue

import (
	"fmt"
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestQueue_PushAndPop(t *testing.T) {
	v := NewQueueWith[int](1)
	v.Push(2, 3)
	stltest.AssertEq(t, v.Pop(), 1)
	stltest.AssertEq(t, v.Pop(), 2)
	stltest.AssertEq(t, v.Pop(), 3)
}

func TestQueue_Peek(t *testing.T) {
	v := NewQueueWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Peek(), 1)
	stltest.AssertEq(t, v.Pop(), 1)
	stltest.AssertEq(t, v.Peek(), 2)
}

func TestDynArray_Clear(t *testing.T) {
	v := NewQueueWith[int](1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestDynArray_Empty(t *testing.T) {
	v := NewQueueWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Empty(), false)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}

func TestDynArray_Append(t *testing.T) {
	var v Queue[int]
	v.Append(NewQueueWith[int](1, 2, 3))
	stltest.AssertEq(t, v.Pop(), 1)
	stltest.AssertEq(t, v.Pop(), 2)
	stltest.AssertEq(t, v.Pop(), 3)
}

func TestDynArray_ToSlice(t *testing.T) {
	v := NewQueueWith[int](1, 2, 3)
	stltest.AssertEq(t, fmt.Sprintf("%v", v.ToSlice()), fmt.Sprintf("%v", []int{1, 2, 3}))
}
