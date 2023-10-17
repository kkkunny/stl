package treemap

import (
	"fmt"
	"strings"

	"github.com/HuKeping/rbtree"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
)

type _TreeMapEntry[K, V any] pair.Pair[K, V]

func (self _TreeMapEntry[K, V]) Equal(dst _TreeMapEntry[K, V]) bool {
	return stlbasic.Equal(self.First, dst.First)
}

func (self _TreeMapEntry[K, V]) Order(dst _TreeMapEntry[K, V]) int {
	return stlbasic.Order(self.First, dst.First)
}

func (self *_TreeMapEntry[K, V]) Less(dst rbtree.Item) bool {
	return self.Order(*dst.(*_TreeMapEntry[K, V])) < 0
}

// TreeMap 有序表
type TreeMap[K, V any] struct {
	tree *rbtree.Rbtree
}

func NewTreeMap[K, V any]() TreeMap[K, V] {
	return TreeMap[K, V]{tree: rbtree.New()}
}

func NewTreeMapWith[K, V any](vs ...any) TreeMap[K, V] {
	self := NewTreeMap[K, V]()
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

func (_ TreeMap[K, V]) NewWithIterator(iter iterator.Iterator[pair.Pair[K, V]]) TreeMap[K, V] {
	self := NewTreeMap[K, V]()
	for iter.Next() {
		entry := iter.Value()
		self.Set(entry.First, entry.Second)
	}
	return self
}

func (self TreeMap[K, V]) Length() uint {
	return self.tree.Len()
}

func (self TreeMap[K, V]) Equal(dst TreeMap[K, V]) bool {
	if self.Length() != dst.Length() {
		return false
	}

	equal := true
	dstIter := dst.Iterator()
	self.Iterator().Foreach(func(v pair.Pair[K, V]) bool {
		dstIter.Next()
		dv := dstIter.Value()

		if !stlbasic.Equal(v.First, dv.First) {
			equal = false
			return equal
		} else if !stlbasic.Equal(v.Second, dv.Second) {
			equal = false
			return equal
		}
		return equal
	})
	return equal
}

func (self TreeMap[K, V]) Get(k K) V {
	item := self.tree.Get(&_TreeMapEntry[K, V]{First: k})
	if item == nil {
		var pv V
		return pv
	}
	return item.(*_TreeMapEntry[K, V]).Second
}

func (self TreeMap[K, V]) ContainKey(k K) bool {
	return self.tree.Get(&_TreeMapEntry[K, V]{First: k}) != nil
}

func (self *TreeMap[K, V]) Set(k K, v V) V {
	entry := &_TreeMapEntry[K, V]{First: k, Second: v}
	item := self.tree.Get(entry)
	if item == nil {
		self.tree.Insert(entry)
		var pv V
		return pv
	}
	entry = item.(*_TreeMapEntry[K, V])
	pv := entry.Second
	entry.Second = v
	return pv
}

func (self *TreeMap[K, V]) Remove(k K) V {
	item := self.tree.Delete(&_TreeMapEntry[K, V]{First: k})
	if item == nil {
		var v V
		return v
	}
	return item.(*_TreeMapEntry[K, V]).Second
}

func (self TreeMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var i uint
	iter := self.Iterator()
	iter.Foreach(func(v pair.Pair[K, V]) bool {
		buf.WriteString(fmt.Sprintf("%v", v.First))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", v.Second))
		if i < iter.Length()-1 {
			buf.WriteString(", ")
		}
		i++
		return true
	})
	buf.WriteByte('}')
	return buf.String()
}

func (self TreeMap[K, V]) Debug(prefix uint) string {
	var buf strings.Builder
	buf.WriteString("treemap{")
	var i uint
	iter := self.Iterator()
	iter.Foreach(func(v pair.Pair[K, V]) bool {
		buf.WriteString(stlbasic.Debug(v.First, prefix))
		buf.WriteString(": ")
		buf.WriteString(stlbasic.Debug(v.Second, prefix))
		if i < iter.Length()-1 {
			buf.WriteString(", ")
		}
		i++
		return true
	})
	buf.WriteByte('}')
	return buf.String()
}

func (self TreeMap[K, V]) Clone() TreeMap[K, V] {
	tm := NewTreeMap[K, V]()
	self.Iterator().Foreach(func(v pair.Pair[K, V]) bool {
		tm.Set(v.First, v.Second)
		return true
	})
	return tm
}

func (self *TreeMap[K, V]) Clear() {
	self.tree = rbtree.New()
}

func (self TreeMap[K, V]) Empty() bool {
	return self.Length() == 0
}

func (self TreeMap[K, V]) Iterator() iterator.Iterator[pair.Pair[K, V]] {
	return iterator.NewIterator[pair.Pair[K, V]](_NewIterator[K, V](&self))
}

func (self TreeMap[K, V]) Keys() dynarray.DynArray[K] {
	da := dynarray.NewDynArrayWithCapacity[K](self.Length())
	var i uint
	self.Iterator().Foreach(func(v pair.Pair[K, V]) bool {
		da.Set(i, v.First)
		i++
		return true
	})
	return da
}

func (self TreeMap[K, V]) Values() dynarray.DynArray[V] {
	da := dynarray.NewDynArrayWithCapacity[V](self.Length())
	var i uint
	self.Iterator().Foreach(func(v pair.Pair[K, V]) bool {
		da.Set(i, v.Second)
		i++
		return true
	})
	return da
}

func (self TreeMap[K, V]) Back() (K, V) {
	entry := self.tree.Max().(*_TreeMapEntry[K, V])
	return entry.First, entry.Second
}

func (self TreeMap[K, V]) Front() (K, V) {
	entry := self.tree.Min().(*_TreeMapEntry[K, V])
	return entry.First, entry.Second
}
