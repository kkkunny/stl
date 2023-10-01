package hashmap

import (
	"math/rand"
	"strconv"
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNewHashMap(t *testing.T) {
	v := NewHashMap[int, int]()
	stltest.AssertEq(t, v.Length(), 0)
}

func TestNewHashMapWithCapacity(t *testing.T) {
	v := NewHashMapWithCapacity[int, int](10)
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
	v := NewHashMapWith[int, int](1, 1, 2, 2, 3, 3)
	v.Clear()
	stltest.AssertEq(t, int(v.Length()), 0)
}

func TestHashMap_Empty(t *testing.T) {
	v := NewHashMap[int, int]()
	stltest.AssertEq(t, v.Empty(), true)
}

func TestHashMap_ContainKey(t *testing.T) {
	v := NewHashMap[int, int]()
	stltest.AssertEq(t, v.ContainKey(1), false)
	v.Set(1, 2)
	stltest.AssertEq(t, v.ContainKey(1), true)
}

func BenchmarkWrite_map(b *testing.B) {
	hm := make(map[string]int)
	for i := 0; i < b.N; i++ {
		hm[strconv.FormatInt(int64(i), 10)] = i
	}
}

func BenchmarkWrite_HashMap(b *testing.B) {
	hm := NewHashMap[string, int]()
	for i := 0; i < b.N; i++ {
		hm.Set(strconv.FormatInt(int64(i), 10), i)
	}
}

func BenchmarkRead_map(b *testing.B) {
	hm := make(map[string]int)
	for i := 0; i < 10000; i++ {
		hm[strconv.FormatInt(int64(i), 10)] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := rand.Int63n(10000)
		_ = hm[strconv.FormatInt(key, 10)]
	}
}

func BenchmarkRead_HashMap(b *testing.B) {
	hm := NewHashMap[string, int]()
	for i := 0; i < 10000; i++ {
		hm.Set(strconv.FormatInt(int64(i), 10), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := rand.Int63n(10000)
		_ = hm.Get(strconv.FormatInt(key, 10))
	}
}
