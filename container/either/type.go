package either

import stlbasic "github.com/kkkunny/stl/basic"

type Either[L, R any] struct {
	left *bool
	data any
}

func Left[L, R any](v L) Either[L, R] {
	left := true
	return Either[L, R]{
		left: &left,
		data: v,
	}
}

func Right[L, R any](v R) Either[L, R] {
	left := false
	return Either[L, R]{
		left: &left,
		data: v,
	}
}

func (self Either[L, R]) Default() Either[L, R] {
	return Left[L, R](stlbasic.Default[L]())
}
