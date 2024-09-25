package stlval

import (
	stlbasic "github.com/kkkunny/stl/basic"
	stlslices "github.com/kkkunny/stl/container/slices"
)

// ValueOr 获取值，如果为零值则返回defaultVal
func ValueOr[T comparable](v T, defaultVal ...T) T {
	if v != stlbasic.Default[T]() {
		return v
	}
	return stlslices.Last(defaultVal)
}

// DerefPtrOr 解引用，如果为nil则返回defaultVal
func DerefPtrOr[T any](p *T, defaultVal ...T) T {
	if p != nil {
		return *p
	}
	return stlslices.Last(defaultVal)
}
