package tuple

import (
	"fmt"

	stlcmp "github.com/kkkunny/stl/cmp"
	stlval "github.com/kkkunny/stl/value"
)

type Tuple2[T, E any] struct {
	t T
	e E
}

func Pack2[T, E any](t T, e E) Tuple2[T, E] {
	return Tuple2[T, E]{t: t, e: e}
}

func (self Tuple2[T, E]) Unpack() (T, E) {
	return self.t, self.e
}

func (self Tuple2[T, E]) Equal(dstObj any) bool {
	dst, ok := dstObj.(Tuple2[T, E])
	if !ok {
		return false
	}
	return stlcmp.Equal(self.t, dst.t) && stlcmp.Equal(self.e, dst.e)
}

func (self Tuple2[T, E]) Clone() any {
	return Pack2(stlval.Clone(self.t), stlval.Clone(self.e))
}

func (self Tuple2[T, E]) String() string {
	return fmt.Sprintf("(%v, %v)", self.t, self.e)
}

func (self Tuple2[T, E]) E1() T {
	return self.t
}

func (self Tuple2[T, E]) E2() E {
	return self.e
}
