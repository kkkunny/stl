package stlbasic

import (
    "fmt"
    "reflect"
)

// Lengthable 可获取长度的
type Lengthable interface {
    Length() uint
}

// Length 获取长度
func Length[T any](v T) uint {
    if vv, ok := any(v).(Lengthable); ok {
        return vv.Length()
    } else {
        vtype := reflect.TypeOf(v)
        switch vtype.Kind() {
        case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
            return uint(reflect.ValueOf(v).Len())
        default:
            panic(fmt.Errorf("type `%s` cannot be get length", vtype))
        }
    }
}
