package hashmap

import (
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/iterator"
	"github.com/kkkunny/stl/container/pair"
)

// HashMap 哈希表
type HashMap[K comparable, V any] struct {
	data map[K]V
}

func NewHashMap[K comparable, V any]() HashMap[K, V] {
	return HashMap[K, V]{data: make(map[K]V)}
}

func NewHashMapWithCapacity[K comparable, V any](cap uint) HashMap[K, V] {
	return HashMap[K, V]{data: make(map[K]V, cap)}
}

func NewDynArrayWith[K comparable, V any](vs ...any) HashMap[K, V] {
	self := NewHashMapWithCapacity[K, V](uint(len(vs) / 2))
	// TODO
	return self
}

func (_ HashMap[K, V]) NewWithIterator(iter iterator.Iterator[HashMap[K, V], pair.Pair[K, V]]) HashMap[K, V] {
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

func (self HashMap[K, V]) Equal(dst any) bool {
	hm, ok := dst.(HashMap[K, V])
	if !ok {
		return false
	}

	if self.Length() != hm.Length() {
		return false
	}

	for k, v := range self.data {
		dv, ok := hm.data[k]
		if !ok{
			return false
		}
		if !stlbasic.Equal(v, dv) {
			return false
		}
	}
	return true
}

func (self HashMap[K, V]) Get(k K) V {
	return self.data[k]
}

func (self HashMap[K, V]) ContainKey(k K) bool {
	_, ok := self.data[k]
	return ok
}

func (self *HashMap[K, V]) Set(k K, v V) V {
	pv := self.data[k]
	self.data[k] = v
	return pv
}

func (self *HashMap[K, V]) Remove(k K) V {
	pv := self.data[k]
	delete(self.data, k)
	return pv
}

func (self HashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var i int
	for k, v := range self.data {
		buf.WriteString(fmt.Sprintf("%v", k))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < len(self.data)-1 {
			buf.WriteString(", ")
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self HashMap[K, V]) Clone() any {
	hm := NewHashMapWithCapacity[K, V](self.Length())
	for k, v := range self.data {
		hm.Set(k, v)
	}
	return hm
}

func (self *HashMap[K, V]) Clear() {
	self.data = make(map[K]V)
}

func (self HashMap[K, V]) Empty() bool {
	return self.Length() == 0
}

func (self HashMap[K, V]) Iterator() iterator.Iterator[HashMap[K, V], pair.Pair[K, V]] {
	return iterator.NewIterator[HashMap[K, V], pair.Pair[K, V]](_NewIterator[K, V](&self))
}
