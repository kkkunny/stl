//go:build goexperiment.rangefunc || go1.23

package hashmap

import "iter"

func (self *_ConcurrentMap[K, V]) Iter2() iter.Seq2[K, V] {
	kvs := self.data.IterBuffered()
	return func(yield func(K, V) bool) {
		kv, ok := <-kvs
		if !ok {
			return
		}
		yield(kv.Key, kv.Val)
	}
}
