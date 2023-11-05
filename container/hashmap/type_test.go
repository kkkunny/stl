package hashmap

import (
	"math/rand"
	"testing"
)

func BenchmarkWrite_map(b *testing.B) {
	hm := make(map[int]int)
	for i := 0; i < b.N; i++ {
		hm[i] = i
	}
}

func BenchmarkWrite_HashMap(b *testing.B) {
	hm := NewHashMap[int, int]()
	for i := 0; i < b.N; i++ {
		hm.Set(i, i)
	}
}

func BenchmarkRead_map(b *testing.B) {
	hm := make(map[int]int)
	for i := 0; i < 10000; i++ {
		hm[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := rand.Int63n(10000)
		_ = hm[int(key)]
	}
}

func BenchmarkRead_HashMap(b *testing.B) {
	hm := NewHashMap[int, int]()
	for i := 0; i < 10000; i++ {
		hm.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := rand.Int63n(10000)
		_ = hm.Get(int(key))
	}
}
