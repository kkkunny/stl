package hashmap

import (
	"math/rand"
	"testing"

	"github.com/kkkunny/stl/container/pair"
	stlslices "github.com/kkkunny/stl/container/slices"
	stlmath "github.com/kkkunny/stl/math"
	stltest "github.com/kkkunny/stl/test"
	stlval "github.com/kkkunny/stl/value"
)

func BenchmarkWrite_GenericHashMap(b *testing.B) {
	hm := _NewGenericHashMap[int, int]()
	for i := 0; i < b.N; i++ {
		hm.Set(i, i)
	}
}

func BenchmarkRead_GenericHashMap(b *testing.B) {
	hm := _NewGenericHashMap[int, int]()
	for i := 0; i < 10000; i++ {
		hm.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := rand.Int63n(10000)
		_ = hm.Get(int(key))
	}
}

func TestGenericHashMap_Capacity(t *testing.T) {
	hm := _NewGenericHashMap[int, int]()
	stltest.AssertEq(t, hm.Capacity(), initGenericHashMapCapacity)
	hm = _NewGenericHashMapWithCapacity[int, int](20)
	stltest.AssertEq(t, hm.Capacity(), stlmath.RoundToPowerOf[uint](20, 2))
}

func TestGenericHashMap_Clone(t *testing.T) {
	hm1 := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	hm2 := stlval.Clone(hm1)
	stltest.AssertEq(t, hm1, hm2)
}

func TestGenericHashMap_Equal(t *testing.T) {
	hm1 := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	hm2 := _NewGenericHashMapWith[int, int](2, 2, 1, 1)
	hm3 := _NewGenericHashMapWith[int, int](1, 2, 2, 1)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}

func TestGenericHashMap_Length(t *testing.T) {
	hm := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}

func TestGenericHashMap_Set(t *testing.T) {
	hm := _NewGenericHashMap[int, int]()
	stltest.AssertEq(t, hm.Set(1, 1), 0)
	stltest.AssertEq(t, hm.Set(1, 2), 1)
	stltest.AssertEq(t, hm.Set(2, 2), 0)
	stltest.AssertEq(t, hm.Set(2, 3), 2)
}

func TestGenericHashMap_Get(t *testing.T) {
	hm := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Get(1, 2), 1)
	stltest.AssertEq(t, hm.Get(3), 0)
	stltest.AssertEq(t, hm.Get(3, 2), 2)
}

func TestGenericHashMap_ContainKey(t *testing.T) {
	hm := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Contain(1), true)
	stltest.AssertEq(t, hm.Contain(3), false)
}

func TestGenericHashMap_Remove(t *testing.T) {
	hm := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Remove(3), 0)
	stltest.AssertEq(t, hm.Remove(3, 3), 3)
	stltest.AssertEq(t, hm.Remove(1), 1)
}

func TestGenericHashMap_Clear(t *testing.T) {
	hm := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestGenericHashMap_Empty(t *testing.T) {
	hm := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestGenericHashMap_Keys(t *testing.T) {
	hm := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	keys := hm.Keys()
	stltest.AssertEq(t, stlslices.Contain(keys, 1), true)
	stltest.AssertEq(t, stlslices.Contain(keys, 3), false)
}

func TestGenericHashMap_Values(t *testing.T) {
	hm := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	values := hm.Values()
	stltest.AssertEq(t, stlslices.Contain(values, 1), true)
	stltest.AssertEq(t, stlslices.Contain(values, 3), false)
}

func TestGenericHashMap_KeyValues(t *testing.T) {
	hm := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	pairs := hm.KeyValues()
	stltest.AssertEq(t, stlslices.Contain(pairs, pair.NewPair(1, 1)), true)
	stltest.AssertEq(t, stlslices.Contain(pairs, pair.NewPair(1, 2)), false)
}

func TestGenericHashMap_String(t *testing.T) {
	hm := _NewGenericHashMapWith[int, int](1, 1, 2, 2)
	stltest.AssertNotEq(t, hm.String(), "")
}
