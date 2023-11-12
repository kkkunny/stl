package stlbasic

import (
	"fmt"
	"reflect"
)

// Comparable 可比较的
type Comparable[Self any] interface {
	Equal(dst Self) bool
}

// Equal 比较是否相等
func Equal[T any](lv, rv T) bool {
	if eqlv, ok := any(lv).(Comparable[T]); ok {
		if !eqlv.Equal(rv) {
			return false
		}
	} else {
		vtype := reflect.TypeOf(lv)
		if vtype.Comparable() {
			return reflect.ValueOf(lv).Equal(reflect.ValueOf(rv))
		} else {
			panic(fmt.Errorf("type `%s` cannot be compared", vtype))
		}
	}
	return true
}
