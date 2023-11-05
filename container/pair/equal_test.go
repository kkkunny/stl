package pair

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestPair_Equal(t *testing.T) {
	pair := NewPair(1, 1)
	stltest.AssertEq(t, pair, NewPair(1, 1))
	stltest.AssertNotEq(t, pair, NewPair(1, 2))
}
