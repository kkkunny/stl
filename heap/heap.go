package heap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 堆
type Heap[T constraints.Ordered] interface {
	fmt.Stringer
	Length() int
	Empty() bool
	Push(e ...T)
	Pop() T
	Peek() T
	Clear()
}
