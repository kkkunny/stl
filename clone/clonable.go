package clone

import (
	"fmt"
	"reflect"

	reflect2 "github.com/kkkunny/stl/internal/reflect"
)

// Cloneable 可克隆的
type Cloneable[Self any] interface {
	Clone() Self
}

// GetCloneFunc 获取克隆函数，若没有会panic
func GetCloneFunc[T any]() (f func(v T) T) {
	t := reflect2.TypeFor[T]()
	switch {
	case t.Implements(reflect2.TypeFor[Cloneable[T]]()):
		return func(v T) T {
			return any(v).(Cloneable[T]).Clone()
		}
	case t.Implements(reflect2.TypeFor[Cloneable[any]]()):
		return func(v T) T {
			return any(v).(Cloneable[any]).Clone().(T)
		}
	default:
		f, ok := getCloneFunc(t, true)
		if !ok {
			panic(fmt.Errorf("type `%s` cannot be cloned", t))
		}
		return func(v T) T {
			return f(v).(T)
		}
	}
}

func getCloneFunc(t reflect.Type, useRuntime bool) (f func(v any) any, ok bool) {
	ret, ok := cloneFuncCache.Get(t.String())
	if ok {
		return ret.f, ret.ok
	}
	defer func() {
		cloneFuncCache.Add(t.String(), retType{f: f, ok: ok})
	}()

	it := reflect2.TypeFor[Cloneable[any]]()
	switch {
	case t.Implements(it):
		return func(v any) any {
			return v.(Cloneable[any]).Clone()
		}, true
	default:
		method, ok := t.MethodByName("Clone")
		if ok && method.Type.NumOut() == 1 {
			methods := make([]reflect.Method, it.NumMethod())
			for i := 0; i < len(methods); i++ {
				methods[i] = it.Method(i)
			}
			if reflect2.HasAllMethod(t, method.Type.Out(0), methods...) {
				return func(v any) any {
					vv := reflect.ValueOf(v)
					method := vv.MethodByName("Clone")
					return method.Call(nil)[0].Interface()
				}, true
			}
		}
		if !useRuntime {
			return nil, false
		}
		return getRuntimeCloneFunc(t)
	}
}

func getRuntimeCloneFunc(t reflect.Type) (func(v any) any, bool) {
	switch t.Kind() {
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
		return func(v any) any {
			return v
		}, true
	case reflect.Array:
		length := t.Len()
		elemFn, ok := getCloneFunc(t.Elem(), true)
		if !ok {
			return nil, false
		}
		return func(v any) any {
			vv := reflect.ValueOf(v)
			newPtr := reflect.New(t)
			for i := 0; i < length; i++ {
				elem := reflect.ValueOf(elemFn(vv.Index(i).Interface()))
				newPtr.Elem().Index(i).Set(elem)
			}
			return newPtr.Elem().Interface()
		}, true
	case reflect.Slice:
		et := t.Elem()
		length := t.Len()
		elemFn, ok := getCloneFunc(et, true)
		if !ok {
			return nil, false
		}
		return func(v any) any {
			vv := reflect.ValueOf(v)
			newVal := reflect.MakeSlice(et, length, vv.Cap())
			for i := 0; i < length; i++ {
				elem := reflect.ValueOf(elemFn(vv.Index(i).Interface()))
				newVal.Index(i).Set(elem)
			}
			return newVal.Interface()
		}, true
	case reflect.Map:
		keyT, valT := t.Key(), t.Elem()
		keyFn, ok := getCloneFunc(keyT, true)
		if !ok {
			return nil, false
		}
		valFn, ok := getCloneFunc(valT, true)
		if !ok {
			return nil, false
		}
		return func(v any) any {
			vv := reflect.ValueOf(v)
			newVal := reflect.MakeMap(t)
			for iter := vv.MapRange(); iter.Next(); {
				key, val := reflect.ValueOf(keyFn(iter.Key().Interface())), reflect.ValueOf(valFn(iter.Value().Interface()))
				newVal.SetMapIndex(key, val)
			}
			return newVal.Interface()
		}, true
	case reflect.Struct:
		f, ok := getCloneFunc(t, false)
		if ok {
			return f, true
		}

		length := t.NumField()
		fieldFns := make([]func(any) any, length)
		for i := 0; i < length; i++ {
			fieldFns[i], ok = getCloneFunc(t.Field(i).Type, true)
			if !ok {
				return nil, false
			}
		}
		return func(v any) any {
			vv := reflect.ValueOf(t)
			newVal := reflect.New(t)
			for i := 0; i < length; i++ {
				field := reflect.ValueOf(fieldFns[i](vv.Field(i).Interface()))
				newVal.Elem().Field(i).Set(field)
			}
			return newVal.Elem().Interface()
		}, true
	default:
		return nil, false
	}
}

// Clone 克隆
func Clone[T any](v T) T {
	return GetCloneFunc[T]()(v)
}
