package linkedhashset

import (
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/linkedhashmap"
)

type LinkedHashSet[T any] interface {
	stlbasic.Defaultable[LinkedHashSet[T]]
	stlbasic.Cloneable[LinkedHashSet[T]]
	stlcmp.Comparable[LinkedHashSet[T]]
	stliter.IteratorContainer[T]
	stlbasic.Lengthable
	Add(v T) bool
	Contain(v T) bool
	Remove(v T) bool
	Clear()
	Empty() bool
	ToSlice() []T
	fmt.Stringer
	getData() linkedhashmap.LinkedHashMap[T, struct{}]
}

// StdWith 使用go的默认hash函数，相比map，write更慢，read更快
func StdWith[T comparable](vs ...T) LinkedHashSet[T] {
	if len(vs) == 0 {
		return _NewStdLinkedHashSet[T]()
	} else {
		return _NewStdLinkedHashSetWith[T](vs...)
	}
}

// StdWithCap 使用go的默认hash函数，相比map，write更慢，read更快
func StdWithCap[T comparable](cap uint) LinkedHashSet[T] {
	return _NewStdLinkedHashSetWithCapacity[T](cap)
}

// AnyWith 使用自定义hash函数，相比map，全方位慢
func AnyWith[T any](vs ...T) LinkedHashSet[T] {
	if len(vs) == 0 {
		return _NewAnyLinkedHashSet[T]()
	} else {
		return _NewAnyLinkedHashSetWith[T](vs...)
	}
}

// AnyWithCap 使用自定义hash函数，相比map，全方位慢
func AnyWithCap[T any](cap uint) LinkedHashSet[T] {
	return _NewAnyLinkedHashSetWithCapacity[T](cap)
}
