package hashset

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestStdHashSet_Clone(t *testing.T) {
	hs1 := _NewStdHashSetWith[int](1, 2)
	hs2 := hs1.Clone()
	stltest.AssertEq(t, hs1, hs2)
}

func TestStdHashSet_Equal(t *testing.T) {
	hs1 := _NewStdHashSetWith[int](1)
	hs2 := _NewStdHashSetWith[int](1)
	hs3 := _NewStdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs1, hs2)
	stltest.AssertNotEq(t, hs1, hs3)
	stltest.AssertNotEq(t, hs2, hs3)
}

func TestStdHashSet_Length(t *testing.T) {
	hs := _NewStdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Length(), 2)
}

func TestStdHashSet_Add(t *testing.T) {
	hs := _NewStdHashSet[int]()
	stltest.AssertEq(t, hs.Add(1), true)
	stltest.AssertEq(t, hs.Add(1), false)
	stltest.AssertEq(t, hs.Add(2), true)
}

func TestStdHashSet_Contain(t *testing.T) {
	hs := _NewStdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Contain(1), true)
	stltest.AssertEq(t, hs.Contain(3), false)
}

func TestStdHashSet_Remove(t *testing.T) {
	hs := _NewStdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Remove(3), false)
	stltest.AssertEq(t, hs.Remove(1), true)
}

func TestStdHashSet_Clear(t *testing.T) {
	hs := _NewStdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Length(), 2)
	hs.Clear()
	stltest.AssertEq(t, hs.Length(), 0)
}

func TestStdHashSet_Empty(t *testing.T) {
	hs := _NewStdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Empty(), false)
	hs.Clear()
	stltest.AssertEq(t, hs.Empty(), true)
}

func TestStdHashSet_String(t *testing.T) {
	hs := _NewStdHashSetWith[int](1, 2)
	stltest.AssertNotEq(t, hs.String(), "")
}
