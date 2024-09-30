package hashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestAnyHashSet_Clone(t *testing.T) {
	hs1 := _NewAnyHashSetWith[int](1, 2)
	hs2 := hs1.Clone()
	stltest.AssertEq(t, hs1, hs2)
}

func TestAnyHashSet_Equal(t *testing.T) {
	hs1 := _NewAnyHashSetWith[int](1)
	hs2 := _NewAnyHashSetWith[int](1)
	hs3 := _NewAnyHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs1, hs2)
	stltest.AssertNotEq(t, hs1, hs3)
	stltest.AssertNotEq(t, hs2, hs3)
}

func TestAnyHashSet_Length(t *testing.T) {
	hs := _NewAnyHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Length(), 2)
}

func TestAnyHashSet_Add(t *testing.T) {
	hs := _NewAnyHashSet[int]()
	stltest.AssertEq(t, hs.Add(1), true)
	stltest.AssertEq(t, hs.Add(1), false)
	stltest.AssertEq(t, hs.Add(2), true)
}

func TestAnyHashSet_Contain(t *testing.T) {
	hs := _NewAnyHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Contain(1), true)
	stltest.AssertEq(t, hs.Contain(3), false)
}

func TestAnyHashSet_Remove(t *testing.T) {
	hs := _NewAnyHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Remove(3), false)
	stltest.AssertEq(t, hs.Remove(1), true)
}

func TestAnyHashSet_Clear(t *testing.T) {
	hs := _NewAnyHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Length(), 2)
	hs.Clear()
	stltest.AssertEq(t, hs.Length(), 0)
}

func TestAnyHashSet_Empty(t *testing.T) {
	hs := _NewAnyHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Empty(), false)
	hs.Clear()
	stltest.AssertEq(t, hs.Empty(), true)
}

func TestAnyHashSet_String(t *testing.T) {
	hs := _NewAnyHashSetWith[int](1, 2)
	stltest.AssertNotEq(t, hs.String(), "")
}
