package either

import (
	stlbasic "github.com/kkkunny/stl/cmp"
	"github.com/kkkunny/stl/container/tuple"
)

func (self Either[L, R]) Equal(dst Either[L, R]) bool {
	self.init()
	dst.init()

	if self.IsLeft() {
		return stlbasic.Equal(tuple.Pack2(self.Left()), tuple.Pack2(dst.Left()))
	} else {
		return stlbasic.Equal(tuple.Pack2(self.Right()), tuple.Pack2(dst.Right()))
	}
}
