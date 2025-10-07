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

type _StdBiMap[T, E comparable] struct {
	keys   hashmap.HashMap[T, E]
	values hashmap.HashMap[E, T]
}

func _NewStdBiMap[T, E comparable]() BiMap[T, E] {
	return &_StdBiMap[T, E]{
		keys:   hashmap.AnyWith[T, E](),
		values: hashmap.AnyWith[E, T](),
	}
}

func _NewStdBiMapWithCapacity[T, E comparable](cap uint) BiMap[T, E] {
	return &_StdBiMap[T, E]{
		keys:   hashmap.AnyWithCap[T, E](cap),
		values: hashmap.AnyWithCap[E, T](cap),
	}
}

func _NewStdBiMapWith[T, E comparable](vs ...any) BiMap[T, E] {
	self := _NewStdBiMapWithCapacity[T, E](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Put(vs[i].(T), vs[i+1].(E))
	}
	return self
}

func (self *_StdBiMap[T, E]) Clone() stlmaps.MapObj[T, E] {
	return &_StdBiMap[T, E]{
		keys:   clone.Clone(self.keys),
		values: clone.Clone(self.values),
	}
}

func (self *_StdBiMap[T, E]) Equal(dstObj stlmaps.MapObj[T, E]) bool {
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

func (_ *_StdBiMap[T, E]) NewWithIterator(iter stliter.Iterator[tuple.Tuple2[T, E]]) any {
	self := _NewStdBiMapWithCapacity[T, E](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Put(item.Unpack())
	}
	return self
}

func (self *_StdBiMap[T, E]) Iterator() stliter.Iterator[tuple.Tuple2[T, E]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_StdBiMap[T, E]) Iter2() iter.Seq2[T, E] {
	return self.keys.Iter2()
}

func (self *_StdBiMap[T, E]) Length() uint {
	return self.keys.Length()
}

// Set 插入键值对
func (self *_StdBiMap[T, E]) Set(k T, v E) E {
	_, pv := self.Put(k, v)
	return pv
}

// Put 插入键值对
func (self *_StdBiMap[T, E]) Put(k T, v E) (T, E) {
	pv, pk := self.RemoveKey(k), self.RemoveValue(v)
	self.keys.Set(k, v)
	self.values.Set(v, k)
	return pk, pv
}

// Get 获取值
func (self *_StdBiMap[T, E]) Get(k T, defaultValue ...E) E {
	return self.GetValue(k, defaultValue...)
}

// GetValue 获取值
func (self *_StdBiMap[T, E]) GetValue(k T, defaultValue ...E) E {
	return self.keys.Get(k, defaultValue...)
}

// GetKey 获取键
func (self *_StdBiMap[T, E]) GetKey(v E, defaultKey ...T) T {
	return self.values.Get(v, defaultKey...)
}

// Contain 是否包含键
func (self *_StdBiMap[T, E]) Contain(k T) bool {
	return self.ContainKey(k)
}

// ContainKey 是否包含键
func (self *_StdBiMap[T, E]) ContainKey(k T) bool {
	return self.keys.Contain(k)
}

// ContainValue 是否包含值
func (self *_StdBiMap[T, E]) ContainValue(v E) bool {
	return self.values.Contain(v)
}

// Remove 移除键
func (self *_StdBiMap[T, E]) Remove(k T, defaultValue ...E) E {
	return self.RemoveKey(k, defaultValue...)
}

// RemoveKey 移除键
func (self *_StdBiMap[T, E]) RemoveKey(k T, defaultValue ...E) E {
	exist := self.Contain(k)
	pv := self.keys.Remove(k, defaultValue...)
	if exist {
		self.values.Remove(pv)
	}
	return pv
}

// RemoveValue 移除值
func (self *_StdBiMap[T, E]) RemoveValue(v E, defaultKey ...T) T {
	exist := self.ContainValue(v)
	pk := self.values.Remove(v, defaultKey...)
	if exist {
		self.keys.Remove(pk)
	}
	return pk
}

// Clear 清空
func (self *_StdBiMap[T, E]) Clear() {
	self.keys.Clear()
	self.values.Clear()
}

// Empty 是否为空
func (self *_StdBiMap[T, E]) Empty() bool {
	return self.keys.Empty()
}

// Keys 获取所有键
func (self *_StdBiMap[T, E]) Keys() []T {
	return self.keys.Keys()
}

// Values 获取所有值
func (self *_StdBiMap[T, E]) Values() []E {
	return self.values.Keys()
}

// KeyValues 获取所有键值对
func (self *_StdBiMap[T, E]) KeyValues() []tuple.Tuple2[T, E] {
	return self.keys.KeyValues()
}

func (self *_StdBiMap[T, E]) String() string {
	var buf strings.Builder
	buf.WriteString("BiMap{")
	for i, p := range self.KeyValues() {
		k, v := p.Unpack()
		buf.WriteString(fmt.Sprintf("%v", k))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < int(self.Length())-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_StdBiMap[T, E]) getKeyData() hashmap.HashMap[T, E] {
	return self.keys
}

func (self *_StdBiMap[T, E]) getValueData() hashmap.HashMap[E, T] {
	return self.values
}

func (self *_StdBiMap[K, V]) MarshalJSON() ([]byte, error) {
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

func (self *_StdBiMap[K, V]) UnmarshalJSON(data []byte) error {
	for kvs, err := range json2.UnmarshalToMapObj[K, V](bytes.NewReader(data)) {
		if err != nil {
			return err
		}
		self.Set(kvs.Unpack())
	}
	return nil
}
