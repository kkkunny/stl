package stlbits

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestBits(t *testing.T) {
	stltest.AssertEq(t, NotWithLength[int8](100, 8), ^int8(100))
}
