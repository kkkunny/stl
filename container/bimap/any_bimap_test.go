package bimap

import (
	"testing"

	"github.com/kkkunny/stl/container/pair"
	stlslices "github.com/kkkunny/stl/container/slices"
	stltest "github.com/kkkunny/stl/test"
)

func TestAnyBiMap_Clone(t *testing.T) {
	hm1 := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	hm2 := hm1.Clone()
	stltest.AssertEq(t, hm1, hm2)
}

func TestAnyBiMap_Equal(t *testing.T) {
	hm1 := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	hm2 := _NewAnyBiMapWith[int, int](2, 2, 1, 1)
	hm3 := _NewAnyBiMapWith[int, int](1, 2, 2, 1)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}

func TestAnyBiMap_Length(t *testing.T) {
	hm := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}

func TestAnyBiMap_Set(t *testing.T) {
	hm := _NewAnyBiMap[int, int]()
	stltest.AssertEq(t, pair.NewPair(hm.Set(1, 1)), pair.NewPair(0, 0))
	stltest.AssertEq(t, pair.NewPair(hm.Set(1, 2)), pair.NewPair(0, 1))
	stltest.AssertEq(t, pair.NewPair(hm.Set(2, 1)), pair.NewPair(0, 0))
}

func TestAnyBiMap_Get(t *testing.T) {
	hm := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.GetValue(1, 2), 1)
	stltest.AssertEq(t, hm.GetKey(1), 1)
	stltest.AssertEq(t, hm.GetKey(3, 2), 2)
}

func TestAnyBiMap_Contain(t *testing.T) {
	hm := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.ContainKey(1), true)
	stltest.AssertEq(t, hm.ContainKey(3), false)
	stltest.AssertEq(t, hm.ContainValue(1), true)
	stltest.AssertEq(t, hm.ContainValue(3), false)
}

func TestAnyBiMap_Remove(t *testing.T) {
	hm := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.RemoveKey(3), 0)
	stltest.AssertEq(t, hm.RemoveValue(3, 3), 3)
	stltest.AssertEq(t, hm.RemoveKey(1), 1)
}

func TestAnyBiMap_Clear(t *testing.T) {
	hm := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestAnyBiMap_Empty(t *testing.T) {
	hm := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestAnyBiMap_Keys(t *testing.T) {
	hm := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	keys := hm.Keys()
	stltest.AssertEq(t, stlslices.Contain[int](keys, 1), true)
	stltest.AssertEq(t, stlslices.Contain[int](keys, 3), false)
}

func TestAnyBiMap_Values(t *testing.T) {
	hm := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	values := hm.Values()
	stltest.AssertEq(t, stlslices.Contain[int](values, 1), true)
	stltest.AssertEq(t, stlslices.Contain[int](values, 3), false)
}

func TestAnyBiMap_KeyValues(t *testing.T) {
	hm := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	pairs := hm.KeyValues()
	stltest.AssertEq(t, stlslices.Contain[pair.Pair[int, int]](pairs, pair.NewPair(1, 1)), true)
	stltest.AssertEq(t, stlslices.Contain[pair.Pair[int, int]](pairs, pair.NewPair(1, 2)), false)
}

func TestAnyBiMap_String(t *testing.T) {
	hm := _NewAnyBiMapWith[int, int](1, 1, 2, 2)
	stltest.AssertNotEq(t, hm.String(), "")
}
