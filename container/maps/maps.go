package stlmaps

import (
	"math/rand"
	"time"

	"golang.org/x/exp/maps"
)

// Reverse 反转键值对
func Reverse[K, V comparable, KV ~map[K]V, VK ~map[V]K](m KV) VK {
	res := make(map[V]K, len(m))
	for k, v := range m {
		res[v] = k
	}
	return res
}

func Empty[K comparable, V any, KV ~map[K]V](hmap KV) bool {
	return len(hmap) == 0
}

func Map[K1 comparable, V1 any, KV1 ~map[K1]V1, K2 comparable, V2 any, KV2 ~map[K2]V2](hmap KV1, f func(K1, V1) (K2, V2)) KV2 {
	res := make(KV2, len(hmap))
	for k1, v1 := range hmap {
		k2, v2 := f(k1, v1)
		res[k2] = v2
	}
	return res
}

func MapError[K1 comparable, V1 any, KV1 ~map[K1]V1, K2 comparable, V2 any, KV2 ~map[K2]V2](hmap KV1, f func(K1, V1) (K2, V2, error)) (KV2, error) {
	res := make(KV2, len(hmap))
	for k1, v1 := range hmap {
		k2, v2, err := f(k1, v1)
		if err != nil {
			return nil, err
		}
		res[k2] = v2
	}
	return res, nil
}

func FlatMap[K1 comparable, V1 any, KV1 ~map[K1]V1, K2 comparable, V2 any, KV2 ~map[K2]V2](hmap KV1, f func(K1, V1) KV2) KV2 {
	res := make(KV2, len(hmap))
	for k1, v1 := range hmap {
		hmap2 := f(k1, v1)
		for k2, v2 := range hmap2 {
			res[k2] = v2
		}
	}
	return res
}

func FlatMapError[K1 comparable, V1 any, KV1 ~map[K1]V1, K2 comparable, V2 any, KV2 ~map[K2]V2](hmap KV1, f func(K1, V1) (KV2, error)) (KV2, error) {
	res := make(KV2, len(hmap))
	for k1, v1 := range hmap {
		hmap2, err := f(k1, v1)
		if err != nil {
			return nil, err
		}
		for k2, v2 := range hmap2 {
			res[k2] = v2
		}
	}
	return res, nil
}

func ToSlice[K comparable, V any, KV ~map[K]V, T any, TS ~[]T](hmap KV, mapFn func(K, V) T) TS {
	res := make(TS, len(hmap))
	var i int
	for k, v := range hmap {
		res[i] = mapFn(k, v)
		i++
	}
	return res
}

func Random[K comparable, V any, KV ~map[K]V](hmap KV) (K, V) {
	keys := maps.Keys(hmap)
	index := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(keys))
	return keys[index], hmap[keys[index]]
}
