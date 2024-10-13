package stlreflect

import (
	"reflect"

	stlval "github.com/kkkunny/stl/value"
)

func Type[T any]() reflect.Type {
	var v T
	return reflect.TypeOf(&v).Elem()
}

func Zero[T any]() reflect.Value {
	return reflect.ValueOf(stlval.New[T]()).Elem()
}
