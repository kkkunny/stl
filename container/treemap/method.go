package treemap

import (
	"github.com/HuKeping/rbtree"

	"github.com/kkkunny/stl/container/pair"
)

// 初始化
func (self *TreeMap[K, V]) init() {
	if self.tree != nil {
		return
	}
	self.tree = rbtree.New()
}

// Set 插入键值对
func (self *TreeMap[K, V]) Set(k K, v V) V {
	self.init()

	node := &entry[K, V]{First: k, Second: v}
	item := self.tree.Get(node)
	if item == nil {
		self.tree.Insert(node)
		var pv V
		return pv
	}
	node = item.(*entry[K, V])
	pv := node.Second
	node.Second = v
	return pv
}

// Get 获取值
func (self TreeMap[K, V]) Get(k K, defaultValue ...V) V {
	self.init()

	item := self.tree.Get(&entry[K, V]{First: k})
	if item == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if item == nil {
		var v V
		return v
	}
	return item.(*entry[K, V]).Second
}

// ContainKey 是否包含键
func (self TreeMap[K, V]) ContainKey(k K) bool {
	self.init()
	return self.tree.Get(&entry[K, V]{First: k}) != nil
}

// Remove 移除键值对
func (self *TreeMap[K, V]) Remove(k K, defaultValue ...V) V {
	self.init()

	item := self.tree.Delete(&entry[K, V]{First: k})
	if item == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if item == nil {
		var v V
		return v
	}
	return item.(*entry[K, V]).Second
}

// Clear 清空
func (self *TreeMap[K, V]) Clear() {
	self.tree = nil
	self.init()
}

// Empty 是否为空
func (self TreeMap[K, V]) Empty() bool {
	self.init()
	return self.Length() == 0
}

// Keys 获取所有键
func (self TreeMap[K, V]) Keys() []K {
	self.init()

	keys := make([]K, self.Length())
	if !self.Empty() {
		var i uint
		self.tree.Ascend(self.tree.Min(), func(item rbtree.Item) bool {
			keys[i] = item.(*entry[K, V]).First
			i++
			return true
		})
	}
	return keys
}

// Values 获取所有值
func (self TreeMap[K, V]) Values() []V {
	self.init()

	values := make([]V, self.Length())
	if !self.Empty() {
		var i uint
		self.tree.Ascend(self.tree.Min(), func(item rbtree.Item) bool {
			values[i] = item.(*entry[K, V]).Second
			i++
			return true
		})
	}
	return values
}

// KeyValues 获取所有键值对
func (self TreeMap[K, V]) KeyValues() []pair.Pair[K, V] {
	self.init()

	pairs := make([]pair.Pair[K, V], self.Length())
	if !self.Empty() {
		var i uint
		self.tree.Ascend(self.tree.Min(), func(item rbtree.Item) bool {
			node := item.(*entry[K, V])
			pairs[i] = pair.NewPair(node.First, node.Second)
			i++
			return true
		})
	}
	return pairs
}

// Back 末尾的元素
func (self TreeMap[K, V]) Back() (K, V) {
	self.init()
	node := self.tree.Max().(*entry[K, V])
	return node.First, node.Second
}

// Front 开头的元素
func (self TreeMap[K, V]) Front() (K, V) {
	self.init()
	node := self.tree.Min().(*entry[K, V])
	return node.First, node.Second
}
