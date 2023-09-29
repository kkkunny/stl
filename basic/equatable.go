package stlbasic

import (
	"fmt"
	"reflect"
)

// Comparable 可比较的
type Comparable interface {
	Equal(dst any) bool
}

// Equal 比较是否相等
func Equal[T any](lv, rv T) bool {
	if _, ok := any(lv).(Comparable); ok {
		if eqlv, ok := any(lv).(Comparable); ok {
			if !eqlv.Equal(rv) {
				return false
			}
		}
	} else if tp := reflect.TypeOf(lv); tp.Comparable() {
		if !reflect.ValueOf(lv).Equal(reflect.ValueOf(rv)) {
			return false
		}
	} else {
		panic(fmt.Errorf("type `%s` cannot be compared", tp))
	}
	return true
}
