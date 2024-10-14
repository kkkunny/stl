package stlunsafe

import (
	"reflect"
	"unsafe"
)

// FlatPointer 获取指向底层值的指针
func FlatPointer(v any) unsafe.Pointer {
	vv := reflect.ValueOf(v)
	switch vt := vv.Type(); vt.Kind() {
	case reflect.Pointer:
		if vt.Elem().Kind() != reflect.Pointer {
			return reflect.ValueOf(vv).FieldByName("ptr").UnsafePointer()
		}
		return FlatPointer(vv.Elem().Interface())
	default:
		return reflect.ValueOf(vv).FieldByName("ptr").UnsafePointer()
	}
}

// Pointer 获取指向值的指针
func Pointer(v any) unsafe.Pointer {
	vv := reflect.ValueOf(v)
	switch vt := vv.Type(); vt.Kind() {
	case reflect.Pointer:
		vp := reflect.New(vt)
		vp.Elem().Set(vv)
		return vp.UnsafePointer()
	default:
		return reflect.ValueOf(vv).FieldByName("ptr").UnsafePointer()
	}
}
