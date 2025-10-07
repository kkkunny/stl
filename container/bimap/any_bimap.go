package bimap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"iter"
	"strings"

	"github.com/kkkunny/stl/clone"
	"github.com/kkkunny/stl/container/hashmap"
	stliter "github.com/kkkunny/stl/container/iter"
	stlmaps "github.com/kkkunny/stl/container/maps"
	"github.com/kkkunny/stl/container/tuple"
	json2 "github.com/kkkunny/stl/internal/json"
)

type _AnyBiMap[T, E any] struct {
	keys   hashmap.HashMap[T, E]
	values hashmap.HashMap[E, T]
}

func _NewAnyBiMap[T, E any]() BiMap[T, E] {
	return &_AnyBiMap[T, E]{
		keys:   hashmap.AnyWith[T, E](),
		values: hashmap.AnyWith[E, T](),
	}
}

func _NewAnyBiMapWithCapacity[T, E any](cap uint) BiMap[T, E] {
	return &_AnyBiMap[T, E]{
		keys:   hashmap.AnyWithCap[T, E](cap),
		values: hashmap.AnyWithCap[E, T](cap),
	}
}

func _NewAnyBiMapWith[T, E any](vs ...any) BiMap[T, E] {
	self := _NewAnyBiMapWithCapacity[T, E](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Put(vs[i].(T), vs[i+1].(E))
	}
	return self
}

func (self *_AnyBiMap[T, E]) Clone() stlmaps.MapObj[T, E] {
	return &_AnyBiMap[T, E]{
		keys:   clone.Clone(self.keys),
		values: clone.Clone(self.values),
	}
}

func (self *_AnyBiMap[T, E]) Equal(dstObj stlmaps.MapObj[T, E]) bool {
	if dstObj == nil && self == nil {
		return true
	} else if dstObj == nil {
		return false
	}

	dst, ok := dstObj.(BiMap[T, E])
	if !ok {
		return false
	}

	return self.keys.Equal(dst.getKeyData())
}

func (_ *_AnyBiMap[T, E]) NewWithIterator(iter stliter.Iterator[tuple.Tuple2[T, E]]) any {
	self := _NewAnyBiMapWithCapacity[T, E](iter.Length())
	for iter.Next() {
		self.Put(iter.Value().Unpack())
	}
	return self
}

func (self *_AnyBiMap[T, E]) Iterator() stliter.Iterator[tuple.Tuple2[T, E]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_AnyBiMap[T, E]) Iter2() iter.Seq2[T, E] {
	return self.keys.Iter2()
}

func (self *_AnyBiMap[T, E]) Length() uint {
	return self.keys.Length()
}

// Set 插入键值对
func (self *_AnyBiMap[T, E]) Set(k T, v E) E {
	_, pv := self.Put(k, v)
	return pv
}

// Put 插入键值对
func (self *_AnyBiMap[T, E]) Put(k T, v E) (T, E) {
	pv, pk := self.RemoveKey(k), self.RemoveValue(v)
	self.keys.Set(k, v)
	self.values.Set(v, k)
	return pk, pv
}

// Get 获取值
func (self *_AnyBiMap[T, E]) Get(k T, defaultValue ...E) E {
	return self.GetValue(k, defaultValue...)
}

// GetValue 获取值
func (self *_AnyBiMap[T, E]) GetValue(k T, defaultValue ...E) E {
	return self.keys.Get(k, defaultValue...)
}

// GetKey 获取键
func (self *_AnyBiMap[T, E]) GetKey(v E, defaultKey ...T) T {
	return self.values.Get(v, defaultKey...)
}

// Contain 是否包含键
func (self *_AnyBiMap[T, E]) Contain(k T) bool {
	return self.ContainKey(k)
}

// ContainKey 是否包含键
func (self *_AnyBiMap[T, E]) ContainKey(k T) bool {
	return self.keys.Contain(k)
}

// ContainValue 是否包含值
func (self *_AnyBiMap[T, E]) ContainValue(v E) bool {
	return self.values.Contain(v)
}

// Remove 移除键
func (self *_AnyBiMap[T, E]) Remove(k T, defaultValue ...E) E {
	return self.RemoveKey(k, defaultValue...)
}

// RemoveKey 移除键
func (self *_AnyBiMap[T, E]) RemoveKey(k T, defaultValue ...E) E {
	exist := self.Contain(k)
	pv := self.keys.Remove(k, defaultValue...)
	if exist {
		self.values.Remove(pv)
	}
	return pv
}

// RemoveValue 移除值
func (self *_AnyBiMap[T, E]) RemoveValue(v E, defaultKey ...T) T {
	exist := self.ContainValue(v)
	pk := self.values.Remove(v, defaultKey...)
	if exist {
		self.keys.Remove(pk)
	}
	return pk
}

// Clear 清空
func (self *_AnyBiMap[T, E]) Clear() {
	self.keys.Clear()
	self.values.Clear()
}

// Empty 是否为空
func (self *_AnyBiMap[T, E]) Empty() bool {
	return self.keys.Empty()
}

// Keys 获取所有键
func (self *_AnyBiMap[T, E]) Keys() []T {
	return self.keys.Keys()
}

// Values 获取所有值
func (self *_AnyBiMap[T, E]) Values() []E {
	return self.values.Keys()
}

// KeyValues 获取所有键值对
func (self *_AnyBiMap[T, E]) KeyValues() []tuple.Tuple2[T, E] {
	return self.keys.KeyValues()
}

func (self *_AnyBiMap[T, E]) String() string {
	var buf strings.Builder
	buf.WriteString("BiMap{")
	for i, p := range self.KeyValues() {
		buf.WriteString(fmt.Sprintf("%v", p.E1()))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", p.E2()))
		if i < int(self.Length())-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_AnyBiMap[T, E]) getKeyData() hashmap.HashMap[T, E] {
	return self.keys
}

func (self *_AnyBiMap[T, E]) getValueData() hashmap.HashMap[E, T] {
	return self.values
}

func (self *_AnyBiMap[K, V]) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	err := buf.WriteByte('{')
	if err != nil {
		return nil, err
	}
	for i, p := range self.KeyValues() {
		_, err = buf.WriteString(fmt.Sprintf("\"%+v\"", p.E1()))
		if err != nil {
			return nil, err
		}
		err = buf.WriteByte(':')
		if err != nil {
			return nil, err
		}
		vs, err := json.Marshal(p.E2())
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(vs)
		if err != nil {
			return nil, err
		}
		if i < int(self.keys.Length())-1 {
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

func (self *_AnyBiMap[K, V]) UnmarshalJSON(data []byte) error {
	for kvs, err := range json2.UnmarshalToMapObj[K, V](bytes.NewReader(data)) {
		if err != nil {
			return err
		}
		self.Set(kvs.Unpack())
	}
	return nil
}
