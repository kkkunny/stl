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
	Insert(i Usize, e ...T)
	Remove(i Usize) T
	Get(i Usize) T
	First() T
	Last() T
	Set(i Usize, e T) T
	Clear()
}
