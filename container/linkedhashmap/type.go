package linkedhashmap

import (
	"github.com/kkkunny/stl/container/hashmap"
	"github.com/kkkunny/stl/container/pair"
	"github.com/kkkunny/stl/internal/list"
)

// LinkedHashMap 插入顺序哈希表
type LinkedHashMap[K, V any] struct {
	hashmap.HashMap[K, *list.Element[pair.Pair[K, V]]]
	list *list.List[pair.Pair[K, V]]
}

func NewLinkedHashMap[K, V any]() LinkedHashMap[K, V] {
	return LinkedHashMap[K, V]{
		HashMap: hashmap.NewHashMap[K, *list.Element[pair.Pair[K, V]]](),
		list:    list.New[pair.Pair[K, V]](),
	}
}

func NewLinkedHashMapWithCapacity[K, V any](cap uint) LinkedHashMap[K, V] {
	return LinkedHashMap[K, V]{
		HashMap: hashmap.NewHashMapWithCapacity[K, *list.Element[pair.Pair[K, V]]](cap),
		list:    list.New[pair.Pair[K, V]](),
	}
}

func NewLinkedHashMapWith[K, V any](vs ...any) LinkedHashMap[K, V] {
	self := NewLinkedHashMapWithCapacity[K, V](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}
