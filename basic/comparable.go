package stlbasic

import (
	"fmt"
	"reflect"
)

// Comparable 可比较的
type Comparable interface {
	Equal(dst any) bool
}

// Equal 比较是否相等
func Equal[T any](lv, rv T) bool {
	if _, ok := any(lv).(Comparable); ok {
		if eqlv, ok := any(lv).(Comparable); ok {
			if !eqlv.Equal(rv) {
				return false
			}
		}
	} else {
		vtype := reflect.TypeOf(lv)
		switch vtype.Kind() {
		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Uintptr,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128,
			reflect.Func,
			reflect.String,
			reflect.Chan,
			reflect.UnsafePointer,
			reflect.Interface,
			reflect.Pointer:
			return reflect.ValueOf(lv).Equal(reflect.ValueOf(rv))
		case reflect.Array:
			lvv, rvv := reflect.ValueOf(lv), reflect.ValueOf(rv)
			length := lvv.Len()
			for i := 0; i < length; i++ {
				if Equal(lvv.Index(i).Interface(), rvv.Index(i).Interface()) {
					return false
				}
			}
			return true
		case reflect.Slice:
			lvv, rvv := reflect.ValueOf(lv), reflect.ValueOf(rv)
			llen, rlen := lvv.Len(), rvv.Len()
			if llen != rlen {
				return false
			}
			for i := 0; i < llen; i++ {
				if Equal(lvv.Index(i).Interface(), rvv.Index(i).Interface()) {
					return false
				}
			}
			return true
		case reflect.Map:
			lvv, rvv := reflect.ValueOf(lv), reflect.ValueOf(rv)
			llen, rlen := lvv.Len(), rvv.Len()
			if llen != rlen {
				return false
			}
			liter := lvv.MapRange()
			for liter.Next() {
				lk, lv := liter.Key(), liter.Value()
				if Equal(lv.Interface(), rvv.MapIndex(lk).Interface()) {
					return false
				}
			}
			return true
		case reflect.Struct:
			lvv, rvv := reflect.ValueOf(lv), reflect.ValueOf(rv)
			for i := 0; i < lvv.NumField(); i++ {
				if !Equal(lvv.Field(i).Interface(), rvv.Field(i).Interface()) {
					return false
				}
			}
			return true
		default:
			panic(fmt.Errorf("type `%s` cannot be compared", vtype))
		}
	}
	return true
}
