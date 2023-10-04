package treemap

import (
	"github.com/kkkunny/stl/container/pair"
	"github.com/kkkunny/stl/container/tree/btree"
)

type _iterator[K, V any] struct {
	src  *TreeMap[K, V]
	nodes []*_TreeMapEntry[K, V]
	next uint
}

func _NewIterator[K, V any](src *TreeMap[K, V]) *_iterator[K, V] {
	nodes := make([]*_TreeMapEntry[K, V], src.Length())
	if cursor := src.tree.Top(); cursor != nil{
		stack := make([]*btree.BTreeNode[_TreeMapEntry[K, V]], 0, src.Length()/2)
		var i int
		for cursor!=nil || len(stack) != 0 {
			for cursor!=nil{
				stack = append(stack, cursor)
				cursor = cursor.Left
			}
			if len(stack) != 0{
				cursor = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				nodes[i] = &cursor.Value
				i++
				cursor = cursor.Right
			}
		}
	}
	return &_iterator[K, V]{
		src:  src,
		nodes: nodes,
		next: 0,
	}
}

func (self *_iterator[K, V]) Length() uint {
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
