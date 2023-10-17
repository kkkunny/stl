package stlbasic

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Orderable 可排序的
type Orderable[Self any] interface {
	Comparable[Self]
	Order(dst Self) int
}

// Order 排序
func Order[T any](lv, rv T) int {
	if eqlv, ok := any(lv).(Orderable[T]); ok {
		return eqlv.Order(rv)
	} else {
		vtype := reflect.TypeOf(lv)
		switch vtype.Kind() {
		case reflect.Int:
			lvv, rvv := any(lv).(int), any(rv).(int)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Int8:
			lvv, rvv := any(lv).(int8), any(rv).(int8)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Int16:
			lvv, rvv := any(lv).(int16), any(rv).(int16)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Int32:
			lvv, rvv := any(lv).(int32), any(rv).(int32)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Int64:
			lvv, rvv := any(lv).(int64), any(rv).(int64)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Uint:
			lvv, rvv := any(lv).(uint), any(rv).(uint)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Uint8:
			lvv, rvv := any(lv).(uint8), any(rv).(uint8)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Uint16:
			lvv, rvv := any(lv).(uint16), any(rv).(uint16)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Uint32:
			lvv, rvv := any(lv).(uint32), any(rv).(uint32)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Uint64:
			lvv, rvv := any(lv).(uint64), any(rv).(uint64)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Uintptr:
			lvv, rvv := any(lv).(uintptr), any(rv).(uintptr)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Float32:
			lvv, rvv := any(lv).(float32), any(rv).(float32)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.Float64:
			lvv, rvv := any(lv).(float64), any(rv).(float64)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.String:
			lvv, rvv := any(lv).(string), any(rv).(string)
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		case reflect.UnsafePointer:
			lvv, rvv := uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&lv))), uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&rv)))
			if lvv < rvv{
				return -1
			}else if lvv == rvv{
				return 0
			}else{
				return 1
			}
		default:
			panic(fmt.Errorf("type `%s` cannot be ordered", vtype))
		}
	}
}
