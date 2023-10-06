package treeset

import "github.com/HuKeping/rbtree"

type _iterator[T any] struct {
	src  *TreeSet[T]
	nodes []_TreeSetItem[T]
	next uint
}

func _NewIterator[T any](src *TreeSet[T]) *_iterator[T] {
    nodes := make([]_TreeSetItem[T], src.Length())
	var i int
	src.tree.Ascend(src.tree.Min(), func(item rbtree.Item) bool{
		nodes[i] = item.(_TreeSetItem[T])
		i++
		return true
	})
	return &_iterator[T]{
		src:  src,
		nodes: nodes,
		next: 0,
	}
}

func (self _iterator[T]) Length() uint {
	return uint(len(self.nodes))
}

func (self *_iterator[T]) Next() bool {
	if self.next >= self.Length() {
		return false
	}
	self.next++
	return true
}

func (self _iterator[T]) HasNext() bool {
	return self.next < self.Length()
}

func (self _iterator[T]) Value() T {
	return self.nodes[self.next-1].value
}

func (self *_iterator[T]) Reset() {
	self.next = 0
}
