package hashmap

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewHashMap(t *testing.T) {
	v := NewHashMap[int, int]()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestNewHashMapWithCapacity(t *testing.T) {
	v := NewHashMap[int, int]()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestNewHashMapWith(t *testing.T) {
	v := NewHashMapWith[int, int](1, 1, 2, 2, 3, 3)
	stltest.AssertEq(t, v.Length(), 3)
}

func TestHashMap_Get(t *testing.T) {
	v := NewHashMapWith[int, int](1, 1, 2, 2, 3, 3)
	stltest.AssertEq(t, v.Get(1), 1)
	stltest.AssertEq(t, v.Get(2), 2)
	stltest.AssertEq(t, v.Get(3), 3)
	stltest.AssertEq(t, v.Get(4), 0)
}

func TestHashMap_Set(t *testing.T) {
	v := NewHashMapWith[int, int](1, 0, 2, 1, 3, 2)
	stltest.AssertEq(t, v.Set(1, 1), 0)
	stltest.AssertEq(t, v.Get(1), 1)
	stltest.AssertEq(t, v.Set(4, 4), 0)
	stltest.AssertEq(t, v.Get(4), 4)
}

func TestHashMap_Remove(t *testing.T) {
	v := NewHashMapWith[int, int](1, 1, 2, 2, 3, 3)
	stltest.AssertEq(t, v.Remove(1), 1)
	stltest.AssertEq(t, v.Length(), 2)
}

func TestHashMap_Clear(t *testing.T) {
	v := NewHashMap[int, int]()
	v.Clear()
	stltest.AssertEq(t, int(v.Length()), 0)
}

func TestHashMap_Empty(t *testing.T) {
	v := NewHashMap[int, int]()
	stltest.AssertEq(t, v.Empty(), true)
}
