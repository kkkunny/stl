package hashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestHashSet_Empty(t *testing.T) {
	var hs HashSet[int]
	stltest.AssertEq(t, hs.Empty(), true)
	hs = NewHashSetWith[int](1, 2, 2)
	stltest.AssertEq(t, hs.Empty(), false)
}

func TestHashSet_Clear(t *testing.T) {
	hs := NewHashSetWith[int](1, 2, 2)
	stltest.AssertEq(t, hs.Empty(), false)
	hs.Clear()
	stltest.AssertEq(t, hs.Empty(), true)
}

func TestHashSet_Push(t *testing.T) {
	hs := NewHashSetWith[int](1, 2, 2)
	stltest.AssertEq(t, hs.Push(1), false)
	stltest.AssertEq(t, hs.Push(3), true)
}

func TestHashSet_Remove(t *testing.T) {
	hs := NewHashSetWith[int](1, 2, 2)
	stltest.AssertEq(t, hs.Remove(1), 1)
	stltest.AssertEq(t, hs.Remove(3), 3)
}
