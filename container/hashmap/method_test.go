package hashmap

import (
	"testing"

	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
	stltest "github.com/kkkunny/stl/test"
)

func TestHashMap_Set(t *testing.T) {
	var hm HashMap[int, int]
	stltest.AssertEq(t, hm.Set(1, 1), 0)
	stltest.AssertEq(t, hm.Set(1, 2), 1)
	stltest.AssertEq(t, hm.Set(2, 2), 0)
	stltest.AssertEq(t, hm.Set(2, 3), 2)
}

func TestHashMap_Get(t *testing.T) {
	hm := NewHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Get(1, 2), 1)
	stltest.AssertEq(t, hm.Get(3), 0)
	stltest.AssertEq(t, hm.Get(3, 2), 2)
}

func TestHashMap_ContainKey(t *testing.T) {
	hm := NewHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.ContainKey(1), true)
	stltest.AssertEq(t, hm.ContainKey(3), false)
}

func TestHashMap_Remove(t *testing.T) {
	hm := NewHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Remove(3), 0)
	stltest.AssertEq(t, hm.Remove(3, 3), 3)
	stltest.AssertEq(t, hm.Remove(1), 1)
}

func TestHashMap_Clear(t *testing.T) {
	hm := NewHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestHashMap_Empty(t *testing.T) {
	hm := NewHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestHashMap_Keys(t *testing.T) {
	hm := NewHashMapWith[int, int](1, 1, 2, 2)
	keys := hm.Keys()
	stltest.AssertEq(t, stliter.Contain[int](keys, 1), true)
	stltest.AssertEq(t, stliter.Contain[int](keys, 3), false)
}

func TestHashMap_Values(t *testing.T) {
	hm := NewHashMapWith[int, int](1, 1, 2, 2)
	values := hm.Values()
	stltest.AssertEq(t, stliter.Contain[int](values, 1), true)
	stltest.AssertEq(t, stliter.Contain[int](values, 3), false)
}

func TestHashMap_KeyValues(t *testing.T) {
	hm := NewHashMapWith[int, int](1, 1, 2, 2)
	pairs := hm.KeyValues()
	stltest.AssertEq(t, stliter.Contain[pair.Pair[int, int]](pairs, pair.NewPair(1, 1)), true)
	stltest.AssertEq(t, stliter.Contain[pair.Pair[int, int]](pairs, pair.NewPair(1, 2)), false)
}
