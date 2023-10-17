package hashmap

import (
	"fmt"
	"strings"

	"github.com/tidwall/hashmap"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
)

// HashMap 哈希表
type HashMap[K, V any] struct {
	data *hashmap.Map[uint64, pair.Pair[K, V]]
}

func NewHashMap[K, V any]() HashMap[K, V] {
	var data hashmap.Map[uint64, pair.Pair[K, V]]
	return HashMap[K, V]{
		data: &data,
	}
}

func NewHashMapWithCapacity[K, V any](cap uint) HashMap[K, V] {
	data := hashmap.New[uint64, pair.Pair[K, V]](int(cap))
	return HashMap[K, V]{
		data: data,
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
	return uint(self.data.Len())
}

func (self HashMap[K, V]) Equal(dst HashMap[K, V]) bool {
	if self.Length() != dst.Length() {
		return false
	}

	for _, hash := range self.data.Keys() {
		p2, ok := dst.data.Get(hash)
		if !ok {
			return false
		}
		p1, _ := self.data.Get(hash)
		if !stlbasic.Equal(p1.First, p2.First) || !stlbasic.Equal(p1.Second, p2.Second) {
			return false
		}
	}
	return true
}

func (self HashMap[K, V]) Get(k K) V {
	p, _ := self.data.Get(stlbasic.Hash(k))
	return p.Second
}

func (self HashMap[K, V]) ContainKey(k K) bool {
	_, ok := self.data.Get(stlbasic.Hash(k))
	return ok
}

func (self *HashMap[K, V]) Set(k K, v V) V {
	p, _ := self.data.Set(stlbasic.Hash(k), pair.NewPair(k, v))
	return p.Second
}

func (self *HashMap[K, V]) Remove(k K) V {
	p, _ := self.data.Delete(stlbasic.Hash(k))
	return p.Second
}

func (self HashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	for i, p := range self.data.Values() {
		buf.WriteString(fmt.Sprintf("%v", p.First))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", p.Second))
		if i < self.data.Len()-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self HashMap[K, V]) Debug(prefix uint) string {
	var buf strings.Builder
	buf.WriteString("hashmap{")
	for i, p := range self.data.Values() {
		buf.WriteString(stlbasic.Debug(p.First, prefix))
		buf.WriteString(": ")
		buf.WriteString(stlbasic.Debug(p.Second, prefix))
		if i < self.data.Len()-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self HashMap[K, V]) Clone() HashMap[K, V] {
	return HashMap[K, V]{data: self.data.Copy()}
}

func (self *HashMap[K, V]) Clear() {
	var data hashmap.Map[uint64, pair.Pair[K, V]]
	self.data = &data
}

func (self HashMap[K, V]) Empty() bool {
	return self.data.Len() == 0
}

func (self HashMap[K, V]) Iterator() iterator.Iterator[pair.Pair[K, V]] {
	return iterator.NewIterator[pair.Pair[K, V]](_NewIterator[K, V](&self))
}

func (self HashMap[K, V]) Keys() dynarray.DynArray[K] {
	da := dynarray.NewDynArrayWithCapacity[K](self.Length())
	var i uint
	for _, p := range self.data.Values() {
		da.Set(i, p.First)
		i++
	}
	return da
}

func (self HashMap[K, V]) Values() dynarray.DynArray[V] {
	da := dynarray.NewDynArrayWithCapacity[V](self.Length())
	var i uint
	for _, p := range self.data.Values() {
		da.Set(i, p.Second)
		i++
	}
	return da
}
