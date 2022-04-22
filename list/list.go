package list

import (
	"encoding/json"
	"fmt"
)

// 列表
type List[T any] interface {
	fmt.Stringer
	json.Marshaler
	json.Unmarshaler
	Length() int
	Empty() bool
	Add(e ...T)
	Insert(i int, e ...T)
	Remove(i int) T
	Get(i int) T
	PopFront() T
	PopBack() T
	First() T
	Last() T
	Set(i int, e T) T
	Clear()
	Any(f func(i int, v T) bool) bool
	Every(f func(i int, v T) bool) bool
}
