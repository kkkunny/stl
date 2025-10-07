package hashmap

import (
	"encoding/json"
	"iter"

	stlmaps "github.com/kkkunny/stl/container/maps"
)

type HashMap[K, V any] interface {
	stlmaps.MapObj[K, V]
	Iter2() iter.Seq2[K, V]
	json.Marshaler
	json.Unmarshaler
}

// StdWith 使用go的默认hash函数，相比map，write更慢，read更快
func StdWith[K comparable, V any](vs ...any) HashMap[K, V] {
	if len(vs) == 0 {
		return _NewSwissTable[K, V]()
	} else {
		return _NewSwissTableWith[K, V](vs...)
	}
}

// StdWithCap 使用go的默认hash函数，相比map，write更慢，read更快
func StdWithCap[K comparable, V any](cap uint) HashMap[K, V] {
	return _NewSwissTableWithCapacity[K, V](cap)
}

// AnyWith 使用自定义hash函数，相比map，全方位慢
func AnyWith[K, V any](vs ...any) HashMap[K, V] {
	if len(vs) == 0 {
		return _NewGenericHashMap[K, V]()
	} else {
		return _NewGenericHashMapWith[K, V](vs...)
	}
}

// AnyWithCap 使用自定义hash函数，相比map，全方位慢
func AnyWithCap[K, V any](cap uint) HashMap[K, V] {
	return _NewGenericHashMapWithCapacity[K, V](cap)
}

// ThreadSafeStdWith 使用go的默认hash函数，相比map，write更慢，read更快
func ThreadSafeStdWith[K comparable, V any](vs ...any) HashMap[K, V] {
	if len(vs) == 0 {
		return _NewConcurrentMap[K, V]()
	} else {
		return _NewConcurrentMapWith[K, V](vs...)
	}
}
