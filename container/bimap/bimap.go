package bimap

import (
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
	stlcmp "github.com/kkkunny/stl/cmp"
	"github.com/kkkunny/stl/container/hashmap"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
)

type BiMap[T, E any] interface {
	stlbasic.Defaultable[BiMap[T, E]]
	stlbasic.Capacityable
	stlbasic.Cloneable[BiMap[T, E]]
	stlcmp.Comparable[BiMap[T, E]]
	stliter.IteratorContainer[pair.Pair[T, E]]
	stlbasic.Lengthable
	Set(k T, v E) (T, E)
	GetValue(k T, defaultValue ...E) E
	GetKey(v E, defaultKey ...T) T
	ContainKey(k T) bool
	ContainValue(v E) bool
	RemoveKey(k T, defaultValue ...E) E
	RemoveValue(v E, defaultKey ...T) T
	Clear()
	Empty() bool
	Keys() []T
	Values() []E
	KeyValues() []pair.Pair[T, E]
	fmt.Stringer
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
