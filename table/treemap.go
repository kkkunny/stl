package table

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"

	"github.com/kkkunny/stl/list"
	"github.com/kkkunny/stl/tree"
)

type treeMapEntry[K constraints.Ordered, V any] Entry[K, V]

func (self treeMapEntry[K, V]) Compare(dst treeMapEntry[K, V]) int {
	if self.Key < dst.Key {
		return -1
	} else if self.Key > dst.Key {
		return 1
	} else {
		return 0
	}
}

// 有序表
type TreeMap[K constraints.Ordered, V any] struct {
	data   *tree.RBTree[treeMapEntry[K, V]]
	length int
}

// 新建有序表
func NewTreeMap[K constraints.Ordered, V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{data: tree.NewRBTree[treeMapEntry[K, V]]()}
}

// 转成字符串 O(N)
func (self *TreeMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		buf.WriteString(fmt.Sprintf("%v: %v", iter.Key(), iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度 O(1)
func (self *TreeMap[K, V]) Length() int {
	return self.length
}

// 是否为空 O(1)
func (self *TreeMap[K, V]) Empty() bool {
	return self.length == 0
}

// 设置键值对 O(logN)
func (self *TreeMap[K, V]) Set(k K, v V) {
	node := treeMapEntry[K, V]{Key: k, Value: v}
	findNode := self.data.Find(node)
	if findNode == nil {
		self.length++
		self.data.Add(node)
	} else {
		findNode.Value.Value = v
	}
}

// 获取值 O(logN)
func (self *TreeMap[K, V]) Get(k K, v ...V) V {
	node := self.data.Find(treeMapEntry[K, V]{Key: k})
	if node != nil {
		return node.Value.Value
	} else if len(v) == 0 {
		var vv V
		return vv
	} else {
		return v[0]
	}
}

// 移除键值对 O(logN)
func (self *TreeMap[K, V]) Remove(k K, v ...V) V {
	node := treeMapEntry[K, V]{Key: k}
	findNode := self.data.Find(node)
	if findNode == nil {
		if len(v) == 0 {
			var vv V
			return vv
		} else {
			return v[0]
		}
	}
	self.length--
	value := findNode.Value.Value
	self.data.Delete(node)
	return value
}

// 是否存在键 O(logN)
func (self *TreeMap[K, V]) ContainKey(k K) bool {
	return self.data.Find(treeMapEntry[K, V]{Key: k}) != nil
}

// 获取键 O(N)
func (self *TreeMap[K, V]) Keys() *list.ArrayList[K] {
	keys := list.NewArrayList[K](self.Length(), self.Length())
	var index int
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		keys.Set(index, iter.Key())
		index++
	}
	return keys
}

// 获取值 O(N)
func (self *TreeMap[K, V]) Values() *list.ArrayList[V] {
	values := list.NewArrayList[V](self.Length(), self.Length())
	var index int
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		values.Set(index, iter.Value())
		index++
	}
	return values
}

// 清空 O(1)
func (self *TreeMap[K, V]) Clear() {
	if self.Empty() {
		return
	}
	self.data = tree.NewRBTree[treeMapEntry[K, V]]()
	self.length = 0
}

// 克隆 O(NlogN)
func (self *TreeMap[K, V]) Clone() *TreeMap[K, V] {
	tm := NewTreeMap[K, V]()
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		tm.Set(iter.Key(), iter.Value())
	}
	return tm
}

// 过滤 O(N)
func (self *TreeMap[K, V]) Filter(f func(k K, v V) bool) *TreeMap[K, V] {
	tm := NewTreeMap[K, V]()
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		key, value := iter.Key(), iter.Value()
		if f(key, value) {
			tm.Set(key, value)
		}
	}
	return tm
}

// 任意一个满足条件 O(N)
func (self *TreeMap[K, V]) Any(f func(k K, v V) bool) bool {
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		if f(iter.Key(), iter.Value()) {
			return true
		}
	}
	return false
}

// 全部满足条件 O(N)
func (self *TreeMap[K, V]) Every(f func(k K, v V) bool) bool {
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		if !f(iter.Key(), iter.Value()) {
			return false
		}
	}
	return true
}

// 获取起始迭代器
func (self *TreeMap[K, V]) Begin() *TreeMapIterator[K, V] {
	ll := list.NewDoubleLinkedList[treeMapEntry[K, V]]()
	var mid func(node *tree.RBTreeNode[treeMapEntry[K, V]])
	mid = func(node *tree.RBTreeNode[treeMapEntry[K, V]]) {
		if node.Left != nil {
			mid(node.Left)
		}
		ll.Add(node.Value)
		if node.Right != nil {
			mid(node.Right)
		}
	}
	mid(self.data.Root)
	return &TreeMapIterator[K, V]{iter: ll.Begin()}
}

// 获取结束迭代器
func (self *TreeMap[K, V]) End() *TreeMapIterator[K, V] {
	ll := list.NewDoubleLinkedList[treeMapEntry[K, V]]()
	var mid func(node *tree.RBTreeNode[treeMapEntry[K, V]])
	mid = func(node *tree.RBTreeNode[treeMapEntry[K, V]]) {
		if node.Left != nil {
			mid(node.Left)
		}
		ll.Add(node.Value)
		if node.Right != nil {
			mid(node.Right)
		}
	}
	mid(self.data.Root)
	return &TreeMapIterator[K, V]{iter: ll.End()}
}

// 迭代器
type TreeMapIterator[K constraints.Ordered, V any] struct {
	iter *list.DoubleLinkedListIterator[treeMapEntry[K, V]]
}

// 是否存在值
func (self *TreeMapIterator[K, V]) HasValue() bool {
	return self.iter.HasValue()
}

// 是否存在上一个
func (self *TreeMapIterator[K, V]) HasPrev() bool {
	return self.iter.HasPrev()
}

// 是否存在下一个
func (self *TreeMapIterator[K, V]) HasNext() bool {
	return self.iter.HasNext()
}

// 上一个
func (self *TreeMapIterator[K, V]) Prev() {
	self.iter.Prev()
}

// 下一个
func (self *TreeMapIterator[K, V]) Next() {
	self.iter.Next()
}

// 获取键
func (self *TreeMapIterator[K, V]) Key() K {
	return self.iter.Value().Key
}

// 获取值
func (self *TreeMapIterator[K, V]) Value() V {
	return self.iter.Value().Value
}
