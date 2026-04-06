package either

import (
	stlbasic "github.com/kkkunny/stl/cmp"
)

func (e Either[L, R]) Equal(dst Either[L, R]) bool {
	if e.right != dst.right {
		return false
	}
	if !e.right {
		return stlbasic.Equal(e.l, dst.l)
	} else {
		return stlbasic.Equal(e.r, dst.r)
	}
}
