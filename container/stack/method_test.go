package stack

import (
	"fmt"
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestStack_PushAndPop(t *testing.T) {
	v := NewStackWith[int](1)
	v.Push(2, 3)
	stltest.AssertEq(t, v.Pop(), 3)
	stltest.AssertEq(t, v.Pop(), 2)
	stltest.AssertEq(t, v.Pop(), 1)
}

func TestStack_Peek(t *testing.T) {
	v := NewStackWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Peek(), 3)
	stltest.AssertEq(t, v.Pop(), 3)
	stltest.AssertEq(t, v.Peek(), 2)
}

func TestStack_Clear(t *testing.T) {
	v := NewStackWith[int](1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestStack_Empty(t *testing.T) {
	v := NewStackWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Empty(), false)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}

func TestStack_ToSlice(t *testing.T) {
	v := NewStackWith[int](1, 2, 3)
	stltest.AssertEq(t, fmt.Sprintf("%v", v.ToSlice()), fmt.Sprintf("%v", []int{1, 2, 3}))
}
