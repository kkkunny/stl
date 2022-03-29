package list

import (
	"testing"

	"github.com/kkkunny/stl/types"
)

func BenchmarkList(b *testing.B) {
	l := NewLinkedList[types.I32]()
	for i := types.I32(0); i < 1000000; i++ {
		l.Add(i)
	}
	for i := 0; i < 100; i++ {
		_ = l.Get(999999)
	}
}
