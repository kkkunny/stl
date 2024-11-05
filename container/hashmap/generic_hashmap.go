package hashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"

	hm "github.com/zyedidia/generic/hashmap"

	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	stlmaps "github.com/kkkunny/stl/container/maps"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/container/tuple"
	stlhash "github.com/kkkunny/stl/hash"
)

const initGenericHashMapCapacity = 8

type genericHashMapDataEntry[K, V any] struct {
	key    K
	filled bool
	value  V
}

type genericHashMapDataOps[T any] struct {
	equals func(a, b T) bool
	hash   func(t T) uint64
}

type genericHashMapData[K, V any] struct {
	entries  []genericHashMapDataEntry[K, V]
	capacity uint64
	length   uint64
	readonly bool

	ops genericHashMapDataOps[K]
}

func toGenericHashMapData[K, V any](h *hm.Map[K, V]) *genericHashMapData[K, V] {
	if h == nil {
		return nil
	}
	return (*genericHashMapData[K, V])(unsafe.Pointer(h))
}

func (self *genericHashMapData[K, V]) toData() *hm.Map[K, V] {
	if self == nil {
		return nil
	}
	return (*hm.Map[K, V])(unsafe.Pointer(self))
}

func (self *genericHashMapData[K, V]) Each(fn func(k K, v V) bool) {
	for _, ent := range self.entries {
		if ent.filled {
			if !fn(ent.key, ent.value) {
				return
			}
		}
	}
}

type _GenericHashMap[K, V any] struct {
	data *hm.Map[K, V]
}

func _NewGenericHashMap[K, V any]() HashMap[K, V] {
	return _NewGenericHashMapWithCapacity[K, V](initGenericHashMapCapacity)
}

func _NewGenericHashMapWithCapacity[K, V any](cap uint) HashMap[K, V] {
	return &_GenericHashMap[K, V]{data: hm.New[K, V](uint64(cap), stlcmp.GetEqualFunc[K](), stlhash.GetHashFunc[K]())}
}

func _NewGenericHashMapWith[K, V any](vs ...any) HashMap[K, V] {
	self := _NewGenericHashMapWithCapacity[K, V](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

func (self *_GenericHashMap[K, V]) Capacity() uint {
	h := toGenericHashMapData(self.data)
	return uint(h.capacity)
}

func (self *_GenericHashMap[K, V]) Clone() stlmaps.MapObj[K, V] {
	return &_GenericHashMap[K, V]{data: self.data.Copy()}
}

func (self *_GenericHashMap[K, V]) Equal(dstObj stlmaps.MapObj[K, V]) (eq bool) {
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

	for _, p := range self.KeyValues() {
		k, v := p.Unpack()
		if !dst.Contain(k) || !stlcmp.Equal(v, dst.Get(k)) {
			return false
		}
	}
	return true
}

func (_ *_GenericHashMap[K, V]) NewWithIterator(iter stliter.Iterator[tuple.Tuple2[K, V]]) any {
	self := _NewGenericHashMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		self.Set(iter.Value().Unpack())
	}
	return self
}

func (self *_GenericHashMap[K, V]) Iterator() stliter.Iterator[tuple.Tuple2[K, V]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_GenericHashMap[K, V]) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	err := buf.WriteByte('{')
	if err != nil {
		return nil, err
	}
	for i, p := range self.KeyValues() {
		k, v := p.Unpack()
		_, err = buf.WriteString(fmt.Sprintf("\"%+v\"", k))
		if err != nil {
			return nil, err
		}
		err = buf.WriteByte(':')
		if err != nil {
			return nil, err
		}
		vs, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(vs)
		if err != nil {
			return nil, err
		}
		if i < self.data.Size()-1 {
			err = buf.WriteByte(',')
			if err != nil {
				return nil, err
			}
		}
	}
	err = buf.WriteByte('}')
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (self *_GenericHashMap[K, V]) Length() uint {
	return uint(self.data.Size())
}

// Set 插入键值对
func (self *_GenericHashMap[K, V]) Set(k K, v V) V {
	oldV, _ := self.data.Get(k)
	self.data.Put(k, v)
	return oldV
}

// Get 获取值
func (self *_GenericHashMap[K, V]) Get(k K, defaultValue ...V) V {
	v, ok := self.data.Get(k)
	if !ok {
		return stlslices.Last(defaultValue)
	}
	return v
}

// Contain 是否包含键
func (self *_GenericHashMap[K, V]) Contain(k K) bool {
	_, ok := self.data.Get(k)
	return ok
}

// Remove 移除键值对
func (self *_GenericHashMap[K, V]) Remove(k K, defaultValue ...V) V {
	v, ok := self.data.Get(k)
	if !ok {
		return stlslices.Last(defaultValue)
	}
	self.data.Remove(k)
	return v
}

// Clear 清空
func (self *_GenericHashMap[K, V]) Clear() {
	self.data.Clear()
}

// Empty 是否为空
func (self *_GenericHashMap[K, V]) Empty() bool {
	return self.data.Size() == 0
}

// Keys 获取所有键
func (self *_GenericHashMap[K, V]) Keys() []K {
	ks := make([]K, self.Length())
	var i int
	self.data.Each(func(k K, _ V) {
		ks[i] = k
		i++
	})
	return ks
}

// Values 获取所有值
func (self *_GenericHashMap[K, V]) Values() []V {
	vs := make([]V, self.Length())
	var i int
	self.data.Each(func(_ K, v V) {
		vs[i] = v
		i++
	})
	return vs
}

// KeyValues 获取所有键值对
func (self *_GenericHashMap[K, V]) KeyValues() []tuple.Tuple2[K, V] {
	kvs := make([]tuple.Tuple2[K, V], self.Length())
	var i int
	self.data.Each(func(k K, v V) {
		kvs[i] = tuple.Pack2(k, v)
		i++
	})
	return kvs
}

func (self *_GenericHashMap[K, V]) String() string {
	var buf strings.Builder
	var i int
	buf.WriteString("HashMap{")
	self.data.Each(func(k K, v V) {
		buf.WriteString(fmt.Sprintf("%v", k))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < self.data.Size()-1 {
			buf.WriteString(", ")
		}
		i++
	})
	buf.WriteByte('}')
	return buf.String()
}
