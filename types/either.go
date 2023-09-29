package types

type Either[L, R any] struct {
	value any
}

func Left[L, R any](v L) Either[L, R] {
	return Either[L, R]{value: v}
}

func Right[L, R any](v R) Either[L, R] {
	return Either[L, R]{value: v}
}

func (self Either[L, R]) Left() (L, bool) {
	v, ok := self.value.(L)
	return v, ok
}

func (self Either[L, R]) Right() (R, bool) {
	v, ok := self.value.(R)
	return v, ok
}

func (self Either[L, R]) IsLeft() bool {
	_, ok := self.value.(L)
	return ok
}

func (self Either[L, R]) IsRight() bool {
	_, ok := self.value.(R)
	return ok
}
