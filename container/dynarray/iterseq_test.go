//go:build goexperiment.rangefunc

package dynarray

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestDynArray_IterSeq(t *testing.T) {
	a := NewDynArrayWith(1, 2, 3)
	for i, e := range a.IterSeq() {
		stltest.AssertEq(t, e, i+1)
	}
}
