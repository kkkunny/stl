package util

/*
#include <stdio.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

// Ternary 三元运算
func Ternary[T any](c bool, t, f T) T {
	if c {
		return t
	} else {
		return f
	}
}

// Alloc 分配内存
func Alloc(size uintptr) (unsafe.Pointer, error) {
	ptr := C.calloc(1, C.size_t(size))
	if ptr == nil {
		return nil, errors.New("heap memory alloc error")
	}
	return ptr, nil
}

// Free 释放内存
func Free(ptr unsafe.Pointer) {
	C.free(ptr)
}
