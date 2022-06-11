package queue

import (
	"github.com/kkkunny/stl/list"
	"testing"
)

func BenchmarkArrayListQueue(b *testing.B) {
	al := list.NewArrayList[int](0, 0)
	for i := 0; i < 10; i++ {
		al.Add(i)
	}
	for i := 0; i < 10; i++ {
		al.Remove(0)
	}
}

func BenchmarkSingleLinkedListListQueue(b *testing.B) {
	sll := list.NewSingleLinkedList[int](0, 0)
	for i := 0; i < 10; i++ {
		sll.Add(i)
	}
	for i := 0; i < 10; i++ {
		sll.Remove(0)
	}
}
