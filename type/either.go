package stltype

import stlbasic "github.com/kkkunny/stl/basic"

type Either[L, R any] struct {
	data any
}

func Left[L, R any](v L) Either[L, R] {
	return Either[L, R]{data: v}
}

func Right[L, R any](v R) Either[L, R] {
	return Either[L, R]{data: v}
}

func (self Either[L, R]) Left() (L, bool) {
	v, ok := self.data.(L)
	return v, ok
}

func (self Either[L, R]) Right() (R, bool) {
	v, ok := self.data.(R)
	return v, ok
}

func (self Either[L, R]) IsLeft() bool {
	_, ok := self.data.(L)
	return ok
}

func (self Either[L, R]) IsRight() bool {
	_, ok := self.data.(R)
	return ok
}

func (self Either[L, R]) Clone() any {
	return Either[L, R]{data: stlbasic.Clone(self.data)}
}
