package hashmap

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/kkkunny/stl/clone"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/container/tuple"
	stltest "github.com/kkkunny/stl/test"
)

func BenchmarkWrite_ConcurrentMap(b *testing.B) {
	hm := _NewConcurrentMap[int, int]()
	for i := 0; i < b.N; i++ {
		hm.Set(i, i)
	}
}

func BenchmarkRead_ConcurrentMap(b *testing.B) {
	hm := _NewConcurrentMap[int, int]()
	for i := 0; i < 10000; i++ {
		hm.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := rand.Int63n(10000)
		_ = hm.Get(int(key))
	}
}

func TestConcurrentMap_Clone(t *testing.T) {
	hm1 := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	hm2 := clone.Clone(hm1)
	stltest.AssertEq(t, hm1, hm2)
}

func TestConcurrentMap_Equal(t *testing.T) {
	hm1 := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	hm2 := _NewConcurrentMapWith[int, int](2, 2, 1, 1)
	hm3 := _NewConcurrentMapWith[int, int](1, 2, 2, 1)
	fmt.Println(hm1, hm2, hm3)
	stltest.AssertEq(t, hm1, hm2)
	stltest.AssertNotEq(t, hm1, hm3)
	stltest.AssertNotEq(t, hm2, hm3)
}

func TestConcurrentMap_Length(t *testing.T) {
	hm := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
}

func TestConcurrentMap_Set(t *testing.T) {
	hm := _NewConcurrentMap[int, int]()
	stltest.AssertEq(t, hm.Set(1, 1), 0)
	stltest.AssertEq(t, hm.Set(1, 2), 1)
	stltest.AssertEq(t, hm.Set(2, 2), 0)
	stltest.AssertEq(t, hm.Set(2, 3), 2)
}

func TestConcurrentMap_Get(t *testing.T) {
	hm := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Get(1, 2), 1)
	stltest.AssertEq(t, hm.Get(3), 0)
	stltest.AssertEq(t, hm.Get(3, 2), 2)
}

func TestConcurrentMap_ContainKey(t *testing.T) {
	hm := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Contain(1), true)
	stltest.AssertEq(t, hm.Contain(3), false)
}

func TestConcurrentMap_Remove(t *testing.T) {
	hm := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Remove(3), 0)
	stltest.AssertEq(t, hm.Remove(3, 3), 3)
	stltest.AssertEq(t, hm.Remove(1), 1)
}

func TestConcurrentMap_Clear(t *testing.T) {
	hm := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Length(), 2)
	hm.Clear()
	stltest.AssertEq(t, hm.Length(), 0)
}

func TestConcurrentMap_Empty(t *testing.T) {
	hm := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	stltest.AssertEq(t, hm.Empty(), false)
	hm.Clear()
	stltest.AssertEq(t, hm.Empty(), true)
}

func TestConcurrentMap_Keys(t *testing.T) {
	hm := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	keys := hm.Keys()
	stltest.AssertEq(t, stlslices.Contain(keys, 1), true)
	stltest.AssertEq(t, stlslices.Contain(keys, 3), false)
}

func TestConcurrentMap_Values(t *testing.T) {
	hm := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	values := hm.Values()
	stltest.AssertEq(t, stlslices.Contain(values, 1), true)
	stltest.AssertEq(t, stlslices.Contain(values, 3), false)
}

func TestConcurrentMap_KeyValues(t *testing.T) {
	hm := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	pairs := hm.KeyValues()
	stltest.AssertEq(t, stlslices.Contain(pairs, tuple.Pack2(1, 1)), true)
	stltest.AssertEq(t, stlslices.Contain(pairs, tuple.Pack2(1, 2)), false)
}

func TestConcurrentMap_String(t *testing.T) {
	hm := _NewConcurrentMapWith[int, int](1, 1, 2, 2)
	stltest.AssertNotEq(t, hm.String(), "")
}
