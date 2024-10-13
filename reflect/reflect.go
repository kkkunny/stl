package stlreflect

import (
	"reflect"

	reflect2 "github.com/kkkunny/stl/internal/reflect"
)

func Zero[T any]() reflect.Value {
	return reflect.Zero(reflect2.TypeFor[T]())
}
