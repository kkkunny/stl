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
	lvobj, rvobj := reflect.ValueOf(lv), reflect.ValueOf(rv)
	vt := lvobj.Type()
	switch {
	case vt.Comparable():
		return lvobj.Equal(rvobj)
	case vt.Kind() == reflect.Slice:
		if lvobj.Len() != rvobj.Len() {
			return false
		}
		for i := 0; i < lvobj.Len(); i++ {
			if !reflectEqual(lvobj.Index(i).Interface(), rvobj.Index(i).Interface()) {
				return false
			}
		}
		return true
	default:
		panic(fmt.Errorf("type `%s` cannot be compared", vt))
	}
}
