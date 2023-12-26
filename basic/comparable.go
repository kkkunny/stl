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
	switch lvv := any(lv).(type) {
	case Comparable[T]:
		return lvv.Equal(rv)
	default:
		return reflectEqual(lv, rv)
	}
}

func reflectEqual[T any](lv, rv T) bool {
	vt := reflect.TypeOf(lv)
	if vt.Comparable() {
		return reflect.ValueOf(lv).Equal(reflect.ValueOf(rv))
	} else {
		panic(fmt.Errorf("type `%s` cannot be compared", vt))
	}
}
