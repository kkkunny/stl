package stlos

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestSize_String(t *testing.T) {
	s := Gib*2 + Kib*30 + Byte*23
	stltest.AssertEq(t, s.String(), "2 Gib 30 Kib 23 Byte")
}
