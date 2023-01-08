package types

type Either[L, R any] struct {
	left  *L
	right *R
}

func Left[L, R any](v L) Either[L, R] {
	return Either[L, R]{left: &v}
}

func Right[L, R any](v R) Either[L, R] {
	return Either[L, R]{right: &v}
}

// Left dangerous
func (self Either[L, R]) Left() L {
	return *self.left
}

// Right dangerous
func (self Either[L, R]) Right() R {
	return *self.right
}

func (self Either[L, R]) IsLeft() bool {
	return self.left != nil
}

func (self Either[L, R]) IsRight() bool {
	return self.right != nil
}
