//go:build goexperiment.rangefunc || go1.23

package linkedhashmap

import (
	"iter"
)

func (self *_AnyLinkedHashMap[K, V]) Iter2() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
			if !yield(cursor.Value.Unpack()) {
				return
			}
		}
	}
}
