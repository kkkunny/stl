//go:build !goexperiment.rangefunc && !go1.23

package stliter

import stlbasic "github.com/kkkunny/stl/basic"

// Iterator 迭代器
type Iterator[T any] interface {
	stlbasic.Lengthable
	Next() bool
	HasNext() bool
	Value() T
	Reset()
}
