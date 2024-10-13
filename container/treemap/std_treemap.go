package treemap

import (
	"cmp"
	"fmt"
	"strings"

	rbtree "github.com/sakeven/RbTree"

	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
	stlslices "github.com/kkkunny/stl/container/slices"
)

type _StdTreeMap[K cmp.Ordered, V any] struct {
	data *rbtree.Tree[K, V]
}

func _NewStdTreeMap[K cmp.Ordered, V any]() TreeMap[K, V] {
	return &_StdTreeMap[K, V]{data: rbtree.NewTree[K, V]()}
}

func _NewStdTreeMapWith[K cmp.Ordered, V any](vs ...any) TreeMap[K, V] {
	self := _NewStdTreeMap[K, V]()
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

// Clone 克隆
func (self *_StdTreeMap[K, V]) Clone() any {
	tm := _NewStdTreeMap[K, V]()
	for iter := self.data.Iterator(); iter != nil; iter = iter.Next() {
		tm.Set(iter.Key, iter.Value)
	}
	return tm
}

// Equal 比较
func (self *_StdTreeMap[K, V]) Equal(dstObj any) bool {
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
	return stlcmp.Equal(self.KeyValues(), dst.KeyValues())
}

func (_ *_StdTreeMap[K, V]) NewWithIterator(iter stliter.Iterator[pair.Pair[K, V]]) any {
	self := _NewStdTreeMap[K, V]()
	for iter.Next() {
		node := iter.Value()
		self.Set(node.First, node.Second)
	}
	return self
}

func (self *_StdTreeMap[K, V]) Iterator() stliter.Iterator[pair.Pair[K, V]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_StdTreeMap[K, V]) Length() uint {
	return uint(self.data.Size())
}

// Set 插入键值对
func (self *_StdTreeMap[K, V]) Set(k K, v V) V {
	var pv V
	node := self.data.FindIt(k)
	if node != nil {
		pv = node.Value
		node.Value = v
		return pv
	}
	self.data.Insert(k, v)
	return pv
}

// Get 获取值
func (self *_StdTreeMap[K, V]) Get(k K, defaultValue ...V) V {
	node := self.data.FindIt(k)
	if node == nil {
		return stlslices.Last(defaultValue)
	}
	return node.Value
}

// Contain 是否包含键
func (self *_StdTreeMap[K, V]) Contain(k K) bool {
	return self.data.FindIt(k) != nil
}

// Remove 移除键值对
func (self *_StdTreeMap[K, V]) Remove(k K, defaultValue ...V) V {
	node := self.data.FindIt(k)
	if node == nil {
		return stlslices.Last(defaultValue)
	}
	self.data.Delete(k)
	return node.Value
}

// Clear 清空
func (self *_StdTreeMap[K, V]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_StdTreeMap[K, V]) Empty() bool {
	return self.data.Empty()
}

// Keys 获取所有键
func (self *_StdTreeMap[K, V]) Keys() []K {
	keys := make([]K, self.Length())
	var i int
	for iter := self.data.Iterator(); iter != nil; iter = iter.Next() {
		keys[i] = iter.Key
		i++
	}
	return keys
}

// Values 获取所有值
func (self *_StdTreeMap[K, V]) Values() []V {
	values := make([]V, self.Length())
	var i int
	for iter := self.data.Iterator(); iter != nil; iter = iter.Next() {
		values[i] = iter.Value
		i++
	}
	return values
}

// KeyValues 获取所有键值对
func (self *_StdTreeMap[K, V]) KeyValues() []pair.Pair[K, V] {
	pairs := make([]pair.Pair[K, V], self.Length())
	var i int
	for iter := self.data.Iterator(); iter != nil; iter = iter.Next() {
		pairs[i] = pair.NewPair(iter.Key, iter.Value)
		i++
	}
	return pairs
}

// Back 末尾的元素
func (self *_StdTreeMap[K, V]) Back() (K, V) {
	iter := self.data.Iterator()
	for ; iter != nil && iter.Next() != nil; iter = iter.Next() {
	}
	return iter.Key, iter.Value
}

// Front 开头的元素
func (self *_StdTreeMap[K, V]) Front() (K, V) {
	iter := self.data.Iterator()
	return iter.Key, iter.Value
}

// String 转成字符串
func (self *_StdTreeMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteString("TreeMap{")
	for iter := self.data.Iterator(); iter != nil; iter = iter.Next() {
		buf.WriteString(fmt.Sprintf("%v", iter.Key))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", iter.Value))
		if iter.Next() != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
