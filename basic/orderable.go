package stlbasic

import (
	"fmt"
	"math/cmplx"
	"reflect"
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
		lvv, rvv := reflect.ValueOf(lv), reflect.ValueOf(rv)
		vtype := reflect.TypeOf(lv)
		switch vtype.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			lv, rv := lvv.Int(), rvv.Int()
			if lv < rv {
				return -1
			} else if lv == rv {
				return 0
			} else {
				return 1
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			lv, rv := lvv.Uint(), rvv.Uint()
			if lv < rv {
				return -1
			} else if lv == rv {
				return 0
			} else {
				return 1
			}
		case reflect.Float32, reflect.Float64:
			lv, rv := lvv.Float(), rvv.Float()
			if lv < rv {
				return -1
			} else if lv == rv {
				return 0
			} else {
				return 1
			}
		case reflect.Complex64, reflect.Complex128:
			lv, rv := cmplx.Abs(lvv.Complex()), cmplx.Abs(rvv.Complex())
			if lv < rv {
				return -1
			} else if lv == rv {
				return 0
			} else {
				return 1
			}
		case reflect.String:
			lv, rv := lvv.String(), rvv.String()
			if lv < rv {
				return -1
			} else if lv == rv {
				return 0
			} else {
				return 1
			}
		case reflect.UnsafePointer:
			lv, rv := uintptr(lvv.UnsafePointer()), uintptr(rvv.UnsafePointer())
			if lv < rv {
				return -1
			} else if lv == rv {
				return 0
			} else {
				return 1
			}
		default:
			panic(fmt.Errorf("type `%s` cannot be ordered", vtype))
		}
	}
}
