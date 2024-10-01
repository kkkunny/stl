//go:build goexperiment.rangefunc || go1.23

package linkedlist

import (
	"iter"
)

func (self LinkedList[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for cursor := self.root; cursor != nil; cursor = cursor.Next {
			if !yield(cursor.Value) {
				return
			}
		}
	}
}
