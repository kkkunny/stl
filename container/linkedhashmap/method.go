package linkedhashmap

import (
	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/pair"
	"github.com/kkkunny/stl/internal/list"
)

// 初始化
func (self *LinkedHashMap[K, V]) init() {
	if self.list != nil {
		return
	}
	self.list = list.New[pair.Pair[K, V]]()
}

// Set 插入键值对
func (self *LinkedHashMap[K, V]) Set(k K, v V) V {
	self.init()

	if node := self.HashMap.Get(k); node != nil {
		pv := node.Value.Second
		node.Value = pair.Pair[K, V]{First: k, Second: v}
		self.list.MoveToBack(node)
		return pv
	}
	self.HashMap.Set(k, self.list.PushBack(pair.Pair[K, V]{First: k, Second: v}))
	var pv V
	return pv
}

// Get 获取值
func (self LinkedHashMap[K, V]) Get(k K, defaultValue ...V) V {
	self.init()

	node := self.HashMap.Get(k)
	if node == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if node == nil {
		var v V
		return v
	}
	return node.Value.Second
}

// ContainKey 是否包含键
func (self LinkedHashMap[K, V]) ContainKey(k K) bool {
	self.init()
	return self.HashMap.ContainKey(k)
}

// Remove 移除键值对
func (self *LinkedHashMap[K, V]) Remove(k K, defaultValue ...V) V {
	self.init()

	node := self.HashMap.Remove(k)
	if node == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if node == nil {
		var v V
		return v
	}
	return self.list.Remove(node).Second
}

// Clear 清空
func (self *LinkedHashMap[K, V]) Clear() {
	self.list = nil
	self.HashMap.Clear()
	self.init()
}

// Empty 是否为空
func (self LinkedHashMap[K, V]) Empty() bool {
	self.init()
	return self.list.Len() == 0
}

// Keys 获取所有键
func (self LinkedHashMap[K, V]) Keys() dynarray.DynArray[K] {
	self.init()

	keys := dynarray.NewDynArrayWithLength[K](self.Length())
	var i uint
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		keys.Set(i, cursor.Value.First)
		i++
	}
	return keys
}

// Values 获取所有值
func (self LinkedHashMap[K, V]) Values() dynarray.DynArray[V] {
	self.init()

	values := dynarray.NewDynArrayWithLength[V](self.Length())
	var i uint
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		values.Set(i, cursor.Value.Second)
		i++
	}
	return values
}

// KeyValues 获取所有键值对
func (self LinkedHashMap[K, V]) KeyValues() dynarray.DynArray[pair.Pair[K, V]] {
	self.init()

	pairs := dynarray.NewDynArrayWithLength[pair.Pair[K, V]](self.Length())
	var i uint
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		pairs.Set(i, pair.NewPair(cursor.Value.First, cursor.Value.Second))
		i++
	}
	return pairs
}
