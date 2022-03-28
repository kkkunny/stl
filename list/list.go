package list

import (
	"fmt"

	. "github.com/kkkunny/stl/types"
)

// 列表
type List[T any] interface {
	fmt.Stringer
	Length() Usize
	Empty() bool
	Add(e ...T)
	Insert(i int, e ...T)
	Remove(i int) T
	Get(i int) T
	First() T
	Last() T
	Set(i int, e T) T
	Clear()
}
