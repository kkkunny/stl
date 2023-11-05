package linkedhashmap

import (
	"testing"

	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/pair"
	stltest "github.com/kkkunny/stl/test"
)

func TestLinkedHashMap_Set(t *testing.T) {
	var hm LinkedHashMap[int, int]
	stltest.AssertEq(t, hm.Set(1, 1), 0)
	stltest.AssertEq(t, hm.Set(1, 2), 1)
	stltest.AssertEq(t, hm.Set(2, 2), 0)
	stltest.AssertEq(t, hm.Set(2, 3), 2)
}

func TestLinkedHashMap_Get(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Get(1, 2), 1)
	stltest.AssertEq(t, hm.Get(3), 0)
	stltest.AssertEq(t, hm.Get(3, 2), 2)
}

func TestLinkedHashMap_ContainKey(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.ContainKey(1), true)
	stltest.AssertEq(t, hm.ContainKey(3), false)
}

func TestLinkedHashMap_Remove(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Remove(3), 0)
	stltest.AssertEq(t, hm.Remove(3, 3), 3)
	stltest.AssertEq(t, hm.Remove(1), 1)
}

func TestLinkedHashMap_Clear(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestLinkedHashMap_Empty(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestLinkedHashMap_Keys(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Keys(), dynarray.NewDynArrayWith(1, 2))
}

func TestLinkedHashMap_Values(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Values(), dynarray.NewDynArrayWith(1, 2))
}

func TestLinkedHashMap_KeyValues(t *testing.T) {
	hm := NewLinkedHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.KeyValues(), dynarray.NewDynArrayWith(pair.NewPair(1, 1), pair.NewPair(2, 2)))
}
