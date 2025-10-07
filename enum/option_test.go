package enum

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestWithLower(t *testing.T) {
	e1 := New[struct {
		A string `enum:"A"`
		B string
	}](WithLower)
	e2 := struct {
		A string `enum:"A"`
		B string
	}{
		A: "A",
		B: "b",
	}
	stltest.AssertEq(t, e1, e2)
}

func TestWithNumber(t *testing.T) {
	e1 := New[struct {
		A int `enum:"0"`
		B int
		C int
	}](WithBitmask)
	e2 := struct {
		A int `enum:"0"`
		B int
		C int
	}{
		A: 0,
		B: 2,
		C: 4,
	}
	stltest.AssertEq(t, e1, e2)
}
