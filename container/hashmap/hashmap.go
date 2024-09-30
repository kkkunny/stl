package hashmap

import (
	"encoding/json"
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
)

type HashMap[K, V any] interface {
	stlbasic.Defaultable[HashMap[K, V]]
	stlbasic.Capacityable
	stlbasic.Cloneable[HashMap[K, V]]
	stlcmp.Comparable[HashMap[K, V]]
	stliter.IteratorContainer[pair.Pair[K, V]]
	json.Marshaler
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
	fmt.Stringer
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
