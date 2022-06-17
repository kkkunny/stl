package table

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/list"
)

// 哈希表
type HashMap[K comparable, V any] struct {
	data map[K]V
}

// 新建哈希表
func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		data: make(map[K]V),
	}
}

// 转成字符串 O(N)
func (self *HashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var index int
	for k, v := range self.data {
		buf.WriteString(fmt.Sprintf("%v: %v", k, v))
		if index < len(self.data)-1 {
			buf.WriteString(", ")
		}
		index++
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度 O(1)
func (self *HashMap[K, V]) Length() int {
	return len(self.data)
}

// 是否为空 O(1)
func (self *HashMap[K, V]) Empty() bool {
	return len(self.data) == 0
}

// 设置键值对 O(1)
func (self *HashMap[K, V]) Set(k K, v V) {
	self.data[k] = v
}

// 获取值 O(1)
func (self *HashMap[K, V]) Get(k K, v ...V) V {
	if vv, ok := self.data[k]; ok {
		return vv
	} else if len(v) == 0 {
		return vv
	} else {
		return v[0]
	}
}

// 移除键值对 O(1)
func (self *HashMap[K, V]) Remove(k K, v ...V) V {
	if vv, ok := self.data[k]; ok {
		delete(self.data, k)
		return vv
	} else if len(v) == 0 {
		return vv
	} else {
		return v[0]
	}
}

// 是否存在键 O(1)
func (self *HashMap[K, V]) ContainKey(k K) bool {
	_, ok := self.data[k]
	return ok
}

// 获取键 O(N)
func (self *HashMap[K, V]) Keys() *list.ArrayList[K] {
	keys := list.NewArrayList[K](len(self.data), len(self.data))
	var index int
	for k := range self.data {
		keys.Set(index, k)
		index++
	}
	return keys
}

// 获取值 O(N)
func (self *HashMap[K, V]) Values() *list.ArrayList[V] {
	values := list.NewArrayList[V](len(self.data), len(self.data))
	var index int
	for _, v := range self.data {
		values.Set(index, v)
		index++
	}
	return values
}

// 清空 O(1)
func (self *HashMap[K, V]) Clear() {
	if !self.Empty() {
		self.data = make(map[K]V)
	}
}

// 克隆 O(N)
func (self *HashMap[K, V]) Clone() *HashMap[K, V] {
	hm := NewHashMap[K, V]()
	for k, v := range self.data {
		hm.data[k] = v
	}
	return hm
}

// 过滤 O(N)
func (self *HashMap[K, V]) Filter(f func(k K, v V) bool) *HashMap[K, V] {
	hm := NewHashMap[K, V]()
	for k, v := range self.data {
		if f(k, v) {
			hm.data[k] = v
		}
	}
	return hm
}

// 任意一个满足条件 O(N)
func (self *HashMap[K, V]) Any(f func(k K, v V) bool) bool {
	for k, v := range self.data {
		if f(k, v) {
			return true
		}
	}
	return false
}

// 每一个满足条件 O(N)
func (self *HashMap[K, V]) Every(f func(k K, v V) bool) bool {
	for k, v := range self.data {
		if !f(k, v) {
			return false
		}
	}
	return true
}

// HashMapMap 映射
func HashMapMap[K1 comparable, V1 any, K2 comparable, V2 any](src *HashMap[K1, V1], f func(k K1, v V1) (K2, V2)) *HashMap[K2, V2] {
	newMap := NewHashMap[K2, V2]()
	for k, v := range src.data {
		k2, v2 := f(k, v)
		newMap.data[k2] = v2
	}
	return newMap
}

// 获取迭代器
func (self *HashMap[K, V]) Iterator() *HashMapIterator[K, V] {
	data := make([]Entry[K, V], len(self.data))
	var index int
	for k, v := range self.data {
		data[index] = Entry[K, V]{
			Key:   k,
			Value: v,
		}
		index++
	}
	return &HashMapIterator[K, V]{
		data:  data,
		index: 0,
	}
}

// 迭代器
type HashMapIterator[K comparable, V any] struct {
	data  []Entry[K, V]
	index int
}

// 是否存在值
func (self *HashMapIterator[K, V]) HasValue() bool {
	return 0 <= self.index && self.index < len(self.data)
}

// 是否存在下一个
func (self *HashMapIterator[K, V]) HasNext() bool {
	return self.index+1 < len(self.data)
}

// 下一个
func (self *HashMapIterator[K, V]) Next() {
	self.index++
}

// 获取键
func (self *HashMapIterator[K, V]) Key() K {
	return self.data[self.index].Key
}

// 获取值
func (self *HashMapIterator[K, V]) Value() V {
	return self.data[self.index].Value
}
