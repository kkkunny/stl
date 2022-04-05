package table

import (
	"fmt"
	"strings"

	"github.com/kkkunny/stl/list"
	"github.com/kkkunny/stl/tree"
	. "github.com/kkkunny/stl/types"
)

type treeMapNode[K Comparator[K], V any] Entry[K, V]

func (self treeMapNode[K, V]) Compare(dst treeMapNode[K, V]) int {
	return self.Key.Compare(dst.Key)
}

// 有序表
type TreeMap[K Comparator[K], V any] struct {
	data   *tree.RBTree[treeMapNode[K, V]]
	length Usize
}

// 新建有序表
func NewTreeMap[K Comparator[K], V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{data: tree.NewRBTree[treeMapNode[K, V]]()}
}

// 转成字符串 O(N)
func (self *TreeMap[K, V]) String() string {
	var buf strings.Builder
	var i Usize
	buf.WriteByte('{')
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		buf.WriteString(fmt.Sprintf("%v: %v", iter.Key(), iter.Value()))
		if i < self.length-1 {
			buf.WriteString(", ")
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度 O(1)
func (self *TreeMap[K, V]) Length() Usize {
	return self.length
}

// 是否为空 O(1)
func (self *TreeMap[K, V]) Empty() bool {
	return self.length == 0
}

// 设置键值对 O(logN)
func (self *TreeMap[K, V]) Set(k K, v V) {
	node := treeMapNode[K, V]{Key: k, Value: v}
	findNode := self.data.Find(node)
	if findNode == nil {
		self.length++
		self.data.Add(node)
	} else {
		findNode.Value.Value = v
	}
}

// 获取值 O(logN)
func (self *TreeMap[K, V]) Get(k K, d V) V {
	node := self.data.Find(treeMapNode[K, V]{Key: k})
	if node == nil {
		return d
	}
	return node.Value.Value
}

// 移除键值对 O(logN)
func (self *TreeMap[K, V]) Remove(k K, d V) V {
	node := treeMapNode[K, V]{Key: k}
	findNode := self.data.Find(node)
	if findNode == nil {
		return d
	}
	self.length--
	value := findNode.Value.Value
	self.data.Delete(node)
	return value
}

// 是否存在键 O(logN)
func (self *TreeMap[K, V]) ContainKey(k K) bool {
	node := self.data.Find(treeMapNode[K, V]{Key: k})
	return node != nil
}

// 获取键 O(N)
func (self *TreeMap[K, V]) Keys() *list.ArrayList[K] {
	keys := list.NewArrayList[K](self.Length(), self.Length())
	var index Usize
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		keys.Set(index, iter.Key())
		index++
	}
	return keys
}

// 获取值 O(N)
func (self *TreeMap[K, V]) Values() *list.ArrayList[V] {
	values := list.NewArrayList[V](self.Length(), self.Length())
	var index Usize
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		values.Set(index, iter.Value())
		index++
	}
	return values
}

// 清空 O(1)
func (self *TreeMap[K, V]) Clear() {
	self.data = tree.NewRBTree[treeMapNode[K, V]]()
	self.length = 0
}

// 克隆 O(NlogN)
func (self *TreeMap[K, V]) Clone() *TreeMap[K, V] {
	newTree := NewTreeMap[K, V]()
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		newTree.Set(iter.Key(), iter.Value())
	}
	return newTree
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

// 获取起始迭代器
func (self *TreeMap[K, V]) Begin() *TreeMapIterator[K, V] {
	l := list.NewLinkedList[treeMapNode[K, V]]()
	var mid func(node *tree.RBTreeNode[treeMapNode[K, V]])
	mid = func(node *tree.RBTreeNode[treeMapNode[K, V]]) {
		if node.Left != nil {
			mid(node.Left)
		}
		l.Add(node.Value)
		if node.Right != nil {
			mid(node.Right)
		}
	}
	mid(self.data.Root)
	return &TreeMapIterator[K, V]{iter: l.Begin()}
}

// 获取结束迭代器
func (self *TreeMap[K, V]) End() *TreeMapIterator[K, V] {
	l := list.NewLinkedList[treeMapNode[K, V]]()
	var mid func(node *tree.RBTreeNode[treeMapNode[K, V]])
	mid = func(node *tree.RBTreeNode[treeMapNode[K, V]]) {
		if node.Left != nil {
			mid(node.Left)
		}
		l.Add(node.Value)
		if node.Right != nil {
			mid(node.Right)
		}
	}
	mid(self.data.Root)
	return &TreeMapIterator[K, V]{iter: l.End()}
}

// 迭代器
type TreeMapIterator[K Comparator[K], V any] struct {
	iter *list.LinkedListIterator[treeMapNode[K, V]]
}

// 是否存在值
func (self *TreeMapIterator[K, V]) HasValue() bool {
	return self.iter.HasValue()
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
	value := self.iter.Value()
	return value.Key
}

// 获取值
func (self *TreeMapIterator[K, V]) Value() V {
	value := self.iter.Value()
	return value.Value
}
