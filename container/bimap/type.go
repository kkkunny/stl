package bimap

import "github.com/kkkunny/stl/container/hashmap"

// BiMap 双向哈希表
type BiMap[K, V any] struct {
	keys   hashmap.HashMap[K, V]
	values hashmap.HashMap[V, K]
}

func NewBiMap[K, V any]() BiMap[K, V] {
	return BiMap[K, V]{
		keys:   hashmap.NewHashMap[K, V](),
		values: hashmap.NewHashMap[V, K](),
	}
}

func NewBiMapWithCapacity[K, V any](cap uint) BiMap[K, V] {
	return BiMap[K, V]{
		keys:   hashmap.NewHashMapWithCapacity[K, V](cap),
		values: hashmap.NewHashMapWithCapacity[V, K](cap),
	}
}

func NewBiMapWith[K, V any](vs ...any) BiMap[K, V] {
	self := NewBiMapWithCapacity[K, V](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}
