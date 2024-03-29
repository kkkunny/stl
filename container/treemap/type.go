package treemap

import (
	"github.com/HuKeping/rbtree"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/pair"
)

type entry[K, V any] pair.Pair[K, V]

func (self entry[K, V]) Equal(dst entry[K, V]) bool {
	return stlbasic.Equal(self.First, dst.First)
}

func (self entry[K, V]) Order(dst entry[K, V]) int {
	return stlbasic.Order(self.First, dst.First)
}

func (self *entry[K, V]) Less(dst rbtree.Item) bool {
	return self.Order(*dst.(*entry[K, V])) < 0
}

// TreeMap 有序表
type TreeMap[K, V any] struct {
	tree *rbtree.Rbtree
}

func NewTreeMap[K, V any]() TreeMap[K, V] {
	return TreeMap[K, V]{tree: rbtree.New()}
}

func NewTreeMapWith[K, V any](vs ...any) TreeMap[K, V] {
	self := NewTreeMap[K, V]()
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

func (self TreeMap[K, V]) Default() TreeMap[K, V] {
	return NewTreeMap[K, V]()
}
