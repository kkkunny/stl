package set

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
	stlval "github.com/kkkunny/stl/value"
)

func TestSet_Clone(t *testing.T) {
	hs1 := StdHashSetWith[int](1, 2)
	hs2 := stlval.Clone(hs1)
	stltest.AssertEq(t, hs1, hs2)
}

func TestSet_Equal(t *testing.T) {
	hs1 := StdHashSetWith[int](1)
	hs2 := StdHashSetWith[int](1)
	hs3 := StdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs1, hs2)
	stltest.AssertNotEq(t, hs1, hs3)
	stltest.AssertNotEq(t, hs2, hs3)
}

func TestSet_Length(t *testing.T) {
	hs := StdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Length(), 2)
}

func TestSet_Add(t *testing.T) {
	hs := StdHashSetWith[int]()
	stltest.AssertEq(t, hs.Add(1), true)
	stltest.AssertEq(t, hs.Add(1), false)
	stltest.AssertEq(t, hs.Add(2), true)
}

func TestSet_Contain(t *testing.T) {
	hs := StdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Contain(1), true)
	stltest.AssertEq(t, hs.Contain(3), false)
}

func TestSet_Remove(t *testing.T) {
	hs := StdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Remove(3), false)
	stltest.AssertEq(t, hs.Remove(1), true)
}

func TestSet_Clear(t *testing.T) {
	hs := StdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Length(), 2)
	hs.Clear()
	stltest.AssertEq(t, hs.Length(), 0)
}

func TestSet_Empty(t *testing.T) {
	hs := StdHashSetWith[int](1, 2)
	stltest.AssertEq(t, hs.Empty(), false)
	hs.Clear()
	stltest.AssertEq(t, hs.Empty(), true)
}

func TestSet_String(t *testing.T) {
	hs := StdHashSetWith[int](1, 2)
	stltest.AssertNotEq(t, hs.String(), "")
}
