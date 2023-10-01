package hashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewHashSet(t *testing.T) {
	v := NewHashSet[int]()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestNewHashSetWithCapacity(t *testing.T) {
	v := NewHashSetWithCapacity[int](10)
	stltest.AssertEq(t, v.Length(), 0)
}

func TestNewHashSetWith(t *testing.T) {
	v := NewHashSetWith[int](1, 2, 3, 1)
	stltest.AssertEq(t, v.Length(), 3)
}

func TestHashSet_Push(t *testing.T) {
	v := NewHashSetWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Push(1), false)
	stltest.AssertEq(t, v.Push(4), true)
}

func TestHashSet_Remove(t *testing.T) {
	v := NewHashSetWith[int](1, 2, 3)
	stltest.AssertEq(t, v.Push(1), false)
	v.Remove(1)
	stltest.AssertEq(t, v.Push(1), true)
}

func TestHashSet_Clear(t *testing.T) {
	v := NewHashSetWith[int](1, 2, 3)
	v.Clear()
	stltest.AssertEq(t, int(v.Length()), 0)
}

func TestHashSet_Empty(t *testing.T) {
	v := NewHashSet[int]()
	stltest.AssertEq(t, v.Empty(), true)
}

func TestHashSet_Contain(t *testing.T) {
	v := NewHashSet[int]()
	stltest.AssertEq(t, v.Contain(1), false)
	v.Push(1)
	stltest.AssertEq(t, v.Contain(1), true)
}
