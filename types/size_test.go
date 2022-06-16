package types

import (
	"fmt"
	"testing"
)

func TestSize(t *testing.T) {
	s := GigaByte * 2
	fmt.Println(s)
	t.Fail()
}
