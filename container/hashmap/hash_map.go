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
    keys   map[uint64]K
    values map[uint64]V
}

func NewHashMap[K, V any]() HashMap[K, V] {
    return HashMap[K, V]{
        keys:   make(map[uint64]K),
        values: make(map[uint64]V),
    }
}

func NewHashMapWithCapacity[K, V any](cap uint) HashMap[K, V] {
    return HashMap[K, V]{
        keys:   make(map[uint64]K, cap),
        values: make(map[uint64]V, cap),
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
    return uint(len(self.keys))
}

func (self HashMap[K, V]) Equal(dst HashMap[K, V]) bool {
    if self.Length() != dst.Length() {
        return false
    }

    for hash, v := range self.values {
        dv, ok := dst.values[hash]
        if !ok {
            return false
        }
        if !stlbasic.Equal(v, dv) {
            return false
        }
    }
    return true
}

func (self HashMap[K, V]) Get(k K) V {
    return self.values[stlbasic.Hash(k)]
}

func (self HashMap[K, V]) ContainKey(k K) bool {
    _, ok := self.keys[stlbasic.Hash(k)]
    return ok
}

func (self *HashMap[K, V]) Set(k K, v V) V {
    hash := stlbasic.Hash(k)
    pv := self.values[hash]
    self.keys[hash] = k
    self.values[hash] = v
    return pv
}

func (self *HashMap[K, V]) Remove(k K) V {
    hash := stlbasic.Hash(k)
    pv := self.values[hash]
    delete(self.keys, hash)
    delete(self.values, hash)
    return pv
}

func (self HashMap[K, V]) String() string {
    var buf strings.Builder
    buf.WriteByte('{')
    var i int
    for h, k := range self.keys {
        v := self.values[h]
        buf.WriteString(fmt.Sprintf("%v", k))
        buf.WriteString(": ")
        buf.WriteString(fmt.Sprintf("%v", v))
        if i < len(self.keys)-1 {
            buf.WriteString(", ")
        }
        i++
    }
    buf.WriteByte('}')
    return buf.String()
}

func (self HashMap[K, V]) Clone() HashMap[K, V] {
    hm := NewHashMapWithCapacity[K, V](self.Length())
    for h, k := range self.keys {
        hm.keys[h] = k
    }
    for h, v := range self.values {
        hm.values[h] = v
    }
    return hm
}

func (self *HashMap[K, V]) Clear() {
    self.keys = make(map[uint64]K)
    self.values = make(map[uint64]V)
}

func (self HashMap[K, V]) Empty() bool {
    return self.Length() == 0
}

func (self HashMap[K, V]) Iterator() iterator.Iterator[pair.Pair[K, V]] {
    return iterator.NewIterator[pair.Pair[K, V]](_NewIterator[K, V](&self))
}

func (self HashMap[K, V]) Keys() dynarray.DynArray[K] {
    da := dynarray.NewDynArrayWithCapacity[K](self.Length())
    var i uint
    for _, k := range self.keys {
        da.Set(i, k)
        i++
    }
    return da
}

func (self HashMap[K, V]) Values() dynarray.DynArray[V] {
    da := dynarray.NewDynArrayWithCapacity[V](self.Length())
    var i uint
    for _, v := range self.values {
        da.Set(i, v)
        i++
    }
    return da
}
