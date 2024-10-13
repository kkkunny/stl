package pair

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
	stlval "github.com/kkkunny/stl/value"
)

func TestPair_Clone(t *testing.T) {
	pair1 := NewPair(1, 2)
	pair2 := stlval.Clone(pair1)
	stltest.AssertEq(t, pair1, pair2)
}
