package either

import (
	stlbasic "github.com/kkkunny/stl/hash"
)

func (self Either[L, R]) Hash() uint64 {
	self.init()

	if self.IsLeft() {
		return stlbasic.Hash(self.data.(L))
	} else {
		return stlbasic.Hash(self.data.(R))
	}
}
