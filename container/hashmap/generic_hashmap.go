package hashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	hm "github.com/zyedidia/generic/hashmap"

	stlcmp "github.com/kkkunny/stl/cmp"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
	stlslices "github.com/kkkunny/stl/container/slices"
	stlhash "github.com/kkkunny/stl/hash"
)

const initGenericHashMapCapacity = 10

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
	return uint(self.data.Size())
}

func (self *_GenericHashMap[K, V]) Clone() HashMap[K, V] {
	return &_GenericHashMap[K, V]{data: self.data.Copy()}
}

func (self *_GenericHashMap[K, V]) Equal(dst HashMap[K, V]) (eq bool) {
	if self.Length() != dst.Length() {
		return false
	}

	for _, p := range self.KeyValues() {
		if !dst.ContainKey(p.First) || !stlcmp.Equal(p.Second, dst.Get(p.First)) {
			return false
		}
	}
	return true
}

func (_ *_GenericHashMap[K, V]) NewWithIterator(iter stliter.Iterator[pair.Pair[K, V]]) any {
	self := _NewGenericHashMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self *_GenericHashMap[K, V]) Iterator() stliter.Iterator[pair.Pair[K, V]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_GenericHashMap[K, V]) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	err := buf.WriteByte('{')
	if err != nil {
		return nil, err
	}
	for i, p := range self.KeyValues() {
		_, err = buf.WriteString(fmt.Sprintf("\"%+v\"", p.First))
		if err != nil {
			return nil, err
		}
		err = buf.WriteByte(':')
		if err != nil {
			return nil, err
		}
		vs, err := json.Marshal(p.Second)
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

// ContainKey 是否包含键
func (self *_GenericHashMap[K, V]) ContainKey(k K) bool {
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
func (self *_GenericHashMap[K, V]) KeyValues() []pair.Pair[K, V] {
	kvs := make([]pair.Pair[K, V], self.Length())
	var i int
	self.data.Each(func(k K, v V) {
		kvs[i] = pair.NewPair(k, v)
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
