package treemap

import (
	"github.com/HuKeping/rbtree"
)

// Clone 克隆
func (self TreeMap[K, V]) Clone() TreeMap[K, V] {
	self.init()
	tm := NewTreeMap[K, V]()
	if !self.Empty() {
		self.tree.Ascend(self.tree.Min(), func(item rbtree.Item) bool {
			node := item.(*entry[K, V])
			tm.Set(node.First, node.Second)
			return true
		})
	}
	return tm
}
