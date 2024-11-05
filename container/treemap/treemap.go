package treemap

import (
	"cmp"

	stlmaps "github.com/kkkunny/stl/container/maps"
)

type TreeMap[K, V any] interface {
	stlmaps.MapObj[K, V]
	Back() (K, V)
	Front() (K, V)
}

// StdWith 使用go的默认cmp函数
func StdWith[K cmp.Ordered, V any](vs ...any) TreeMap[K, V] {
	if len(vs) == 0 {
		return _NewStdTreeMap[K, V]()
	} else {
		return _NewStdTreeMapWith[K, V](vs...)
	}
}

// AnyWith 使用自定义cmp函数，全方位慢
func AnyWith[K, V any](vs ...any) TreeMap[K, V] {
	if len(vs) == 0 {
		return _NewAnyTreeMap[K, V]()
	} else {
		return _NewAnyTreeMapWith[K, V](vs...)
	}
}
