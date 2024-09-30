package linkedhashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestAnyLinkedHashSet_Clone(t *testing.T) {
	hm1 := _NewAnyLinkedHashSetWith[int](1, 2)
	hm2 := hm1.Clone()
	stltest.AssertEq(t, hm1, hm2)
}

func TestAnyLinkedHashSet_Equal(t *testing.T) {
	hm1 := _NewAnyLinkedHashSetWith[int](1)
	hm2 := _NewAnyLinkedHashSetWith[int](1)
	hm3 := _NewAnyLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}

func TestAnyLinkedHashSet_Length(t *testing.T) {
	hm := _NewAnyLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}

func TestAnyLinkedHashSet_Add(t *testing.T) {
	hs := _NewAnyLinkedHashSet[int]()
	stltest.AssertEq(t, hs.Add(1), true)
	stltest.AssertEq(t, hs.Add(1), false)
	stltest.AssertEq(t, hs.Add(2), true)
}

func TestAnyLinkedHashSet_Contain(t *testing.T) {
	hm := _NewAnyLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Contain(1), true)
	stltest.AssertEq(t, hm.Contain(3), false)
}

func TestAnyLinkedHashSet_Remove(t *testing.T) {
	hm := _NewAnyLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Remove(3), false)
	stltest.AssertEq(t, hm.Remove(1), true)
}

func TestAnyLinkedHashSet_Clear(t *testing.T) {
	hm := _NewAnyLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestAnyLinkedHashSet_Empty(t *testing.T) {
	hm := _NewAnyLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestAnyLinkedHashSet_String(t *testing.T) {
	hm := _NewAnyLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.String(), "LinkedHashSet{1, 2}")
}
