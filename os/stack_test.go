package stlos

import (
	"testing"

	stltest "github.com/kkkunny/stl/test"
)

func TestGetCallStacks(t *testing.T) {
	stacks := GetCallStacks(20)
	stltest.AssertEq(t, stacks[0].File, "W:/code/go/github.com/kkkunny/stl/os/stack_test.go")
}
