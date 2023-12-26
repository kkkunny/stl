//go:build goexperiment.rangefunc

package linkedhashmap

import "iter"

func (self LinkedHashMap[K, V]) IterSeq() iter.Seq2[K, V] {
	self.init()
	return func(yield func(K, V) bool) {
		for i := self.Iterator(); i.Next(); {
			elem := i.Value()
			if !yield(elem.First, elem.Second) {
				return
			}
		}
	}
}
