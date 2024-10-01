//go:build goexperiment.rangefunc || go1.23

package bimap

import "iter"

func (self *_AnyBiMap[T, E]) Iter2() iter.Seq2[T, E] {
	return self.keys.Iter2()
}
