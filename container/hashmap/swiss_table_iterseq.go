//go:build goexperiment.rangefunc

package hashmap

import "iter"

func (self *_SwissTable[K, V]) IterSeq() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := self.Iterator(); i.Next(); {
			elem := i.Value()
			if !yield(elem.First, elem.Second) {
				return
			}
		}
	}
}