package stlcmp

import (
	"fmt"
	"reflect"

	stlreflect "github.com/kkkunny/stl/reflect"
)

// Equalable 可相等的
type Equalable[Self any] interface {
	Equal(dst Self) bool
}

// GetEqualFunc 获取比较函数，若没有会panic
func GetEqualFunc[T any]() func(l, r T) bool {
	t := stlreflect.Type[T]()
	switch {
	case t.Implements(stlreflect.Type[Equalable[T]]()):
		return func(l, r T) bool {
			return any(l).(Equalable[T]).Equal(r)
		}
	case t.Implements(stlreflect.Type[Equalable[any]]()):
		return func(l, r T) bool {
			return any(l).(Equalable[any]).Equal(r)
		}
	default:
		f := getReflectEqualFunc(t)
		return func(l, r T) bool {
			return f(l, r)
		}
	}
}

func getReflectEqualFunc(vt reflect.Type) func(l, r any) bool {
	switch {
	case vt.Comparable():
		return func(l, r any) bool {
			lvobj, rvobj := reflect.ValueOf(l), reflect.ValueOf(r)
			return lvobj.Equal(rvobj)
		}
	case vt.Kind() == reflect.Slice:
		elemFn := getReflectEqualFunc(vt.Elem())
		return func(l, r any) bool {
			lvobj, rvobj := reflect.ValueOf(l), reflect.ValueOf(r)
			if lvobj.Len() != rvobj.Len() {
				return false
			}
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
	return GetEqualFunc[T]()(lv, rv)
}
