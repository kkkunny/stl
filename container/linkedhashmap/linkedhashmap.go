package linkedhashmap

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
	"github.com/kkkunny/stl/internal/list"
)

// LinkedHashMap 有序哈希表
type LinkedHashMap[K, V any] struct {
    list *list.List[pair.Pair[K, V]]
	data map[uint64]*list.Element[pair.Pair[K, V]]
}

func NewLinkedHashMap[K, V any]() LinkedHashMap[K, V] {
	return LinkedHashMap[K, V]{
		list:   list.New[pair.Pair[K, V]](),
		data: make(map[uint64]*list.Element[pair.Pair[K, V]]),
	}
}

func NewLinkedHashMapWithCapacity[K, V any](cap uint) LinkedHashMap[K, V] {
	return LinkedHashMap[K, V]{
		list:   list.New[pair.Pair[K, V]](),
		data: make(map[uint64]*list.Element[pair.Pair[K, V]], cap),
	}
}

func NewLinkedHashMapWith[K, V any](vs ...any) LinkedHashMap[K, V] {
	self := NewLinkedHashMapWithCapacity[K, V](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

func (_ LinkedHashMap[K, V]) NewWithIterator(iter iterator.Iterator[pair.Pair[K, V]]) LinkedHashMap[K, V] {
	self := NewLinkedHashMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self LinkedHashMap[K, V]) Length() uint {
	return uint(self.list.Len())
}

func (self LinkedHashMap[K, V]) Equal(dst LinkedHashMap[K, V]) bool {
	if self.Length() != dst.Length() {
		return false
	}

	for c1, c2 :=self.list.Front(), dst.list.Front(); c1!=nil&&c2!=nil; c1, c2 = c1.Next(), c2.Next() {
		v1, v2 := c1.Value, c2.Value
		if !stlbasic.Equal(v1.First, v2.First) || !stlbasic.Equal(v1.Second, v2.Second) {
			return false
		}
	}
	return true
}

func (self LinkedHashMap[K, V]) Get(k K) V {
	node := self.data[stlbasic.Hash(k)]
	if node == nil{
		var v V
		return v
	}
	return node.Value.Second
}

func (self LinkedHashMap[K, V]) ContainKey(k K) bool {
	return self.data[stlbasic.Hash(k)] != nil
}

func (self *LinkedHashMap[K, V]) Set(k K, v V) V {
	hash := stlbasic.Hash(k)
	if node := self.data[hash]; node != nil{
		pv := node.Value.Second
		node.Value = pair.Pair[K, V]{First: k, Second: v}
		self.list.MoveToBack(node)
		return pv
	}
	self.data[hash] = self.list.PushBack(pair.Pair[K, V]{First: k, Second: v})
	var pv V
	return pv
}

func (self *LinkedHashMap[K, V]) Remove(k K) V {
    hash := stlbasic.Hash(k)
	node := self.data[hash]
    if node == nil{
        var v V
        return v
    }
	delete(self.data, hash)
	return self.list.Remove(node).Second
}

func (self LinkedHashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var i int
	for cursor:=self.list.Front(); cursor!=nil; cursor=cursor.Next() {
		buf.WriteString(fmt.Sprintf("%v", cursor.Value.First))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", cursor.Value.Second))
		if cursor.Next() != nil {
			buf.WriteString(", ")
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self LinkedHashMap[K, V]) Clone() LinkedHashMap[K, V] {
	hm := NewLinkedHashMapWithCapacity[K, V](self.Length())
	for cursor:=self.list.Front(); cursor!=nil; cursor=cursor.Next(){
		hm.Set(cursor.Value.First, cursor.Value.Second)
	}
	return hm
}

func (self *LinkedHashMap[K, V]) Clear() {
	self.list = list.New[pair.Pair[K, V]]()
	self.data = make(map[uint64]*list.Element[pair.Pair[K, V]])
}

func (self LinkedHashMap[K, V]) Empty() bool {
	return self.Length() == 0
}

func (self LinkedHashMap[K, V]) Iterator() iterator.Iterator[pair.Pair[K, V]] {
	return iterator.NewIterator[pair.Pair[K, V]](_NewIterator[K, V](&self))
}

func (self LinkedHashMap[K, V]) Keys() dynarray.DynArray[K] {
	da := dynarray.NewDynArrayWithCapacity[K](self.Length())
	var i uint
	for cursor:=self.list.Front(); cursor!=nil; cursor=cursor.Next(){
		da.Set(i, cursor.Value.First)
		i++
	}
	return da
}

func (self LinkedHashMap[K, V]) Values() dynarray.DynArray[V] {
	da := dynarray.NewDynArrayWithCapacity[V](self.Length())
	var i uint
	for cursor:=self.list.Front(); cursor!=nil; cursor=cursor.Next(){
		da.Set(i, cursor.Value.Second)
		i++
	}
	return da
}
