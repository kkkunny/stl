package either

import (
	stlbasic "github.com/kkkunny/stl/value"
)

func (self Either[L, R]) Clone() any {
	self.init()
	return Either[L, R]{
		left: self.left,
		data: stlbasic.Clone(self.data),
	}
}
