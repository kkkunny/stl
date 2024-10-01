//go:build goexperiment.rangefunc || go1.23

package hashmap

import "iter"

func (self *_GenericHashMap[K, V]) Iter2() iter.Seq2[K, V] {
	h := toGenericHashMapData(self.data)
	return func(yield func(K, V) bool) {
		h.Each(yield)
	}
}
