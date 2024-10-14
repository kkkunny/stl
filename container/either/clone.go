package either

import "github.com/kkkunny/stl/clone"

func (self Either[L, R]) Clone() Either[L, R] {
	self.init()
	return Either[L, R]{
		left: self.left,
		data: clone.Clone(self.data),
	}
}
