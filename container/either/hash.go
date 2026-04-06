package either

import (
	stlbasic "github.com/kkkunny/stl/hash"
)

func (e Either[L, R]) Hash() uint64 {
	if !e.right {
		return stlbasic.Hash(e.l)
	} else {
		return stlbasic.Hash(e.r)
	}
}
