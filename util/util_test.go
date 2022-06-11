package util

import (
	"fmt"
	"testing"
)

func test(v int) int {
	fmt.Println(v)
	return v
}

func TestTernary(t *testing.T) {
	cond := true
	fmt.Println(Ternary[int](cond, func() int {
		return test(1)
	}, func() int {
		return test(0)
	}))
	t.Fail()
}
