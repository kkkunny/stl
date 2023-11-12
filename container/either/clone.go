package either

import stlbasic "github.com/kkkunny/stl/basic"

func (self Either[L, R]) Clone() Either[L, R] {
	self.init()
	return Either[L, R]{
		left: self.left,
		data: stlbasic.Clone(self.data),
	}
}
