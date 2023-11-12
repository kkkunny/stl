package pair

import (
	"fmt"
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestPair_String(t *testing.T) {
	pair := NewPair(1, 2)
	fmt.Println(pair)
	stltest.AssertEq(t, pair.String(), "(1, 2)")
}
