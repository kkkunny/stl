package tuple

import (
	"fmt"

	stlcmp "github.com/kkkunny/stl/cmp"
	stlval "github.com/kkkunny/stl/value"
)

type Tuple3[T, E, F any] struct {
	t T
	e E
	f F
}

func Pack3[T, E, F any](t T, e E, f F) Tuple3[T, E, F] {
	return Tuple3[T, E, F]{t: t, e: e, f: f}
}

func (self Tuple3[T, E, F]) Unpack() (T, E, F) {
	return self.t, self.e, self.f
}

func (self Tuple3[T, E, F]) Equal(dstObj any) bool {
	dst, ok := dstObj.(Tuple3[T, E, F])
	if !ok {
		return false
	}
	return stlcmp.Equal(self.t, dst.t) && stlcmp.Equal(self.e, dst.e) && stlcmp.Equal(self.f, dst.f)
}

func (self Tuple3[T, E, F]) Clone() any {
	return Pack3(stlval.Clone(self.t), stlval.Clone(self.e), stlval.Clone(self.f))
}

func (self Tuple3[T, E, F]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", self.t, self.e, self.f)
}

func (self Tuple3[T, E, F]) E1() T {
	return self.t
}

func (self Tuple3[T, E, F]) E2() E {
	return self.e
}

func (self Tuple3[T, E, F]) E3() F {
	return self.f
}
