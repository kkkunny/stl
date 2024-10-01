//go:build goexperiment.rangefunc || go1.23

package treemap

import (
	"iter"

	"github.com/HuKeping/rbtree"
)

func (self *_AnyTreeMap[K, V]) Iter2() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		self.data.Ascend(self.data.Min(), func(item rbtree.Item) bool {
			node := item.(*anyTreeMapEntry[K, V])
			return yield(node.First, node.Second)
		})
	}
}
