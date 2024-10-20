//go:build goexperiment.rangefunc || go1.23

package hashmap

import "iter"

func (self *_SwissTable[K, V]) Iter2() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		self.data.Iter(func(k K, v V) (stop bool) {
			return !yield(k, v)
		})
	}
}
