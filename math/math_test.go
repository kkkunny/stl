package stlmath

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNotWithBits(t *testing.T) {
	stltest.AssertEq(t, NotWithBits[uint8](1, 1), 0)
	stltest.AssertEq(t, NotWithBits[uint8](0, 1), 1)
	stltest.AssertEq(t, NotWithBits[int64](100, 64), ^100)
}
