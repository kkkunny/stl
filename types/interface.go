package types

import "golang.org/x/exp/constraints"

// Comparator 比较
type Comparator[T any] interface {
	Compare(T) int
}

// Number 数字
type Number interface {
	constraints.Integer | constraints.Float
}
