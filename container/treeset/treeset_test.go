package treeset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewTreeSet(t *testing.T) {
	v := NewTreeSet[int]()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestNewTreeSetWith(t *testing.T) {
	v := NewTreeSetWith[int](1, 2, 3, 1)
	stltest.AssertEq(t, v.Length(), 3)
}

func TestTreeSet_Push(t *testing.T) {
	v := NewTreeSetWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Push(1), false)
	stltest.AssertEq(t, v.Push(4), true)
}

func TestTreeSet_Remove(t *testing.T) {
	v := NewTreeSetWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Push(1), false)
	v.Remove(1)
	stltest.AssertEq(t, v.Push(1), true)
}

func TestTreeSet_Clear(t *testing.T) {
	v := NewTreeSetWith[int](1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, int(v.Length()), 0)
}

func TestTreeSet_Empty(t *testing.T) {
	v := NewTreeSet[int]()
	stltest.AssertEq(t, v.Empty(), true)
}

func TestTreeSet_Contain(t *testing.T) {
	v := NewTreeSet[int]()
	stltest.AssertEq(t, v.Contain(1), false)
	v.Push(1)
	stltest.AssertEq(t, v.Contain(1), true)
}

func TestTreeSet_Front(t *testing.T) {
	v := NewTreeSetWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Front(), 1)
}

func TestTreeSet_Back(t *testing.T) {
	v := NewTreeSetWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Back(), 3)
}
