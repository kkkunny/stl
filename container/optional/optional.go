package optional

import (
	"reflect"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlval "github.com/kkkunny/stl/value"
)

type Optional[T any] struct {
	data *T
}

func Some[T any](v T) Optional[T] {
	return Optional[T]{data: &v}
}

func None[T any]() Optional[T] {
	return Optional[T]{data: nil}
}

func UnPtr[T any, Ptr ~*T](v Ptr) Optional[T] {
	if v == nil {
		return None[T]()
	}
	return Some[T](*v)
}

func UnEmpty[T any](v T) Optional[T] {
	vv := reflect.ValueOf(v)
	if !vv.IsValid() || vv.IsZero() {
		return None[T]()
	}
	return Some[T](v)
}

func UnPtrEmpty[T any, Ptr ~*T](v Ptr) Optional[T] {
	if v == nil {
		return None[T]()
	}
	return UnEmpty[T](*v)
}

func (op Optional[T]) IsSome() bool {
	return op.data != nil
}

func (op Optional[T]) IsNone() bool {
	return op.data == nil
}

func (op Optional[T]) Value() (T, bool) {
	if op.IsNone() {
		return stlval.Default[T](), false
	}
	return *op.data, true
}

func (op Optional[T]) ValueWith(defaultValue ...T) T {
	if stlslices.Empty(defaultValue) {
		return stlval.IgnoreWith(op.Value())
	} else if op.IsNone() {
		return stlslices.First(defaultValue)
	} else {
		return *op.data
	}
}

func (op Optional[T]) MustValue() T {
	if op.IsNone() {
		panic("optional is none")
	}
	return *op.data
}
