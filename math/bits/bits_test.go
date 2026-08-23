package stlbits

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestBits(t *testing.T) {
	stltest.AssertEq(t, NotWithLength[int8](100, 8), ^int8(100))
}

func TestBits_SignedInteger_Empty(t *testing.T) {
	stltest.AssertEq(t, Bits{}.SignedInteger(), int64(0))
}

func TestBits_Equal_LengthMismatch(t *testing.T) {
	stltest.AssertEq(t, NewFromInteger[int](1).Equal(NewFromIntegerWithLength(1, 8)), false)
	stltest.AssertEq(t, NewFromInteger[int](1).Equal(NewFromInteger[int](1)), true)
}
