//go:build goexperiment.rangefunc || go1.23

package treemap

import (
	"iter"
)

func (self *_StdTreeMap[K, V]) Iter2() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := self.data.Iterator(); i != nil; i = i.Next() {
			if !yield(i.Key, i.Value) {
				return
			}
		}
	}
}
