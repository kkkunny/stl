package stlbasic

import (
	"fmt"
	"math"
	"reflect"
)

// Hashable 可哈希的
type Hashable interface {
	Hash() uint64
}

// Hash 获取哈希
func Hash[T any](v T) uint64 {
	switch vv := any(v).(type) {
	case Hashable:
		return vv.Hash()
	default:
		return reflectHash(v)
	}
}

func reflectHash[T any](v T) uint64 {
	vv, vt := reflect.ValueOf(v), reflect.TypeOf(v)

	switch vt.Kind() {
	case reflect.Bool:
		if vv.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(vv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return vv.Uint()
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(vv.Float())
	case reflect.String:
		var hash uint64
		for _, b := range vv.String() {
			hash = 31*hash + uint64(b)
		}
		return hash
	case reflect.Chan, reflect.UnsafePointer, reflect.Func, reflect.Pointer, reflect.Map:
		return uint64(vv.Pointer())
	case reflect.Array:
		var hash uint64
		for i := 0; i < vv.Len(); i++ {
			hash = 31*hash + Hash(vv.Index(i).Interface())
		}
		return hash
	case reflect.Slice:
		var hash uint64
		for i := 0; i < vv.Len(); i++ {
			hash = 31*hash + Hash(vv.Index(i).Interface())
		}
		return hash
	case reflect.Struct:
		var hash uint64
		for i := 0; i < vv.NumField(); i++ {
			hash = 31*hash + Hash(vv.Field(i).Interface())
		}
		return hash
	default:
		panic(fmt.Errorf("type `%s` cannot get hash", vt))
	}
}
