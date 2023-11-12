package pair

import stlbasic "github.com/kkkunny/stl/basic"

func (self Pair[T, F]) Clone() Pair[T, F] {
	return Pair[T, F]{
		First:  stlbasic.Clone(self.First),
		Second: stlbasic.Clone(self.Second),
	}
}
