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

// Ptr 获取值指针
func Ptr[T any](v T) *T {
	return &v
}

// New 获取默认值指针
func New[T any]() *T {
	return Ptr(stlbasic.Default[T]())
}
