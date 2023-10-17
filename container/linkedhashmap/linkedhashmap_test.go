package linkedhashmap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewLinkedHashMap(t *testing.T) {
	v := NewLinkedHashMap[int, int]()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestNewLinkedHashMapWithCapacity(t *testing.T) {
	v := NewLinkedHashMapWithCapacity[int, int](10)
	stltest.AssertEq(t, v.Length(), 0)
}

func TestNewLinkedHashMapWith(t *testing.T) {
	v := NewLinkedHashMapWith[int, int](1, 1, 2, 2, 3, 3)
	stltest.AssertEq(t, v.Length(), 3)
}

func TestLinkedHashMap_Get(t *testing.T) {
	v := NewLinkedHashMapWith[int, int](1, 1, 2, 2, 3, 3)
	stltest.AssertEq(t, v.Get(1), 1)
	stltest.AssertEq(t, v.Get(2), 2)
	stltest.AssertEq(t, v.Get(3), 3)
	stltest.AssertEq(t, v.Get(4), 0)
}

func TestLinkedHashMap_Set(t *testing.T) {
	v := NewLinkedHashMapWith[int, int](1, 0, 2, 1, 3, 2)
	stltest.AssertEq(t, v.Set(1, 1), 0)
	stltest.AssertEq(t, v.Get(1), 1)
	stltest.AssertEq(t, v.Set(4, 4), 0)
	stltest.AssertEq(t, v.Get(4), 4)
}

func TestLinkedHashMap_Remove(t *testing.T) {
	v := NewLinkedHashMapWith[int, int](1, 1, 2, 2, 3, 3)
	stltest.AssertEq(t, v.Remove(1), 1)
	stltest.AssertEq(t, v.Length(), 2)
}

func TestLinkedHashMap_Clear(t *testing.T) {
	v := NewLinkedHashMapWith[int, int](1, 1, 2, 2, 3, 3)
	v.Clear()
	stltest.AssertEq(t, int(v.Length()), 0)
}

func TestLinkedHashMap_Empty(t *testing.T) {
	v := NewLinkedHashMap[int, int]()
	stltest.AssertEq(t, v.Empty(), true)
}

func TestLinkedHashMap_ContainKey(t *testing.T) {
	v := NewLinkedHashMap[int, int]()
	stltest.AssertEq(t, v.ContainKey(1), false)
	v.Set(1, 2)
	stltest.AssertEq(t, v.ContainKey(1), true)
}
