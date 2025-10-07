package either

import (
	"testing"

	"github.com/kkkunny/stl/container/tuple"
	stltest "github.com/kkkunny/stl/test"
)

func TestEither_IsLeft(t *testing.T) {
	v := Left[int, int](1)
	stltest.AssertEq(t, v.IsLeft(), true)
	stltest.AssertEq(t, v.IsRight(), false)
}

func TestEither_IsRight(t *testing.T) {
	v := Right[int, int](1)
	stltest.AssertEq(t, v.IsLeft(), false)
	stltest.AssertEq(t, v.IsRight(), true)
}

func TestEither_Left(t *testing.T) {
	v := Left[int, int](1)
	stltest.AssertEq(t, tuple.Pack2(v.Left()), tuple.Pack2(1, true))
	v.SetLeft(2)
	stltest.AssertEq(t, tuple.Pack2(v.Left()), tuple.Pack2(2, true))
}

func TestEither_Right(t *testing.T) {
	v := Right[int, int](1)
	stltest.AssertEq(t, tuple.Pack2(v.Right()), tuple.Pack2(1, true))
	v.SetRight(2)
	stltest.AssertEq(t, tuple.Pack2(v.Right()), tuple.Pack2(2, true))
}
