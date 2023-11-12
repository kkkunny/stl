package pair

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestPair_Clone(t *testing.T) {
	pair1 := NewPair(1, 2)
	pair2 := pair1.Clone()
	stltest.AssertEq(t, pair1, pair2)
}
