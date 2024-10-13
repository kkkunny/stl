package stlval

import (
	"fmt"
	"reflect"

	stlreflect "github.com/kkkunny/stl/reflect"
)

// Cloneable 可克隆的
type Cloneable[Self any] interface {
	Clone() Self
}

// GetCloneFunc 获取克隆函数，若没有会panic
func GetCloneFunc[T any]() func(v T) T {
	t := stlreflect.Type[T]()
	switch {
	case t.Implements(stlreflect.Type[Cloneable[T]]()):
		return func(vv T) T {
			return any(vv).(Cloneable[T]).Clone()
		}
	case t.Implements(stlreflect.Type[Cloneable[any]]()):
		return func(vv T) T {
			return any(vv).(Cloneable[any]).Clone().(T)
		}
	default:
		f := getReflectCloneFunc(t)
		return func(v T) T {
			return f(v).(T)
		}
	}
}

func getReflectCloneFunc(vt reflect.Type) func(v any) any {
	switch vt.Kind() {
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
		}
	case reflect.Array:
		length := vt.Len()
		elemFn := getReflectCloneFunc(vt.Elem())
		return func(v any) any {
			vv := reflect.ValueOf(v)
			newPtr := reflect.New(vt)
			for i := 0; i < length; i++ {
				elem := reflect.ValueOf(elemFn(vv.Index(i).Interface()))
				newPtr.Elem().Index(i).Set(elem)
			}
			return newPtr.Elem().Interface()
		}
	case reflect.Slice:
		et := vt.Elem()
		length := vt.Len()
		elemFn := getReflectCloneFunc(et)
		return func(v any) any {
			vv := reflect.ValueOf(v)
			newVal := reflect.MakeSlice(et, length, vv.Cap())
			for i := 0; i < length; i++ {
				elem := reflect.ValueOf(elemFn(vv.Index(i).Interface()))
				newVal.Index(i).Set(elem)
			}
			return newVal.Interface()
		}
	case reflect.Map:
		keyT, valT := vt.Key(), vt.Elem()
		keyFn, valFn := getReflectCloneFunc(keyT), getReflectCloneFunc(valT)
		return func(v any) any {
			vv := reflect.ValueOf(v)
			newVal := reflect.MakeMap(vt)
			for iter := vv.MapRange(); iter.Next(); {
				key, val := reflect.ValueOf(keyFn(iter.Key().Interface())), reflect.ValueOf(valFn(iter.Value().Interface()))
				newVal.SetMapIndex(key, val)
			}
			return newVal.Interface()
		}
	case reflect.Struct:
		length := vt.NumField()
		fieldFns := make([]func(any) any, length)
		for i := 0; i < length; i++ {
			fieldFns[i] = getReflectCloneFunc(vt.Field(i).Type)
		}
		return func(v any) any {
			vv := reflect.ValueOf(vt)
			newVal := reflect.New(vt)
			for i := 0; i < length; i++ {
				field := reflect.ValueOf(fieldFns[i](vv.Field(i).Interface()))
				newVal.Elem().Field(i).Set(field)
			}
			return newVal.Elem().Interface()
		}
	default:
		panic(fmt.Errorf("type `%s` cannot be cloned", vt))
	}
}

// Clone 克隆
func Clone[T any](v T) T {
	return GetCloneFunc[T]()(v)
}
