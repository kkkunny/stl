package treemap

import (
	"github.com/HuKeping/rbtree"
	"github.com/kkkunny/stl/container/pair"
)

type _iterator[K, V any] struct {
	src  *TreeMap[K, V]
	nodes []*_TreeMapEntry[K, V]
	next uint
}

func _NewIterator[K, V any](src *TreeMap[K, V]) *_iterator[K, V] {
	nodes := make([]*_TreeMapEntry[K, V], src.Length())
	var i int
	src.tree.Ascend(src.tree.Min(), func(item rbtree.Item) bool{
		nodes[i] = item.(*_TreeMapEntry[K, V])
		i++
		return true
	})
	return &_iterator[K, V]{
		src:  src,
		nodes: nodes,
		next: 0,
	}
}

func (self _iterator[K, V]) Length() uint {
	return uint(len(self.nodes))
}

func (self *_iterator[K, V]) Next() bool {
	if self.next >= self.Length() {
		return false
	}
	self.next++
	return true
}

func (self _iterator[K, V]) HasNext() bool {
	return self.next < self.Length()
}

func (self _iterator[K, V]) Value() pair.Pair[K, V] {
	return pair.Pair[K, V](*self.nodes[self.next-1])
}

func (self *_iterator[K, V]) Reset() {
	self.next = 0
}
