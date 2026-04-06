package either

func (e Either[L, R]) TryLeft() (L, bool) {
	return e.l, !e.right
}

func (e Either[L, R]) Left() L {
	return e.l
}

func (e Either[L, R]) TryRight() (R, bool) {
	return e.r, e.right
}

func (e Either[L, R]) Right() R {
	return e.r
}

func (e Either[L, R]) IsLeft() bool {
	return !e.right
}

func (e Either[L, R]) IsRight() bool {
	return e.right
}

func (e *Either[L, R]) SetLeft(v L) L {
	bak := e.l
	e.l = v
	return bak
}

func (e *Either[L, R]) SetRight(v R) R {
	bak := e.r
	e.r = v
	return bak
}
