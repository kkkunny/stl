package stlmaps

import (
	"math/rand"
	"time"

	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/internal/maps"
)

// Reverse 反转键值对
func Reverse[K, V comparable](m map[K]V) map[V]K {
	res := make(map[V]K, len(m))
	for k, v := range m {
		res[v] = k
	}
	return res
}

func Empty[K comparable, V any](hmap map[K]V) bool {
	return len(hmap) == 0
}

func Map[K1 comparable, V1 any, K2 comparable, V2 any](hmap map[K1]V1, f func(K1, V1) (K2, V2)) map[K2]V2 {
	res := make(map[K2]V2, len(hmap))
	for k1, v1 := range hmap {
		k2, v2 := f(k1, v1)
		res[k2] = v2
	}
	return res
}

func MapError[K1 comparable, V1 any, K2 comparable, V2 any](hmap map[K1]V1, f func(K1, V1) (K2, V2, error)) (map[K2]V2, error) {
	res := make(map[K2]V2, len(hmap))
	for k1, v1 := range hmap {
		k2, v2, err := f(k1, v1)
		if err != nil {
			return nil, err
		}
		res[k2] = v2
	}
	return res, nil
}

func FlatMap[K1 comparable, V1 any, K2 comparable, V2 any](hmap map[K1]V1, f func(K1, V1) map[K2]V2) map[K2]V2 {
	res := make(map[K2]V2, len(hmap))
	for k1, v1 := range hmap {
		hmap2 := f(k1, v1)
		for k2, v2 := range hmap2 {
			res[k2] = v2
		}
	}
	return res
}

func FlatMapError[K1 comparable, V1 any, K2 comparable, V2 any](hmap map[K1]V1, f func(K1, V1) (map[K2]V2, error)) (map[K2]V2, error) {
	res := make(map[K2]V2, len(hmap))
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

func ToSlice[K comparable, V any, T any](hmap map[K]V, mapFn func(K, V) T) []T {
	res := make([]T, len(hmap))
	var i int
	for k, v := range hmap {
		res[i] = mapFn(k, v)
		i++
	}
	return res
}

func Random[K comparable, V any](hmap map[K]V) (K, V) {
	keys := maps.Keys(hmap)
	index := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(keys))
	return keys[index], hmap[keys[index]]
}

func ContainKey[K comparable, V any](hmap map[K]V, key K) bool {
	_, ok := hmap[key]
	return ok
}

func ContainAnyKeys[K comparable, V any](hmap map[K]V, keys ...K) bool {
	return stlslices.Any(keys, func(_ int, key K) bool {
		return ContainKey(hmap, key)
	})
}

func ContainAllKeys[K comparable, V any](hmap map[K]V, keys ...K) bool {
	return stlslices.All(keys, func(_ int, key K) bool {
		return ContainKey(hmap, key)
	})
}

func Union[K comparable, V any](l, r map[K]V) map[K]V {
	hmap := make(map[K]V, len(l)+len(r))
	for k, v := range l {
		hmap[k] = v
	}
	for k, v := range r {
		hmap[k] = v
	}
	return hmap
}

// DiffKeyTo 返回l中r没有的KV
func DiffKeyTo[K comparable, V any](l, r map[K]V) map[K]V {
	res := make(map[K]V, len(l))
	for lk, lv := range l {
		if ContainKey(r, lk) {
			continue
		}
		res[lk] = lv
	}
	return res
}

// DiffKey 返回l和r中各在对方没有的KV
func DiffKey[K comparable, V any](l, r map[K]V) map[K]V {
	return Union(DiffKeyTo(l, r), DiffKeyTo(r, l))
}

func Clone[K comparable, V any](hmap map[K]V) map[K]V {
	newHMap := make(map[K]V, len(hmap))
	for k, v := range hmap {
		newHMap[k] = v
	}
	return newHMap
}
