package util

import (
	"fmt"
	"testing"
)

func TestMust(t *testing.T) {
	a := []int{1, 2, 3}
	content := Must(Json(&a))
	fmt.Println(string(content))
}

func TestGetGoroutineID(t *testing.T) {
	id := GetGoroutineID()
	fmt.Println(id)
}
