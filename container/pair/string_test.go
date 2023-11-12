package pair

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestPair_String(t *testing.T) {
	pair := NewPair(1, 2)
	stltest.AssertEq(t, pair.String(), "(1, 2)")
}
