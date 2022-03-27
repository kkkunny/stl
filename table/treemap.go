package table

import (
	"fmt"
	"stl/list"
	"stl/tree"
	. "stl/types"
	"strings"
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

// 转成字符串
func (self *TreeMap[K, V]) String() string {
	var buf strings.Builder
	var i Usize
	buf.WriteByte('{')
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		k, v := iter.Value()
		buf.WriteString(fmt.Sprintf("%v: %v", k, v))
		if i < self.length-1 {
			buf.WriteString(", ")
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度
func (self *TreeMap[K, V]) Length() Usize {
	return self.length
}

// 是否为空
func (self *TreeMap[K, V]) Empty() bool {
	return self.length == 0
}

// 设置键值对
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

// 获取值
func (self *TreeMap[K, V]) Get(k K, d V) V {
	node := self.data.Find(treeMapNode[K, V]{Key: k})
	if node == nil {
		return d
	}
	return node.Value.Value
}

// 移除键值对
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

// 是否存在键
func (self *TreeMap[K, V]) ContainKey(k K) bool {
	node := self.data.Find(treeMapNode[K, V]{Key: k})
	return node != nil
}

// 获取键
func (self *TreeMap[K, V]) Keys() *list.ArrayList[K] {
	keys := list.NewArrayList[K](self.Length(), self.Length())
	var index Usize
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		k, _ := iter.Value()
		keys.Set(index, k)
		index++
	}
	return keys
}

// 获取值
func (self *TreeMap[K, V]) Values() *list.ArrayList[V] {
	values := list.NewArrayList[V](self.Length(), self.Length())
	var index Usize
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		_, v := iter.Value()
		values.Set(index, v)
		index++
	}
	return values
}

// 清空
func (self *TreeMap[K, V]) Clear() {
	self.data = tree.NewRBTree[treeMapNode[K, V]]()
	self.length = 0
}

// 克隆
func (self *TreeMap[K, V]) Clone() *TreeMap[K, V] {
	newTree := NewTreeMap[K, V]()
	for iter := self.Begin(); iter.HasValue(); iter.Next() {
		k, v := iter.Value()
		newTree.Set(k, v)
	}
	return newTree
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

// 获取值
func (self *TreeMapIterator[K, V]) Value() (K, V) {
	value := self.iter.Value()
	return value.Key, value.Value
}
