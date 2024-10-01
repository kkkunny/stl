package bimap

import (
	"github.com/kkkunny/stl/container/hashmap"
)

type BiMap[T, E any] interface {
	hashmap.HashMap[T, E]
	Put(k T, v E) (T, E)
	GetValue(k T, defaultValue ...E) E
	GetKey(v E, defaultKey ...T) T
	ContainKey(k T) bool
	ContainValue(v E) bool
	RemoveKey(k T, defaultValue ...E) E
	RemoveValue(v E, defaultKey ...T) T
	getKeyData() hashmap.HashMap[T, E]
	getValueData() hashmap.HashMap[E, T]
}

// StdWith 使用go的默认hash函数，相比map，write更慢，read更快
func StdWith[T, E comparable](vs ...any) BiMap[T, E] {
	if len(vs) == 0 {
		return _NewStdBiMap[T, E]()
	} else {
		return _NewStdBiMapWith[T, E](vs...)
	}
}

// StdWithCap 使用go的默认hash函数，相比map，write更慢，read更快
func StdWithCap[T, E comparable](cap uint) BiMap[T, E] {
	return _NewStdBiMapWithCapacity[T, E](cap)
}

// AnyWith 使用自定义hash函数，相比map，全方位慢
func AnyWith[T, E any](vs ...any) BiMap[T, E] {
	if len(vs) == 0 {
		return _NewAnyBiMap[T, E]()
	} else {
		return _NewAnyBiMapWith[T, E](vs...)
	}
}

// AnyWithCap 使用自定义hash函数，相比map，全方位慢
func AnyWithCap[T, E any](cap uint) BiMap[T, E] {
	return _NewAnyBiMapWithCapacity[T, E](cap)
}
