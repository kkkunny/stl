package enum

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestNew(t *testing.T) {
	e1 := New[struct {
		A int
		B uint
		C string `enum:"C"`
		D string `enum:"A"`
		E float32
	}]()
	e2 := struct {
		A int
		B uint
		C string `enum:"C"`
		D string `enum:"A"`
		E float32
	}{
		A: 0,
		B: 1,
		C: "C",
		D: "A",
		E: 4.0,
	}
	stltest.AssertEq(t, e1, e2)
}
