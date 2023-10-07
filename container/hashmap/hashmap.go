package hashmap

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
)

// HashMap 哈希表
type HashMap[K, V any] struct {
	data map[uint64]pair.Pair[K, V]
}

func NewHashMap[K, V any]() HashMap[K, V] {
	return HashMap[K, V]{
		data: make(map[uint64]pair.Pair[K, V]),
	}
}

func NewHashMapWithCapacity[K, V any](cap uint) HashMap[K, V] {
	return HashMap[K, V]{
		data: make(map[uint64]pair.Pair[K, V], cap),
	}
}

func NewHashMapWith[K, V any](vs ...any) HashMap[K, V] {
	self := NewHashMapWithCapacity[K, V](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

func (_ HashMap[K, V]) NewWithIterator(iter iterator.Iterator[pair.Pair[K, V]]) HashMap[K, V] {
	self := NewHashMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self HashMap[K, V]) Length() uint {
	return uint(len(self.data))
}

func (self HashMap[K, V]) Equal(dst HashMap[K, V]) bool {
	if self.Length() != dst.Length() {
		return false
	}

	for hash, pair1 := range self.data {
		pair2, ok := dst.data[hash]
		if !ok {
			return false
		}
		if !stlbasic.Equal(pair1.First, pair2.First) || !stlbasic.Equal(pair1.Second, pair2.Second) {
			return false
		}
	}
	return true
}

func (self HashMap[K, V]) Get(k K) V {
	return self.data[stlbasic.Hash(k)].Second
}

func (self HashMap[K, V]) ContainKey(k K) bool {
	_, ok := self.data[stlbasic.Hash(k)]
	return ok
}

func (self *HashMap[K, V]) Set(k K, v V) V {
	hash := stlbasic.Hash(k)
	ppair := self.data[hash]
	self.data[hash] = pair.NewPair(k, v)
	return ppair.Second
}

func (self *HashMap[K, V]) Remove(k K) V {
	hash := stlbasic.Hash(k)
	pair := self.data[hash]
	delete(self.data, hash)
	return pair.Second
}

func (self HashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var i int
	for _, pair := range self.data {
		buf.WriteString(fmt.Sprintf("%v", pair.First))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", pair.Second))
		if i < len(self.data)-1 {
			buf.WriteString(", ")
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self HashMap[K, V]) Clone() HashMap[K, V] {
	return HashMap[K, V]{data: stlbasic.Clone(self.data)}
}

func (self *HashMap[K, V]) Clear() {
	self.data = make(map[uint64]pair.Pair[K, V])
}

func (self HashMap[K, V]) Empty() bool {
	return len(self.data) == 0
}

func (self HashMap[K, V]) Iterator() iterator.Iterator[pair.Pair[K, V]] {
	return iterator.NewIterator[pair.Pair[K, V]](_NewIterator[K, V](&self))
}

func (self HashMap[K, V]) Keys() dynarray.DynArray[K] {
	da := dynarray.NewDynArrayWithCapacity[K](self.Length())
	var i uint
	for _, pair := range self.data {
		da.Set(i, pair.First)
		i++
	}
	return da
}

func (self HashMap[K, V]) Values() dynarray.DynArray[V] {
	da := dynarray.NewDynArrayWithCapacity[V](self.Length())
	var i uint
	for _, pair := range self.data {
		da.Set(i, pair.Second)
		i++
	}
	return da
}
