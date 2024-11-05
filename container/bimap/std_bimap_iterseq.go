//go:build goexperiment.rangefunc || go1.23

package bimap

import "iter"

func (self *_StdBiMap[T, E]) Iter2() iter.Seq2[T, E] {
	return self.keys.Iter2()
}
