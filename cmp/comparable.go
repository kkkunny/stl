package stlcmp

import (
	"cmp"
	"fmt"
	"reflect"

	reflect2 "github.com/kkkunny/stl/internal/reflect"
)

// Comparable 可比较的
type Comparable[Self any] interface {
	Equalable[Self]
	Compare(dst Self) int
}

func GetCompareFunc[T any]() func(l, r T) int {
	t := reflect2.TypeFor[T]()
	switch {
	case t.Implements(reflect2.TypeFor[Comparable[T]]()):
		return func(l, r T) int {
			return any(l).(Comparable[T]).Compare(r)
		}
	default:
		f := getReflectCompareFunc(t)
		return func(l, r T) int {
			return f(l, r)
		}
	}
}

func getReflectCompareFunc(vt reflect.Type) func(l, r any) int {
	switch vt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(l, r any) int {
			lv, rv := reflect.ValueOf(l).Int(), reflect.ValueOf(r).Int()
			return cmp.Compare(lv, rv)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return func(l, r any) int {
			lv, rv := reflect.ValueOf(l).Uint(), reflect.ValueOf(r).Uint()
			return cmp.Compare(lv, rv)
		}
	case reflect.Float32, reflect.Float64:
		return func(l, r any) int {
			lv, rv := reflect.ValueOf(l).Float(), reflect.ValueOf(r).Float()
			return cmp.Compare(lv, rv)
		}
	case reflect.String:
		return func(l, r any) int {
			lv, rv := reflect.ValueOf(l).String(), reflect.ValueOf(r).String()
			return cmp.Compare(lv, rv)
		}
	case reflect.UnsafePointer:
		return func(l, r any) int {
			lv, rv := uintptr(reflect.ValueOf(l).UnsafePointer()), uintptr(reflect.ValueOf(r).UnsafePointer())
			return cmp.Compare(lv, rv)
		}
	default:
		panic(fmt.Errorf("type `%s` cannot be ordered", vt))
	}
}

// Compare 比较大小
func Compare[T any](lv, rv T) int {
	return GetCompareFunc[T]()(lv, rv)
}
