package stlbasic

import (
    "fmt"
    "reflect"
)

// Capacityable 可获取容量的
type Capacityable interface {
    Capacity() uint
}

// Capacity 获取容量
func Capacity[T any](v T) uint {
    if vv, ok := any(v).(Capacityable); ok {
        return vv.Capacity()
    } else {
        vtype := reflect.TypeOf(v)
        switch vtype.Kind() {
        case reflect.Array, reflect.Chan, reflect.Slice:
            return uint(reflect.ValueOf(v).Cap())
        default:
            panic(fmt.Errorf("type `%s` cannot be get capacity", vtype))
        }
    }
}
