package stlmath

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNotWithBits(t *testing.T) {
	stltest.AssertEq(t, NotWithBits(1, 1), 0)
	stltest.AssertEq(t, NotWithBits(0, 1), 1)
}
