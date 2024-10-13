package stlreflect

import (
	"reflect"
)

func Type[T any]() reflect.Type {
	var v T
	return reflect.TypeOf(&v).Elem()
}

func Zero[T any]() reflect.Value {
	var v T
	return reflect.ValueOf(&v).Elem()
}
