package stltype

import (
	"fmt"
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

type _Stringer struct{}

func (_ _Stringer) String() string { return "" }

func TestImplInterface(t *testing.T) {
	stltest.AssertEq(t, ImplInterface[_Stringer, fmt.Stringer](), true)
	stltest.AssertEq(t, ImplInterface[int, fmt.Stringer](), false)
	stltest.AssertEq(t, ImplInterface[error, fmt.Stringer](), false)
	stltest.AssertEq(t, ImplInterface[error, error](), true)
	stltest.AssertEq(t, ImplInterface[int, int](), true)
}