package hashset

import (
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/cmp"
	"github.com/kkkunny/stl/container/hashmap"
	stliter "github.com/kkkunny/stl/container/iter"
)

type HashSet[T any] interface {
	stlbasic.Cloneable[HashSet[T]]
	stlcmp.Equalable[HashSet[T]]
	stliter.IteratorContainer[T]
	stlbasic.Lengthable
	Add(v T) bool
	Contain(v T) bool
	Remove(v T) bool
	Clear()
	Empty() bool
	ToSlice() []T
	fmt.Stringer
	getData() hashmap.HashMap[T, struct{}]
}

// StdWith 使用go的默认hash函数，相比map，write更慢，read更快
func StdWith[T comparable](vs ...T) HashSet[T] {
	if len(vs) == 0 {
		return _NewStdHashSet[T]()
	} else {
		return _NewStdHashSetWith[T](vs...)
	}
}

// StdWithCap 使用go的默认hash函数，相比map，write更慢，read更快
func StdWithCap[T comparable](cap uint) HashSet[T] {
	return _NewStdHashSetWithCapacity[T](cap)
}

// AnyWith 使用自定义hash函数，相比map，全方位慢
func AnyWith[T any](vs ...T) HashSet[T] {
	if len(vs) == 0 {
		return _NewAnyHashSet[T]()
	} else {
		return _NewAnyHashSetWith[T](vs...)
	}
}

// AnyWithCap 使用自定义hash函数，相比map，全方位慢
func AnyWithCap[T any](cap uint) HashSet[T] {
	return _NewAnyHashSetWithCapacity[T](cap)
}
