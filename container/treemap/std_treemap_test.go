package treemap

import (
	"math/rand"
	"testing"

	"github.com/kkkunny/stl/clone"
	"github.com/kkkunny/stl/container/tuple"
	stltest "github.com/kkkunny/stl/test"
)

func BenchmarkWrite_StdTreeMap(b *testing.B) {
	hm := _NewStdTreeMap[int, int]()
	for i := 0; i < b.N; i++ {
		hm.Set(i, i)
	}
}

func BenchmarkRead_StdTreeMap(b *testing.B) {
	hm := _NewStdTreeMap[int, int]()
	for i := 0; i < 10000; i++ {
		hm.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := rand.Int63n(10000)
		_ = hm.Get(int(key))
	}
}

func TestStdTreeMap_Clone(t *testing.T) {
	hm1 := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	hm2 := clone.Clone(hm1)
	stltest.AssertEq(t, hm1, hm2)
}

func TestStdTreeMap_Equal(t *testing.T) {
	hm1 := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	hm2 := _NewStdTreeMapWith[int, int](2, 2, 1, 1)
	hm3 := _NewStdTreeMapWith[int, int](1, 2, 2, 1)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}

func TestStdTreeMap_Length(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}

func TestStdTreeMap_Set(t *testing.T) {
	hm := _NewStdTreeMap[int, int]()
	stltest.AssertEq(t, hm.Set(1, 1), 0)
	stltest.AssertEq(t, hm.Set(1, 2), 1)
	stltest.AssertEq(t, hm.Set(2, 2), 0)
	stltest.AssertEq(t, hm.Set(2, 3), 2)
}

func TestStdTreeMap_Get(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Get(1, 2), 1)
	stltest.AssertEq(t, hm.Get(3), 0)
	stltest.AssertEq(t, hm.Get(3, 2), 2)
}

func TestStdTreeMap_ContainKey(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Contain(1), true)
	stltest.AssertEq(t, hm.Contain(3), false)
}

func TestStdTreeMap_Remove(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Remove(3), 0)
	stltest.AssertEq(t, hm.Remove(3, 3), 3)
	stltest.AssertEq(t, hm.Remove(1), 1)
}

func TestStdTreeMap_Clear(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestStdTreeMap_Empty(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestStdTreeMap_Keys(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](2, 2, 3, 3, 1, 1)
	stltest.AssertEq(t, hm.Keys(), []int{1, 2, 3})
}

func TestStdTreeMap_Values(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](2, 2, 3, 3, 1, 1)
	stltest.AssertEq(t, hm.Values(), []int{1, 2, 3})
}

func TestStdTreeMap_KeyValues(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.KeyValues(), []tuple.Tuple2[int, int]{tuple.Pack2(1, 1), tuple.Pack2(2, 2)})
}

func TestStdTreeMap_Front(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	k, v := hm.Front()
	stltest.AssertEq(t, k, 1)
	stltest.AssertEq(t, v, 1)
}

func TestStdTreeMap_Back(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](1, 1, 2, 2)
	k, v := hm.Back()
	stltest.AssertEq(t, k, 2)
	stltest.AssertEq(t, v, 2)
}

func TestStdTreeMap_String(t *testing.T) {
	hm := _NewStdTreeMapWith[int, int](2, 2, 1, 1)
	stltest.AssertEq(t, hm.String(), "TreeMap{1: 1, 2: 2}")
}
