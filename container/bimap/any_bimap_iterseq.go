//go:build goexperiment.rangefunc

package bimap

import "iter"

func (self *_AnyBiMap[T, E]) IterSeq() iter.Seq2[T, E] {
	return func(yield func(T, E) bool) {
		for i := self.Iterator(); i.Next(); {
			elem := i.Value()
			if !yield(elem.First, elem.Second) {
				return
			}
		}
	}
}
