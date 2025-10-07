package enum

import (
	"reflect"
	"strings"
)

type Option func(i int, field reflect.StructField) (reflect.Value, bool)

func WithLower(_ int, field reflect.StructField) (reflect.Value, bool) {
	if field.Type.Kind() != reflect.String {
		return reflect.Value{}, false
	}
	_, ok := field.Tag.Lookup(Tag)
	if ok {
		return reflect.Value{}, false
	}

	return reflect.ValueOf(strings.ToLower(field.Name)), true
}

func WithBitmask(i int, field reflect.StructField) (reflect.Value, bool) {
	if !field.Type.ConvertibleTo(reflect.TypeFor[int]()) {
		return reflect.Value{}, false
	}
	_, ok := field.Tag.Lookup(Tag)
	if ok {
		return reflect.Value{}, false
	}

	return reflect.ValueOf(1 << i), true
}
