package linkedhashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kkkunny/stl/container/hashmap"
	stliter "github.com/kkkunny/stl/container/iter"
	stlmaps "github.com/kkkunny/stl/container/maps"
	"github.com/kkkunny/stl/container/tuple"
	"github.com/kkkunny/stl/list"
)

type _AnyLinkedHashMap[K, V any] struct {
	kvs  hashmap.HashMap[K, *list.Element[tuple.Tuple2[K, V]]]
	list *list.List[tuple.Tuple2[K, V]]
}

func _NewAnyLinkedHashMap[K, V any]() LinkedHashMap[K, V] {
	return &_AnyLinkedHashMap[K, V]{
		kvs:  hashmap.AnyWith[K, *list.Element[tuple.Tuple2[K, V]]](),
		list: list.New[tuple.Tuple2[K, V]](),
	}
}

func _NewAnyLinkedHashMapWithCapacity[K, V any](cap uint) LinkedHashMap[K, V] {
	return &_AnyLinkedHashMap[K, V]{
		kvs:  hashmap.AnyWithCap[K, *list.Element[tuple.Tuple2[K, V]]](cap),
		list: list.New[tuple.Tuple2[K, V]](),
	}
}

func _NewAnyLinkedHashMapWith[K, V any](vs ...any) LinkedHashMap[K, V] {
	self := _NewAnyLinkedHashMapWithCapacity[K, V](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

func (self *_AnyLinkedHashMap[K, V]) Capacity() uint {
	return self.kvs.Capacity()
}

func (self *_AnyLinkedHashMap[K, V]) Clone() stlmaps.MapObj[K, V] {
	hm := _NewAnyLinkedHashMapWithCapacity[K, V](self.Capacity())
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		hm.Set(cursor.Value.Unpack())
	}
	return hm
}

func (self *_AnyLinkedHashMap[K, V]) Equal(dstObj stlmaps.MapObj[K, V]) bool {
	if dstObj == nil && self == nil {
		return true
	} else if dstObj == nil {
		return false
	}

	dst, ok := dstObj.(LinkedHashMap[K, V])
	if !ok {
		return false
	}

	if self.Length() != dst.Length() {
		return false
	}

	for c1, c2 := self.list.Front(), dst.getList().Front(); c1 != nil && c2 != nil; c1, c2 = c1.Next(), c2.Next() {
		if !c1.Value.Equal(c2.Value) {
			return false
		}
	}
	return true
}

func (_ *_AnyLinkedHashMap[K, V]) NewWithIterator(iter stliter.Iterator[tuple.Tuple2[K, V]]) any {
	self := _NewAnyLinkedHashMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		self.Set(iter.Value().Unpack())
	}
	return self
}

func (self *_AnyLinkedHashMap[K, V]) Iterator() stliter.Iterator[tuple.Tuple2[K, V]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_AnyLinkedHashMap[K, V]) MarshalJSON() ([]byte, error) {
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
		if i < int(self.Length())-1 {
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

func (self *_AnyLinkedHashMap[K, V]) Length() uint {
	return self.kvs.Length()
}

// Set 插入键值对
func (self *_AnyLinkedHashMap[K, V]) Set(k K, v V) V {
	if node := self.kvs.Get(k); node != nil {
		pv := node.Value.E2()
		node.Value = tuple.Pack2[K, V](k, v)
		self.list.MoveToBack(node)
		return pv
	}
	self.kvs.Set(k, self.list.PushBack(tuple.Pack2[K, V](k, v)))
	var pv V
	return pv
}

// Get 获取值
func (self *_AnyLinkedHashMap[K, V]) Get(k K, defaultValue ...V) V {
	node := self.kvs.Get(k)
	if node == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if node == nil {
		var v V
		return v
	}
	return node.Value.E2()
}

// Contain 是否包含键
func (self *_AnyLinkedHashMap[K, V]) Contain(k K) bool {
	return self.kvs.Contain(k)
}

// Remove 移除键值对
func (self *_AnyLinkedHashMap[K, V]) Remove(k K, defaultValue ...V) V {
	node := self.kvs.Remove(k)
	if node == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if node == nil {
		var v V
		return v
	}
	return self.list.Remove(node).E2()
}

// Clear 清空
func (self *_AnyLinkedHashMap[K, V]) Clear() {
	self.list = list.New[tuple.Tuple2[K, V]]()
	self.kvs.Clear()
}

// Empty 是否为空
func (self *_AnyLinkedHashMap[K, V]) Empty() bool {
	return self.list.Len() == 0
}

// Keys 获取所有键
func (self *_AnyLinkedHashMap[K, V]) Keys() []K {
	keys := make([]K, self.Length())
	var i uint
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		keys[i] = cursor.Value.E1()
		i++
	}
	return keys
}

// Values 获取所有值
func (self *_AnyLinkedHashMap[K, V]) Values() []V {
	values := make([]V, self.Length())
	var i uint
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		values[i] = cursor.Value.E2()
		i++
	}
	return values
}

// KeyValues 获取所有键值对
func (self *_AnyLinkedHashMap[K, V]) KeyValues() []tuple.Tuple2[K, V] {
	pairs := make([]tuple.Tuple2[K, V], self.Length())
	var i uint
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		pairs[i] = tuple.Pack2(cursor.Value.Unpack())
		i++
	}
	return pairs
}

func (self *_AnyLinkedHashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteString("LinkedHashMap{")
	var i int
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		k, v := cursor.Value.Unpack()
		buf.WriteString(fmt.Sprintf("%v", k))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", v))
		if cursor.Next() != nil {
			buf.WriteString(", ")
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_AnyLinkedHashMap[K, V]) getList() *list.List[tuple.Tuple2[K, V]] {
	return self.list
}
