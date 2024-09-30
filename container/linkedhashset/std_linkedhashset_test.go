package linkedhashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestStdLinkedHashSet_Clone(t *testing.T) {
	hm1 := _NewStdLinkedHashSetWith[int](1, 2)
	hm2 := hm1.Clone()
	stltest.AssertEq(t, hm1, hm2)
}

func TestStdLinkedHashSet_Equal(t *testing.T) {
	hm1 := _NewStdLinkedHashSetWith[int](1)
	hm2 := _NewStdLinkedHashSetWith[int](1)
	hm3 := _NewStdLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}

func TestStdLinkedHashSet_Length(t *testing.T) {
	hm := _NewStdLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}

func TestStdLinkedHashSet_Add(t *testing.T) {
	hs := _NewStdLinkedHashSet[int]()
	stltest.AssertEq(t, hs.Add(1), true)
	stltest.AssertEq(t, hs.Add(1), false)
	stltest.AssertEq(t, hs.Add(2), true)
}

func TestStdLinkedHashSet_Contain(t *testing.T) {
	hm := _NewStdLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Contain(1), true)
	stltest.AssertEq(t, hm.Contain(3), false)
}

func TestStdLinkedHashSet_Remove(t *testing.T) {
	hm := _NewStdLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Remove(3), false)
	stltest.AssertEq(t, hm.Remove(1), true)
}

func TestStdLinkedHashSet_Clear(t *testing.T) {
	hm := _NewStdLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestStdLinkedHashSet_Empty(t *testing.T) {
	hm := _NewStdLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestStdLinkedHashSet_String(t *testing.T) {
	hm := _NewStdLinkedHashSetWith[int](1, 2)
	stltest.AssertEq(t, hm.String(), "LinkedHashSet{1, 2}")
}
