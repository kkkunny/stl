package pair

import stlbasic "github.com/kkkunny/stl/basic"

func (self Pair[T, F]) Hash() uint64 {
	return stlbasic.Hash(self.First) ^ (stlbasic.Hash(self.Second) << 1)
}
