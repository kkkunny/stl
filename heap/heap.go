package heap

import (
	"fmt"

	"github.com/kkkunny/stl/list"
	. "github.com/kkkunny/stl/types"
)

// 堆
type Heap[T Comparator[T]] interface {
	fmt.Stringer
	Length() Usize
	Empty() bool
	Push(e ...T)
	Pop() T
	Peek() T
	Clear()
	Iterator() *list.ArrayListIterator[T]
}
