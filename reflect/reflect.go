package stlreflect

import (
	"reflect"

	reflect2 "github.com/kkkunny/stl/internal/reflect"
)

func Zero[T any]() reflect.Value {
	return reflect.Zero(reflect.TypeFor[T]())
}

func GetStructOrStructPtrFieldValue(v reflect.Value, fieldName string) reflect.Value {
	return reflect2.GetStructOrStructPtrFieldValue(v, fieldName)
}
