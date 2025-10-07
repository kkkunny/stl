package hashmap

import (
	"fmt"
	"iter"
	"strings"

	cmap "github.com/orcaman/concurrent-map/v2"

	stlbasic "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	stlmaps "github.com/kkkunny/stl/container/maps"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/container/tuple"
	stlhash "github.com/kkkunny/stl/hash"
)

type _ConcurrentMap[K comparable, V any] struct {
	data cmap.ConcurrentMap[K, V]
}

func _NewConcurrentMap[K comparable, V any]() HashMap[K, V] {
	hasher := stlhash.GetMapHashFunc[K]()
	return &_ConcurrentMap[K, V]{
		data: cmap.NewWithCustomShardingFunction[K, V](func(k K) uint32 {
			return uint32(hasher(k))
		}),
	}
}

func _NewConcurrentMapWith[K comparable, V any](vs ...any) HashMap[K, V] {
	self := _NewConcurrentMap[K, V]()
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

func (self *_ConcurrentMap[K, V]) Clone() stlmaps.MapObj[K, V] {
	newMap := _NewConcurrentMap[K, V]()
	self.data.IterCb(func(k K, v V) {
		newMap.Set(k, v)
	})
	return newMap
}

func (self *_ConcurrentMap[K, V]) Equal(dstObj stlmaps.MapObj[K, V]) (eq bool) {
	if dstObj == nil && self == nil {
		return true
	} else if dstObj == nil {
		return false
	}

	dst, ok := dstObj.(HashMap[K, V])
	if !ok {
		return false
	}

	if self.Length() != dst.Length() {
		return false
	}

	for kv := range self.data.IterBuffered() {
		if !dst.Contain(kv.Key) || !stlbasic.Equal(kv.Val, dst.Get(kv.Key)) {
			return false
		}
	}
	return true
}

func (_ *_ConcurrentMap[K, V]) NewWithIterator(iter stliter.Iterator[tuple.Tuple2[K, V]]) any {
	self := _NewConcurrentMapWith[K, V]()
	for iter.Next() {
		item := iter.Value()
		self.Set(item.Unpack())
	}
	return self
}

func (self *_ConcurrentMap[K, V]) Iterator() stliter.Iterator[tuple.Tuple2[K, V]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_ConcurrentMap[K, V]) Iter2() iter.Seq2[K, V] {
	kvs := self.data.IterBuffered()
	return func(yield func(K, V) bool) {
		kv, ok := <-kvs
		if !ok {
			return
		}
		yield(kv.Key, kv.Val)
	}
}

func (self *_ConcurrentMap[K, V]) MarshalJSON() ([]byte, error) {
	return self.data.MarshalJSON()
}

func (self *_ConcurrentMap[K, V]) UnmarshalJSON(data []byte) error {
	return self.data.UnmarshalJSON(data)
}

func (self *_ConcurrentMap[K, V]) Length() uint {
	return uint(self.data.Count())
}

// Set 插入键值对
func (self *_ConcurrentMap[K, V]) Set(k K, v V) V {
	oldV, _ := self.data.Get(k)
	self.data.Set(k, v)
	return oldV
}

// Get 获取值
func (self *_ConcurrentMap[K, V]) Get(k K, defaultValue ...V) V {
	v, ok := self.data.Get(k)
	if !ok {
		return stlslices.Last(defaultValue)
	}
	return v
}

// Contain 是否包含键
func (self *_ConcurrentMap[K, V]) Contain(k K) bool {
	return self.data.Has(k)
}

// Remove 移除键值对
func (self *_ConcurrentMap[K, V]) Remove(k K, defaultValue ...V) V {
	v, ok := self.data.Get(k)
	if !ok {
		return stlslices.Last(defaultValue)
	}
	self.data.Remove(k)
	return v
}

// Clear 清空
func (self *_ConcurrentMap[K, V]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_ConcurrentMap[K, V]) Empty() bool {
	return self.data.IsEmpty()
}

// Keys 获取所有键
func (self *_ConcurrentMap[K, V]) Keys() []K {
	return self.data.Keys()
}

// Values 获取所有值
func (self *_ConcurrentMap[K, V]) Values() []V {
	vs := make([]V, self.Length())
	var i int
	self.data.IterCb(func(_ K, v V) {
		vs[i] = v
		i++
	})
	return vs
}

// KeyValues 获取所有键值对
func (self *_ConcurrentMap[K, V]) KeyValues() []tuple.Tuple2[K, V] {
	kvs := make([]tuple.Tuple2[K, V], self.Length())
	var i int
	self.data.IterCb(func(k K, v V) {
		kvs[i] = tuple.Pack2(k, v)
		i++
	})
	return kvs
}

func (self *_ConcurrentMap[K, V]) String() string {
	var buf strings.Builder
	var i int
	buf.WriteString("HashMap{")
	self.data.IterCb(func(k K, v V) {
		buf.WriteString(fmt.Sprintf("%v", k))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < int(self.data.Count())-1 {
			buf.WriteString(", ")
		}
		i++
	})
	buf.WriteByte('}')
	return buf.String()
}
