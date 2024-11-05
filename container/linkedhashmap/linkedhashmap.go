package linkedhashmap

import (
	"github.com/kkkunny/stl/container/hashmap"
	"github.com/kkkunny/stl/container/tuple"
	"github.com/kkkunny/stl/list"
)

type LinkedHashMap[K, V any] interface {
	hashmap.HashMap[K, V]
	getList() *list.List[tuple.Tuple2[K, V]]
}

// StdWith 使用go的默认hash函数，相比map，write更慢，read更快
func StdWith[K comparable, V any](vs ...any) LinkedHashMap[K, V] {
	if len(vs) == 0 {
		return _NewStdLinkedHashMap[K, V]()
	} else {
		return _NewStdLinkedHashMapWith[K, V](vs...)
	}
}

// StdWithCap 使用go的默认hash函数，相比map，write更慢，read更快
func StdWithCap[K comparable, V any](cap uint) LinkedHashMap[K, V] {
	return _NewStdLinkedHashMapWithCapacity[K, V](cap)
}

// AnyWith 使用自定义hash函数，相比map，全方位慢
func AnyWith[K, V any](vs ...any) LinkedHashMap[K, V] {
	if len(vs) == 0 {
		return _NewAnyLinkedHashMap[K, V]()
	} else {
		return _NewAnyLinkedHashMapWith[K, V](vs...)
	}
}

// AnyWithCap 使用自定义hash函数，相比map，全方位慢
func AnyWithCap[K, V any](cap uint) LinkedHashMap[K, V] {
	return _NewAnyLinkedHashMapWithCapacity[K, V](cap)
}
