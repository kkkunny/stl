package either

import (
	stlbasic "github.com/kkkunny/stl/value"
)

func (self Either[L, R]) Clone() Either[L, R] {
	self.init()
	return Either[L, R]{
		left: self.left,
		data: stlbasic.Clone(self.data),
	}
}
