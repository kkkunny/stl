package treemap

import (
	"cmp"
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
)

type TreeMap[K, V any] interface {
	stlbasic.Cloneable[TreeMap[K, V]]
	stlcmp.Equalable[TreeMap[K, V]]
	stliter.IteratorContainer[pair.Pair[K, V]]
	stlbasic.Lengthable
	Set(k K, v V) V
	Get(k K, defaultValue ...V) V
	ContainKey(k K) bool
	Remove(k K, defaultValue ...V) V
	Clear()
	Empty() bool
	Keys() []K
	Values() []V
	KeyValues() []pair.Pair[K, V]
	Back() (K, V)
	Front() (K, V)
	fmt.Stringer
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
