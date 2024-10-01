package treeset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestStdTreeSet_Clone(t *testing.T) {
	hm1 := _NewStdTreeSetWith[int](1, 2)
	hm2 := hm1.Clone()
	stltest.AssertEq(t, hm1, hm2)
}

func TestStdTreeSet_Equal(t *testing.T) {
	hm1 := _NewStdTreeSetWith[int](1)
	hm2 := _NewStdTreeSetWith[int](1)
	hm3 := _NewStdTreeSetWith[int](1, 2)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}

func TestStdTreeSet_Length(t *testing.T) {
	hm := _NewStdTreeSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}

func TestStdTreeSet_Add(t *testing.T) {
	hs := _NewStdTreeSet[int]()
	stltest.AssertEq(t, hs.Add(1), true)
	stltest.AssertEq(t, hs.Add(1), false)
	stltest.AssertEq(t, hs.Add(2), true)
}

func TestStdTreeSet_Contain(t *testing.T) {
	hm := _NewStdTreeSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Contain(1), true)
	stltest.AssertEq(t, hm.Contain(3), false)
}

func TestStdTreeSet_Remove(t *testing.T) {
	hm := _NewStdTreeSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Remove(3), false)
	stltest.AssertEq(t, hm.Remove(1), true)
}

func TestStdTreeSet_Clear(t *testing.T) {
	hm := _NewStdTreeSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestStdTreeSet_Empty(t *testing.T) {
	hm := _NewStdTreeSetWith[int](1, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestStdTreeSet_String(t *testing.T) {
	hm := _NewStdTreeSetWith[int](2, 1)
	stltest.AssertEq(t, hm.String(), "TreeSet{1, 2}")
}
