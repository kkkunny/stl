package treemap

import (
	"testing"

	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/pair"
	stltest "github.com/kkkunny/stl/test"
)

func TestTreeMap_Set(t *testing.T) {
	var hm TreeMap[int, int]
	stltest.AssertEq(t, hm.Set(1, 1), 0)
	stltest.AssertEq(t, hm.Set(1, 2), 1)
	stltest.AssertEq(t, hm.Set(2, 2), 0)
	stltest.AssertEq(t, hm.Set(2, 3), 2)
}

func TestTreeMap_Get(t *testing.T) {
	hm := NewTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Get(1, 2), 1)
	stltest.AssertEq(t, hm.Get(3), 0)
	stltest.AssertEq(t, hm.Get(3, 2), 2)
}

func TestTreeMap_ContainKey(t *testing.T) {
	hm := NewTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.ContainKey(1), true)
	stltest.AssertEq(t, hm.ContainKey(3), false)
}

func TestTreeMap_Remove(t *testing.T) {
	hm := NewTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Remove(3), 0)
	stltest.AssertEq(t, hm.Remove(3, 3), 3)
	stltest.AssertEq(t, hm.Remove(1), 1)
}

func TestTreeMap_Clear(t *testing.T) {
	hm := NewTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestTreeMap_Empty(t *testing.T) {
	hm := NewTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestTreeMap_Keys(t *testing.T) {
	hm := NewTreeMapWith[int, int](2, 2, 3, 3, 1, 1)
	stltest.AssertEq(t, hm.Keys(), dynarray.NewDynArrayWith(1, 2, 3))
}

func TestTreeMap_Values(t *testing.T) {
	hm := NewTreeMapWith[int, int](2, 2, 3, 3, 1, 1)
	stltest.AssertEq(t, hm.Values(), dynarray.NewDynArrayWith(1, 2, 3))
}

func TestTreeMap_KeyValues(t *testing.T) {
	hm := NewTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.KeyValues(), dynarray.NewDynArrayWith(pair.NewPair(1, 1), pair.NewPair(2, 2)))
}

func TestTreeMap_Front(t *testing.T) {
	hm := NewTreeMapWith[int, int](1, 1, 2, 2)
	k, v := hm.Front()
	stltest.AssertEq(t, k, 1)
	stltest.AssertEq(t, v, 1)
}

func TestTreeMap_Back(t *testing.T) {
	hm := NewTreeMapWith[int, int](1, 1, 2, 2)
	k, v := hm.Back()
	stltest.AssertEq(t, k, 2)
	stltest.AssertEq(t, v, 2)
}
