package bimap

import (
	"github.com/kkkunny/stl/container/pair"
)

// Set 插入键值对
func (self *BiMap[K, V]) Set(k K, v V) (K, V) {
	pv, pk := self.RemoveKey(k), self.RemoveValue(v)
	self.keys.Set(k, v)
	self.values.Set(v, k)
	return pk, pv
}

// GetValue 获取值
func (self BiMap[K, V]) GetValue(k K, defaultValue ...V) V {
	return self.keys.Get(k, defaultValue...)
}

// GetKey 获取键
func (self BiMap[K, V]) GetKey(v V, defaultKey ...K) K {
	return self.values.Get(v, defaultKey...)
}

// ContainKey 是否包含键
func (self BiMap[K, V]) ContainKey(k K) bool {
	return self.keys.ContainKey(k)
}

// ContainValue 是否包含值
func (self BiMap[K, V]) ContainValue(v V) bool {
	return self.values.ContainKey(v)
}

// RemoveKey 移除键
func (self *BiMap[K, V]) RemoveKey(k K, defaultValue ...V) V {
	exist := self.ContainKey(k)
	pv := self.keys.Remove(k, defaultValue...)
	if exist {
		self.values.Remove(pv)
	}
	return pv
}

// RemoveValue 移除值
func (self *BiMap[K, V]) RemoveValue(v V, defaultKey ...K) K {
	exist := self.ContainValue(v)
	pk := self.values.Remove(v, defaultKey...)
	if exist {
		self.keys.Remove(pk)
	}
	return pk
}

// Clear 清空
func (self *BiMap[K, V]) Clear() {
	self.keys.Clear()
	self.values.Clear()
}

// Empty 是否为空
func (self BiMap[K, V]) Empty() bool {
	return self.keys.Empty()
}

// Keys 获取所有键
func (self BiMap[K, V]) Keys() []K {
	return self.keys.Keys()
}

// Values 获取所有值
func (self BiMap[K, V]) Values() []V {
	return self.values.Keys()
}

// KeyValues 获取所有键值对
func (self BiMap[K, V]) KeyValues() []pair.Pair[K, V] {
	return self.keys.KeyValues()
}
