package stlbasic

import (
	"fmt"
	"reflect"
)

// Clonable 可克隆的
type Clonable interface {
	Clone() any
}

// Clone 克隆
func Clone[T any](v T) T {
	if vv, ok := any(v).(Clonable); ok {
		return vv.Clone().(T)
	} else {
		vtype := reflect.TypeOf(v)
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
			return v
		case reflect.Array:
			vv := reflect.ValueOf(v)
			length := vv.Len()
			nvp := reflect.New(vtype)
			for i := 0; i < length; i++ {
				ne := Clone(vv.Index(i).Interface())
				nvp.Elem().Index(i).Set(reflect.ValueOf(ne))
			}
			return nvp.Elem().Interface().(T)
		case reflect.Slice:
			vv := reflect.ValueOf(v)
			et := vtype.Elem()
			length := vv.Len()
			nv := reflect.MakeSlice(et, length, vv.Cap())
			for i := 0; i < length; i++ {
				ne := Clone(vv.Index(i).Interface())
				nv.Index(i).Set(reflect.ValueOf(ne))
			}
			return nv.Interface().(T)
		case reflect.Map:
			vv := reflect.ValueOf(v)
			nv := reflect.MakeMap(vtype)
			iter := vv.MapRange()
			for iter.Next() {
				key, val := iter.Key(), iter.Value()
				ne := Clone(val.Interface())
				nv.SetMapIndex(key, reflect.ValueOf(ne))
			}
			return nv.Interface().(T)
		case reflect.Struct:
			vv := reflect.ValueOf(v)
			length := vv.NumField()
			nv := reflect.New(vtype)
			for i := 0; i < length; i++ {
				nf := Clone(vv.Field(i).Interface())
				nv.Elem().Field(i).Set(reflect.ValueOf(nf))
			}
			return nv.Elem().Interface().(T)
		default:
			panic(fmt.Errorf("type `%s` cannot be cloned", vtype))
		}
	}
}
