package treeset

import (
	"cmp"
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/treemap"
)

type TreeSet[T any] interface {
	stlbasic.Defaultable[TreeSet[T]]
	stlbasic.Cloneable[TreeSet[T]]
	stlcmp.Equalable[TreeSet[T]]
	stliter.IteratorContainer[T]
	stlbasic.Lengthable
	Add(v T) bool
	Contain(v T) bool
	Remove(v T) bool
	Clear()
	Empty() bool
	ToSlice() []T
	fmt.Stringer
	getData() treemap.TreeMap[T, struct{}]
}

// StdWith 使用go的默认cmp函数
func StdWith[T cmp.Ordered](vs ...T) TreeSet[T] {
	if len(vs) == 0 {
		return _NewStdTreeSet[T]()
	} else {
		return _NewStdTreeSetWith[T](vs...)
	}
}

// AnyWith 使用自定义cmp函数，全方位慢
func AnyWith[T any](vs ...T) TreeSet[T] {
	if len(vs) == 0 {
		return _NewAnyTreeSet[T]()
	} else {
		return _NewAnyTreeSetWith[T](vs...)
	}
}
