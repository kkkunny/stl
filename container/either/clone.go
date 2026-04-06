package either

import "github.com/kkkunny/stl/clone"

func (e Either[L, R]) Clone() Either[L, R] {
	return Either[L, R]{
		right: e.right,
		l:     clone.Clone(e.l),
		r:     clone.Clone(e.r),
	}
}
