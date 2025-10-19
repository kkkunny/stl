package lazy

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestGetter(t *testing.T) {
	var i int
	getter := Getter(func() (int, error) {
		i++
		return i, nil
	})
	v, err := getter()
	stltest.AssertEq(t, err, nil)
	stltest.AssertEq(t, v, 1)
	v, err = getter()
	stltest.AssertEq(t, err, nil)
	stltest.AssertEq(t, v, 1)
}

func TestGetterWith(t *testing.T) {
	getter := GetterWith(func(in int) (int, error) {
		return in, nil
	})
	v, err := getter(1)
	stltest.AssertEq(t, err, nil)
	stltest.AssertEq(t, v, 1)
	v, err = getter(2)
	stltest.AssertEq(t, err, nil)
	stltest.AssertEq(t, v, 1)
}
