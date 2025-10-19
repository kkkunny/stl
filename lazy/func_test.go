package lazy

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestLatestFunc(t *testing.T) {
	var j int
	fn := LatestFunc(func(i int) (int, error) {
		j++
		return i + j, nil
	})
	v, err := fn(1)
	stltest.AssertEq(t, err, nil)
	stltest.AssertEq(t, v, 2)
	v, err = fn(1)
	stltest.AssertEq(t, err, nil)
	stltest.AssertEq(t, v, 2)
	v, err = fn(2)
	stltest.AssertEq(t, err, nil)
	stltest.AssertEq(t, v, 4)
}
