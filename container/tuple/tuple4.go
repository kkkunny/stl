package tuple

import (
	"fmt"

	"github.com/kkkunny/stl/clone"
	stlcmp "github.com/kkkunny/stl/cmp"
)

type Tuple4[E1, E2, E3, E4 any] struct {
	e1 E1
	e2 E2
	e3 E3
	e4 E4
}

func Pack4[E1, E2, E3, E4 any](e1 E1, e2 E2, e3 E3, e4 E4) Tuple4[E1, E2, E3, E4] {
	return Tuple4[E1, E2, E3, E4]{e1: e1, e2: e2, e3: e3, e4: e4}
}

func (self Tuple4[E1, E2, E3, E4]) Unpack() (E1, E2, E3, E4) {
	return self.e1, self.e2, self.e3, self.e4
}

func (self Tuple4[E1, E2, E3, E4]) Equal(dstObj any) bool {
	dst, ok := dstObj.(Tuple4[E1, E2, E3, E4])
	if !ok {
		return false
	}
	return stlcmp.Equal(self.e1, dst.e1) && stlcmp.Equal(self.e2, dst.e2) && stlcmp.Equal(self.e3, dst.e3) && stlcmp.Equal(self.e4, dst.e4)
}

func (self Tuple4[E1, E2, E3, E4]) Clone() any {
	return Pack4(clone.Clone(self.e1), clone.Clone(self.e2), clone.Clone(self.e3), clone.Clone(self.e4))
}

func (self Tuple4[E1, E2, E3, E4]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", self.e1, self.e2, self.e3, self.e4)
}

func (self Tuple4[E1, E2, E3, E4]) E1() E1 {
	return self.e1
}

func (self Tuple4[E1, E2, E3, E4]) E2() E2 {
	return self.e2
}

func (self Tuple4[E1, E2, E3, E4]) E3() E3 {
	return self.e3
}

func (self Tuple4[E1, E2, E3, E4]) E4() E4 {
	return self.e4
}
