package treemap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewTreeMap(t *testing.T) {
	v := NewTreeMap[int, int]()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestNewTreeMapWith(t *testing.T) {
	v := NewTreeMapWith[int, int](1, 1, 2, 2, 3, 3)
	stltest.AssertEq(t, v.Length(), 3)
}

func TestTreeMap_Get(t *testing.T) {
	v := NewTreeMapWith[int, int](1, 1, 2, 2, 3, 3)
	stltest.AssertEq(t, v.Get(1), 1)
	stltest.AssertEq(t, v.Get(2), 2)
	stltest.AssertEq(t, v.Get(3), 3)
	stltest.AssertEq(t, v.Get(4), 0)
}

func TestTreeMap_Set(t *testing.T) {
	v := NewTreeMapWith[int, int](1, 0, 2, 1, 3, 2)
	stltest.AssertEq(t, v.Set(1, 1), 0)
	stltest.AssertEq(t, v.Get(1), 1)
	stltest.AssertEq(t, v.Set(4, 4), 0)
	stltest.AssertEq(t, v.Get(4), 4)
}

func TestTreeMap_Remove(t *testing.T) {
	v := NewTreeMapWith[int, int](1, 1, 2, 2, 3, 3)
	stltest.AssertEq(t, v.Remove(1), 1)
	stltest.AssertEq(t, v.Length(), 2)
}

func TestTreeMap_Clear(t *testing.T) {
	v := NewTreeMapWith[int, int](1, 1, 2, 2, 3, 3)
	v.Clear()
	stltest.AssertEq(t, int(v.Length()), 0)
}

func TestTreeMap_Empty(t *testing.T) {
	v := NewTreeMap[int, int]()
	stltest.AssertEq(t, v.Empty(), true)
}

func TestTreeMap_ContainKey(t *testing.T) {
	v := NewTreeMap[int, int]()
	stltest.AssertEq(t, v.ContainKey(1), false)
	v.Set(1, 2)
	stltest.AssertEq(t, v.ContainKey(1), true)
}
