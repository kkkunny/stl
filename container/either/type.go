package either

type Either[L, R any] struct {
	right bool
	l     L
	r     R
}

func Left[L, R any](v L) Either[L, R] {
	return Either[L, R]{
		right: false,
		l:     v,
	}
}

func Right[L, R any](v R) Either[L, R] {
	return Either[L, R]{
		right: true,
		r:     v,
	}
}
