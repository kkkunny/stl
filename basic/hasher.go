package stlbasic

import (
    "fmt"
    "math"
    "reflect"
    "unsafe"
)

// Hashable 可哈希的
type Hashable interface {
    Hash() uint64
}

// Hash 获取哈希
func Hash[T any](v T) uint64 {
    if vv, ok := any(v).(Hashable); ok {
        return vv.Hash()
    } else {
        vtype := reflect.TypeOf(v)
        switch vtype.Kind() {
        case reflect.Bool:
            if any(v).(bool) {
                return 1
            } else {
                return 0
            }
        case reflect.Int:
            return *(*uint64)(unsafe.Pointer(&v))
        case reflect.Int8:
            return uint64(*(*uint8)(unsafe.Pointer(&v)))
        case reflect.Int16:
            return uint64(*(*uint16)(unsafe.Pointer(&v)))
        case reflect.Int32:
            return uint64(*(*uint32)(unsafe.Pointer(&v)))
        case reflect.Int64:
            return *(*uint64)(unsafe.Pointer(&v))
        case reflect.Uint:
            return uint64(any(v).(uint))
        case reflect.Uint8:
            return uint64(any(v).(uint8))
        case reflect.Uint16:
            return uint64(any(v).(uint16))
        case reflect.Uint32:
            return uint64(any(v).(uint32))
        case reflect.Uint64:
            return any(v).(uint64)
        case reflect.Uintptr:
            return uint64(any(v).(uintptr))
        case reflect.Float32:
            return uint64(math.Float32bits(any(v).(float32)))
        case reflect.Float64:
            return math.Float64bits(any(v).(float64))
        case reflect.String:
            var hash uint64
            for _, b := range any(v).(string) {
                hash = 31*hash + uint64(b)
            }
            return hash
        case reflect.UnsafePointer:
            return uint64(uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&v))))
        case reflect.Func:
            return uint64(*(*uintptr)(unsafe.Pointer(&v)))
        case reflect.Pointer:
            return uint64(*(*uintptr)(unsafe.Pointer(&v)))
        case reflect.Array:
            var hash uint64
            vv := reflect.ValueOf(v)
            for i := 0; i < vtype.Len(); i++ {
                hash = 31*hash + Hash(vv.Index(i).Interface())
            }
            return hash
        case reflect.Struct:
            var hash uint64
            vv := reflect.ValueOf(v)
            for i := 0; i < vv.NumField(); i++ {
                hash = 31*hash + Hash(vv.Field(i).Interface())
            }
            return hash
        default:
            panic(fmt.Errorf("type `%s` cannot be get length", vtype))
        }
    }
}
