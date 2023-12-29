package bimap

import (
	"testing"

	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
	stltest "github.com/kkkunny/stl/test"
)

func TestBiMap_Set(t *testing.T) {
	var hm BiMap[int, int]
	stltest.AssertEq(t, pair.NewPair(hm.Set(1, 1)), pair.NewPair(0, 0))
	stltest.AssertEq(t, pair.NewPair(hm.Set(1, 2)), pair.NewPair(0, 1))
	stltest.AssertEq(t, pair.NewPair(hm.Set(2, 1)), pair.NewPair(0, 0))
}

func TestBiMap_Get(t *testing.T) {
	hm := NewBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.GetValue(1, 2), 1)
	stltest.AssertEq(t, hm.GetKey(1), 1)
	stltest.AssertEq(t, hm.GetKey(3, 2), 2)
}

func TestBiMap_Contain(t *testing.T) {
	hm := NewBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.ContainKey(1), true)
	stltest.AssertEq(t, hm.ContainKey(3), false)
	stltest.AssertEq(t, hm.ContainValue(1), true)
	stltest.AssertEq(t, hm.ContainValue(3), false)
}

func TestBiMap_Remove(t *testing.T) {
	hm := NewBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.RemoveKey(3), 0)
	stltest.AssertEq(t, hm.RemoveValue(3, 3), 3)
	stltest.AssertEq(t, hm.RemoveKey(1), 1)
}

func TestBiMap_Clear(t *testing.T) {
	hm := NewBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestBiMap_Empty(t *testing.T) {
	hm := NewBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestBiMap_Keys(t *testing.T) {
	hm := NewBiMapWith[int, int](1, 1, 2, 2)
	keys := hm.Keys()
	stltest.AssertEq(t, stliter.Contain[int](keys, 1), true)
	stltest.AssertEq(t, stliter.Contain[int](keys, 3), false)
}

func TestBiMap_Values(t *testing.T) {
	hm := NewBiMapWith[int, int](1, 1, 2, 2)
	values := hm.Values()
	stltest.AssertEq(t, stliter.Contain[int](values, 1), true)
	stltest.AssertEq(t, stliter.Contain[int](values, 3), false)
}

func TestBiMap_KeyValues(t *testing.T) {
	hm := NewBiMapWith[int, int](1, 1, 2, 2)
	pairs := hm.KeyValues()
	stltest.AssertEq(t, stliter.Contain[pair.Pair[int, int]](pairs, pair.NewPair(1, 1)), true)
	stltest.AssertEq(t, stliter.Contain[pair.Pair[int, int]](pairs, pair.NewPair(1, 2)), false)
}
