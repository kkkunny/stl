package stlcmp

import (
	"fmt"
	"reflect"

	reflect2 "github.com/kkkunny/stl/internal/reflect"
)

// Equalable 可相等的
type Equalable[Self any] interface {
	Equal(dst Self) bool
}

// GetEqualFunc 获取比较函数，若没有会panic
func GetEqualFunc[T any]() func(l, r T) bool {
	t := reflect2.TypeFor[T]()
	switch {
	case t.Implements(reflect2.TypeFor[Equalable[T]]()):
		return func(l, r T) bool {
			return any(l).(Equalable[T]).Equal(r)
		}
	case t.Implements(reflect2.TypeFor[Equalable[any]]()):
		return func(l, r T) bool {
			return any(l).(Equalable[any]).Equal(r)
		}
	default:
		f, ok := getEqualFunc(t, true)
		if !ok {
			panic(fmt.Errorf("type `%s` cannot be compared", t))
		}
		return func(l, r T) bool {
			return f(l, r)
		}
	}
}

func getEqualFunc(t reflect.Type, useRuntime bool) (func(l, r any) bool, bool) {
	it := reflect2.TypeFor[Equalable[any]]()
	switch {
	case t.Implements(it):
		return func(l, r any) bool {
			return l.(Equalable[any]).Equal(r)
		}, true
	default:
		methods := make([]reflect.Method, it.NumMethod())
		for i := 0; i < len(methods); i++ {
			methods[i] = reflect2.ReplaceMethodAnyTo(it.Method(i), t)
		}
		if !reflect2.HasAllMethod(t, methods...) {
			if !useRuntime {
				return nil, false
			}
			f, ok := getRuntimeEqualFunc(t)
			if !ok {
				return nil, false
			}
			return f, true
		}
		return func(l, r any) bool {
			lv, rv := reflect.ValueOf(l), reflect.ValueOf(r)
			method := lv.MethodByName("Equal")
			return method.Call([]reflect.Value{rv})[0].Bool()
		}, true
	}
}

func getRuntimeEqualFunc(t reflect.Type) (func(l, r any) bool, bool) {
	switch {
	case t.Kind() == reflect.Array, t.Kind() == reflect.Slice:
		elemFn, ok := getEqualFunc(t.Elem(), true)
		if !ok {
			return nil, false
		}
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
		}, true
	case t.Kind() == reflect.Map:
		valFn, ok := getEqualFunc(t.Elem(), true)
		if !ok {
			return nil, false
		}
		return func(l, r any) bool {
			lvobj, rvobj := reflect.ValueOf(l), reflect.ValueOf(r)
			if lvobj.Len() != rvobj.Len() {
				return false
			}
			for iter := lvobj.MapRange(); iter.Next(); {
				lv, rv := iter.Value(), rvobj.MapIndex(iter.Key())
				if !rv.IsValid() || !valFn(lv.Interface(), rv.Interface()) {
					return false
				}
			}
			return true
		}, true
	case t.Kind() == reflect.Struct:
		f, ok := getEqualFunc(t, false)
		if ok {
			return f, true
		}

		length := t.NumField()
		fieldFns := make([]func(any, any) bool, length)
		for i := 0; i < length; i++ {
			fieldFns[i], ok = getEqualFunc(t.Field(i).Type, true)
			if !ok {
				return nil, false
			}
		}
		return func(l, r any) bool {
			lvobj, rvobj := reflect.ValueOf(l), reflect.ValueOf(r)
			for i := 0; i < length; i++ {
				lf, rf := lvobj.Field(i), rvobj.Field(i)
				if !fieldFns[i](lf.Interface(), rf.Interface()) {
					return false
				}
			}
			return true
		}, true
	case t.Comparable():
		return func(l, r any) bool {
			lvobj, rvobj := reflect.ValueOf(l), reflect.ValueOf(r)
			return lvobj.Equal(rvobj)
		}, true
	default:
		return nil, false
	}
}

// Equal 比较是否相等
func Equal[T any](lv, rv T) bool {
	return GetEqualFunc[T]()(lv, rv)
}
