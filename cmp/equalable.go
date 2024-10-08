package stlcmp

import (
	"fmt"
	"reflect"
)

// Equalable 可相等的
type Equalable[Self any] interface {
	Equal(dst Self) bool
}

// GetEqualFunc 获取比较函数，若没有会panic
func GetEqualFunc[T any](vs ...T) func(l, r T) bool {
	var v T
	if len(vs) > 0 {
		v = vs[0]
	}

	switch any(v).(type) {
	case Equalable[T]:
		return func(l, r T) bool {
			return any(l).(Equalable[T]).Equal(r)
		}
	case Equalable[any]:
		return func(l, r T) bool {
			return any(l).(Equalable[any]).Equal(r)
		}
	default:
		f := getReflectEqualFunc(reflect.ValueOf(v))
		return func(l, r T) bool {
			return f(l, r)
		}
	}
}

func getReflectEqualFunc(v reflect.Value) func(l, r any) bool {
	if !v.IsValid() {
		return func(_, _ any) bool {
			return true
		}
	}

	vt := v.Type()
	switch {
	case vt.Comparable():
		return func(l, r any) bool {
			lvobj, rvobj := reflect.ValueOf(l), reflect.ValueOf(r)
			return lvobj.Equal(rvobj)
		}
	case vt.Kind() == reflect.Slice:
		return func(l, r any) bool {
			lvobj, rvobj := reflect.ValueOf(l), reflect.ValueOf(r)
			if lvobj.Len() != rvobj.Len() {
				return false
			}
			elemFn := getReflectEqualFunc(reflect.New(vt.Elem()).Elem())
			for i := 0; i < lvobj.Len(); i++ {
				if !elemFn(lvobj.Index(i).Interface(), rvobj.Index(i).Interface()) {
					return false
				}
			}
			return true
		}
	default:
		panic(fmt.Errorf("type `%s` cannot be compared", vt))
	}
}

// Equal 比较是否相等
func Equal[T any](lv, rv T) bool {
	return GetEqualFunc[T](lv, rv)(lv, rv)
}
