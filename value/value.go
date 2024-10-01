package stlval

import (
	"unsafe"

	stlslices "github.com/kkkunny/stl/container/slices"
)

func Default[T any]() T {
	var v T
	return v
}

// ValueOr 获取值，如果为零值则返回defaultVal
func ValueOr[T comparable](v T, defaultVal ...T) T {
	if v != Default[T]() {
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
	return Ptr(Default[T]())
}

// Is 类型是否是
func Is[T any](v any) bool {
	_, ok := v.(T)
	return ok
}

// As 强制转换
func As[From, To any](v From) To {
	return *(*To)(unsafe.Pointer(&v))
}

// Ternary 三目运算
func Ternary[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}

// TernaryAction 三目运算行为
func TernaryAction[T any](cond bool, t, f func() T) T {
	if cond {
		return t()
	}
	return f()
}
