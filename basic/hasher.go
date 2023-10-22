package stlbasic

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

// Hashable 可哈希的
type Hashable interface {
	Hash() uint64
}

// Hash 获取哈希
func Hash[T any](v T) uint64 {
	if vv, ok := any(v).(Hashable); ok {
		return vv.Hash()
	} else {
		vv, vt := reflect.ValueOf(v), reflect.TypeOf(v)
		switch vt.Kind() {
		case reflect.Bool:
			if vv.Bool() {
				return 1
			} else {
				return 0
			}
		case reflect.Int:
			return *(*uint64)(unsafe.Pointer(&v))
		case reflect.Int8:
			return uint64(*(*uint8)(unsafe.Pointer(&v)))
		case reflect.Int16:
			return uint64(*(*uint16)(unsafe.Pointer(&v)))
		case reflect.Int32:
			return uint64(*(*uint32)(unsafe.Pointer(&v)))
		case reflect.Int64:
			return *(*uint64)(unsafe.Pointer(&v))
		case reflect.Uint:
			return vv.Uint()
		case reflect.Uint8:
			return vv.Uint()
		case reflect.Uint16:
			return vv.Uint()
		case reflect.Uint32:
			return vv.Uint()
		case reflect.Uint64:
			return vv.Uint()
		case reflect.Uintptr:
			return vv.Uint()
		case reflect.Float32:
			return uint64(math.Float32bits(float32(vv.Float())))
		case reflect.Float64:
			return math.Float64bits(vv.Float())
		case reflect.String:
			var hash uint64
			for _, b := range vv.String() {
				hash = 31*hash + uint64(b)
			}
			return hash
		case reflect.Chan:
			return uint64(vv.Pointer())
		case reflect.UnsafePointer:
			return uint64(vv.Pointer())
		case reflect.Func:
			return uint64(vv.Pointer())
		case reflect.Pointer:
			return uint64(vv.Pointer())
		case reflect.Array:
			var hash uint64
			vv := reflect.ValueOf(v)
			for i := 0; i < vt.Len(); i++ {
				hash = 31*hash + Hash(vv.Index(i).Interface())
			}
			return hash
		case reflect.Struct:
			var hash uint64
			vv := reflect.ValueOf(v)
			for i := 0; i < vv.NumField(); i++ {
				hash = 31*hash + Hash(vv.Field(i).Interface())
			}
			return hash
		default:
			panic(fmt.Errorf("type `%s` cannot be get length", vt))
		}
	}
}
