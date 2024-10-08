package stlmaps

import (
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
	stlval "github.com/kkkunny/stl/value"
)

type MapObj[K, V any] interface {
	mapIter2[K, V]
	stlval.Cloneable[any]
	stlcmp.Equalable[any]
	stliter.IteratorContainer[pair.Pair[K, V]]
	stlbasic.Lengthable
	Set(k K, v V) V
	Get(k K, defaultValue ...V) V
	Contain(k K) bool
	Remove(k K, defaultValue ...V) V
	Clear()
	Empty() bool
	Keys() []K
	Values() []V
	KeyValues() []pair.Pair[K, V]
	fmt.Stringer
}
