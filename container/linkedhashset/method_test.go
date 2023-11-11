package linkedhashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedHashSet_Add(t *testing.T) {
	var hs LinkedHashSet[int]
	stltest.AssertEq(t, hs.Add(1), true)
	stltest.AssertEq(t, hs.Add(1), false)
	stltest.AssertEq(t, hs.Add(2), true)
}

func TestLinkedHashSet_Contain(t *testing.T) {
	hm := NewLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Contain(1), true)
	stltest.AssertEq(t, hm.Contain(3), false)
}

func TestLinkedHashSet_Remove(t *testing.T) {
	hm := NewLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Remove(3), false)
	stltest.AssertEq(t, hm.Remove(1), true)
}

func TestLinkedHashSet_Clear(t *testing.T) {
	hm := NewLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestLinkedHashSet_Empty(t *testing.T) {
	hm := NewLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}
