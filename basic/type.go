package stlbasic

import "unsafe"

// Is 类型是否是
func Is[T any](v any) bool {
	_, ok := v.(T)
	return ok
}

// As 强制转换
func As[From, To any](v From) To {
	return *(*To)(unsafe.Pointer(&v))
}
