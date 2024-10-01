package treemap

import (
	"fmt"
	"strings"

	"github.com/HuKeping/rbtree"

	"github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
	stlslices "github.com/kkkunny/stl/container/slices"
)

type anyTreeMapEntry[K, V any] pair.Pair[K, V]

func (self *anyTreeMapEntry[K, V]) Less(dst rbtree.Item) bool {
	return stlcmp.Compare(self.First, dst.(*anyTreeMapEntry[K, V]).First) < 0
}

type _AnyTreeMap[K, V any] struct {
	data *rbtree.Rbtree
}

func _NewAnyTreeMap[K, V any]() TreeMap[K, V] {
	return &_AnyTreeMap[K, V]{data: rbtree.New()}
}

func _NewAnyTreeMapWith[K, V any](vs ...any) TreeMap[K, V] {
	self := _NewAnyTreeMap[K, V]()
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

// Clone 克隆
func (self *_AnyTreeMap[K, V]) Clone() any {
	tm := _NewAnyTreeMap[K, V]()
	if !self.Empty() {
		self.data.Ascend(self.data.Min(), func(item rbtree.Item) bool {
			node := item.(*anyTreeMapEntry[K, V])
			tm.Set(node.First, node.Second)
			return true
		})
	}
	return tm
}

// Equal 比较
func (self *_AnyTreeMap[K, V]) Equal(dstObj any) bool {
	if dstObj == nil && self == nil {
		return true
	} else if dstObj == nil {
		return false
	}

	dst, ok := dstObj.(TreeMap[K, V])
	if !ok {
		return false
	}

	if self.Length() != dst.Length() {
		return false
	}
	return stlslices.Equal(self.KeyValues(), dst.KeyValues())
}

func (_ *_AnyTreeMap[K, V]) NewWithIterator(iter stliter.Iterator[pair.Pair[K, V]]) any {
	self := _NewAnyTreeMap[K, V]()
	for iter.Next() {
		node := iter.Value()
		self.Set(node.First, node.Second)
	}
	return self
}

func (self *_AnyTreeMap[K, V]) Iterator() stliter.Iterator[pair.Pair[K, V]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_AnyTreeMap[K, V]) Length() uint {
	return self.data.Len()
}

// Set 插入键值对
func (self *_AnyTreeMap[K, V]) Set(k K, v V) V {
	node := &anyTreeMapEntry[K, V]{First: k, Second: v}
	item := self.data.Get(node)
	if item == nil {
		self.data.Insert(node)
		var pv V
		return pv
	}
	node = item.(*anyTreeMapEntry[K, V])
	pv := node.Second
	node.Second = v
	return pv
}

// Get 获取值
func (self *_AnyTreeMap[K, V]) Get(k K, defaultValue ...V) V {
	item := self.data.Get(&anyTreeMapEntry[K, V]{First: k})
	if item == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if item == nil {
		var v V
		return v
	}
	return item.(*anyTreeMapEntry[K, V]).Second
}

// Contain 是否包含键
func (self *_AnyTreeMap[K, V]) Contain(k K) bool {
	return self.data.Get(&anyTreeMapEntry[K, V]{First: k}) != nil
}

// Remove 移除键值对
func (self *_AnyTreeMap[K, V]) Remove(k K, defaultValue ...V) V {
	item := self.data.Delete(&anyTreeMapEntry[K, V]{First: k})
	if item == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if item == nil {
		var v V
		return v
	}
	return item.(*anyTreeMapEntry[K, V]).Second
}

// Clear 清空
func (self *_AnyTreeMap[K, V]) Clear() {
	self.data = rbtree.New()
}

// Empty 是否为空
func (self *_AnyTreeMap[K, V]) Empty() bool {
	return self.Length() == 0
}

// Keys 获取所有键
func (self *_AnyTreeMap[K, V]) Keys() []K {
	keys := make([]K, self.Length())
	if !self.Empty() {
		var i uint
		self.data.Ascend(self.data.Min(), func(item rbtree.Item) bool {
			keys[i] = item.(*anyTreeMapEntry[K, V]).First
			i++
			return true
		})
	}
	return keys
}

// Values 获取所有值
func (self *_AnyTreeMap[K, V]) Values() []V {
	values := make([]V, self.Length())
	if !self.Empty() {
		var i uint
		self.data.Ascend(self.data.Min(), func(item rbtree.Item) bool {
			values[i] = item.(*anyTreeMapEntry[K, V]).Second
			i++
			return true
		})
	}
	return values
}

// KeyValues 获取所有键值对
func (self *_AnyTreeMap[K, V]) KeyValues() []pair.Pair[K, V] {
	pairs := make([]pair.Pair[K, V], self.Length())
	if !self.Empty() {
		var i uint
		self.data.Ascend(self.data.Min(), func(item rbtree.Item) bool {
			node := item.(*anyTreeMapEntry[K, V])
			pairs[i] = pair.NewPair(node.First, node.Second)
			i++
			return true
		})
	}
	return pairs
}

// Back 末尾的元素
func (self *_AnyTreeMap[K, V]) Back() (K, V) {
	node := self.data.Max().(*anyTreeMapEntry[K, V])
	return node.First, node.Second
}

// Front 开头的元素
func (self *_AnyTreeMap[K, V]) Front() (K, V) {
	node := self.data.Min().(*anyTreeMapEntry[K, V])
	return node.First, node.Second
}

// String 转成字符串
func (self *_AnyTreeMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteString("TreeMap{")
	if !self.Empty() {
		var i uint
		self.data.Ascend(self.data.Min(), func(item rbtree.Item) bool {
			node := item.(*anyTreeMapEntry[K, V])
			buf.WriteString(fmt.Sprintf("%v", node.First))
			buf.WriteString(": ")
			buf.WriteString(fmt.Sprintf("%v", node.Second))
			if i < self.Length()-1 {
				buf.WriteString(", ")
			}
			i++
			return true
		})
	}
	buf.WriteByte('}')
	return buf.String()
}
