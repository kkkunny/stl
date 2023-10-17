package stack

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewStack(t *testing.T) {
	v := NewStack[int]()
	stltest.AssertEq(t, v.Length(), 0)
	stltest.AssertEq(t, v.Capacity(), 0)
}

func TestNewStackWithCapacity(t *testing.T) {
	v := NewStackWithCapacity[int](10)
	stltest.AssertEq(t, v.Capacity(), 10)
}

func TestNewStackWith(t *testing.T) {
	v := NewStackWith(1, 2, 3)
	stltest.AssertEq(t, v.Length(), 3)
	stltest.AssertEq(t, v.Capacity(), 3)
}

func TestStack_Push(t *testing.T) {
	v := NewStackWith(1, 2, 3)
	v.Push(4)
	stltest.AssertEq(t, v.Pop(), 4)
}

func TestStack_Peek(t *testing.T) {
	v := NewStackWith(1, 2, 3)
	stltest.AssertEq(t, v.Peek(), 3)
}

func TestStack_Pop(t *testing.T) {
	v := NewStackWith(1, 2, 3)
	stltest.AssertEq(t, v.Pop(), 3)
	stltest.AssertEq(t, v.Pop(), 2)
	stltest.AssertEq(t, v.Pop(), 1)
}

func TestStack_Clear(t *testing.T) {
	v := NewStackWith(1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, int(v.Length()), 0)
}

func TestStack_Empty(t *testing.T) {
	v := NewStackWith(1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, v.Empty(), true)
}
