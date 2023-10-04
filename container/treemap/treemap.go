package treemap

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
	"github.com/kkkunny/stl/container/tree/btree"
)

type _TreeMapEntry[K, V any] pair.Pair[K, V]

func (self _TreeMapEntry[K, V]) Equal(dst _TreeMapEntry[K, V]) bool{
    return stlbasic.Equal(self.First, dst.First)
}

func (self _TreeMapEntry[K, V]) Order(dst _TreeMapEntry[K, V]) int{
    return stlbasic.Order(self.First, dst.First)
}

// TreeMap 有序表
type TreeMap[K, V any] struct {
	tree btree.BTree[_TreeMapEntry[K, V]]
}

func NewTreeMap[K, V any]() TreeMap[K, V] {
	return TreeMap[K, V]{tree: btree.NewBTree[_TreeMapEntry[K, V]]()}
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
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self TreeMap[K, V]) Length() uint {
	return self.tree.Length()
}

func (self TreeMap[K, V]) Equal(dst TreeMap[K, V]) bool {
	if self.Length() != dst.Length() {
		return false
	}

    equal := true
    dstIter := dst.Iterator()
    self.Iterator().Foreach(func(v pair.Pair[K, V]) bool{
        dstIter.Next()
        dv := dstIter.Value()

		if !stlbasic.Equal(v.First, dv.First) {
            equal = false
			return equal
		}else if !stlbasic.Equal(v.Second, dv.Second) {
            equal = false
			return equal
		}
        return equal
    })
	return equal
}

func (self TreeMap[K, V]) Get(k K) V {
    node := self.tree.Find(_TreeMapEntry[K, V]{First: k})
    if node == nil{
        var v V
        return v
    }
	return node.Value.Second
}

func (self TreeMap[K, V]) ContainKey(k K) bool {
	return self.tree.Find(_TreeMapEntry[K, V]{First: k}) != nil
}

func (self *TreeMap[K, V]) Set(k K, v V) V {
	preNode := self.tree.Find(_TreeMapEntry[K, V]{First: k})
	self.tree.Push(_TreeMapEntry[K, V]{First: k, Second: v})
	if preNode == nil{
		var pv V
		return pv
	}
	pv := preNode.Value.Second
	preNode.Value.Second = v
	return pv
}

func (self *TreeMap[K, V]) Remove(k K) V {
	node := self.tree.Remove(_TreeMapEntry[K, V]{First: k})
    if node == nil{
        var v V
        return v
    }
    return node.Value.Second
}

func (self TreeMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var i uint
    iter := self.Iterator()
    iter.Foreach(func(v pair.Pair[K, V]) bool{
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

func (self TreeMap[K, V]) Clone() TreeMap[K, V] {
	tm := NewTreeMap[K, V]()
    self.Iterator().Foreach(func(v pair.Pair[K, V]) bool{
		tm.Set(v.First, v.Second)
        return true
    })
	return tm
}

func (self *TreeMap[K, V]) Clear() {
	self.tree = btree.NewBTree[_TreeMapEntry[K, V]]()
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
	self.Iterator().Foreach(func(v pair.Pair[K, V]) bool{
		da.Set(i, v.First)
        i++
        return true
    })
	return da
}

func (self TreeMap[K, V]) Values() dynarray.DynArray[V] {
	da := dynarray.NewDynArrayWithCapacity[V](self.Length())
	var i uint
	self.Iterator().Foreach(func(v pair.Pair[K, V]) bool{
		da.Set(i, v.Second)
        i++
        return true
    })
	return da
}
