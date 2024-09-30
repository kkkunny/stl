package either

import (
	stlbasic "github.com/kkkunny/stl/cmp"
	"github.com/kkkunny/stl/container/pair"
)

func (self Either[L, R]) Equal(dst Either[L, R]) bool {
	self.init()
	dst.init()

	if self.IsLeft() {
		return stlbasic.Equal(pair.NewPair(self.Left()), pair.NewPair(dst.Left()))
	} else {
		return stlbasic.Equal(pair.NewPair(self.Right()), pair.NewPair(dst.Right()))
	}
}
