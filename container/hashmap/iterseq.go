//go:build goexperiment.rangefunc

package hashmap

import "iter"

func (self HashMap[K, V]) IterSeq() iter.Seq2[K, V] {
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
