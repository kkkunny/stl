package treemap

import (
	"math/rand"
	"testing"

	"github.com/kkkunny/stl/container/pair"
	stltest "github.com/kkkunny/stl/test"
	stlval "github.com/kkkunny/stl/value"
)

func BenchmarkWrite_AnyTreeMap(b *testing.B) {
	hm := _NewAnyTreeMap[int, int]()
	for i := 0; i < b.N; i++ {
		hm.Set(i, i)
	}
}

func BenchmarkRead_AnyTreeMap(b *testing.B) {
	hm := _NewAnyTreeMap[int, int]()
	for i := 0; i < 10000; i++ {
		hm.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := rand.Int63n(10000)
		_ = hm.Get(int(key))
	}
}

func TestAnyTreeMap_Clone(t *testing.T) {
	hm1 := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	hm2 := stlval.Clone(hm1)
	stltest.AssertEq(t, hm1, hm2)
}

func TestAnyTreeMap_Equal(t *testing.T) {
	hm1 := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	hm2 := _NewAnyTreeMapWith[int, int](2, 2, 1, 1)
	hm3 := _NewAnyTreeMapWith[int, int](1, 2, 2, 1)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}

func TestAnyTreeMap_Length(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}

func TestAnyTreeMap_Set(t *testing.T) {
	hm := _NewAnyTreeMap[int, int]()
	stltest.AssertEq(t, hm.Set(1, 1), 0)
	stltest.AssertEq(t, hm.Set(1, 2), 1)
	stltest.AssertEq(t, hm.Set(2, 2), 0)
	stltest.AssertEq(t, hm.Set(2, 3), 2)
}

func TestAnyTreeMap_Get(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Get(1, 2), 1)
	stltest.AssertEq(t, hm.Get(3), 0)
	stltest.AssertEq(t, hm.Get(3, 2), 2)
}

func TestAnyTreeMap_ContainKey(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Contain(1), true)
	stltest.AssertEq(t, hm.Contain(3), false)
}

func TestAnyTreeMap_Remove(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Remove(3), 0)
	stltest.AssertEq(t, hm.Remove(3, 3), 3)
	stltest.AssertEq(t, hm.Remove(1), 1)
}

func TestAnyTreeMap_Clear(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestAnyTreeMap_Empty(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestAnyTreeMap_Keys(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](2, 2, 3, 3, 1, 1)
	stltest.AssertEq(t, hm.Keys(), []int{1, 2, 3})
}

func TestAnyTreeMap_Values(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](2, 2, 3, 3, 1, 1)
	stltest.AssertEq(t, hm.Values(), []int{1, 2, 3})
}

func TestAnyTreeMap_KeyValues(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.KeyValues(), []pair.Pair[int, int]{pair.NewPair(1, 1), pair.NewPair(2, 2)})
}

func TestAnyTreeMap_Front(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	k, v := hm.Front()
	stltest.AssertEq(t, k, 1)
	stltest.AssertEq(t, v, 1)
}

func TestAnyTreeMap_Back(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](1, 1, 2, 2)
	k, v := hm.Back()
	stltest.AssertEq(t, k, 2)
	stltest.AssertEq(t, v, 2)
}

func TestAnyTreeMap_String(t *testing.T) {
	hm := _NewAnyTreeMapWith[int, int](2, 2, 1, 1)
	stltest.AssertEq(t, hm.String(), "TreeMap{1: 1, 2: 2}")
}
