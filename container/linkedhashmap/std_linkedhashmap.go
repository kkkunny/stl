package linkedhashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	stlbasic "github.com/kkkunny/stl/cmp"
	"github.com/kkkunny/stl/container/hashmap"
	stliter "github.com/kkkunny/stl/container/iter"
	"github.com/kkkunny/stl/container/pair"
	"github.com/kkkunny/stl/list"
)

type _StdLinkedHashMap[K comparable, V any] struct {
	kvs  hashmap.HashMap[K, *list.Element[pair.Pair[K, V]]]
	list *list.List[pair.Pair[K, V]]
}

func _NewStdLinkedHashMap[K comparable, V any]() LinkedHashMap[K, V] {
	return &_StdLinkedHashMap[K, V]{
		kvs:  hashmap.AnyWith[K, *list.Element[pair.Pair[K, V]]](),
		list: list.New[pair.Pair[K, V]](),
	}
}

func _NewStdLinkedHashMapWithCapacity[K comparable, V any](cap uint) LinkedHashMap[K, V] {
	return &_StdLinkedHashMap[K, V]{
		kvs:  hashmap.AnyWithCap[K, *list.Element[pair.Pair[K, V]]](cap),
		list: list.New[pair.Pair[K, V]](),
	}
}

func _NewStdLinkedHashMapWith[K comparable, V any](vs ...any) LinkedHashMap[K, V] {
	self := _NewStdLinkedHashMapWithCapacity[K, V](uint(len(vs) / 2))
	for i := 0; i < len(vs); i += 2 {
		self.Set(vs[i].(K), vs[i+1].(V))
	}
	return self
}

func (self *_StdLinkedHashMap[K, V]) Capacity() uint {
	return self.kvs.Capacity()
}

func (self *_StdLinkedHashMap[K, V]) Clone() LinkedHashMap[K, V] {
	hm := _NewStdLinkedHashMapWithCapacity[K, V](self.Capacity())
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		hm.Set(cursor.Value.First, cursor.Value.Second)
	}
	return hm
}

func (self *_StdLinkedHashMap[K, V]) Equal(dst LinkedHashMap[K, V]) bool {
	if self.Length() != dst.Length() {
		return false
	}

	for c1, c2 := self.list.Front(), dst.getList().Front(); c1 != nil && c2 != nil; c1, c2 = c1.Next(), c2.Next() {
		v1, v2 := c1.Value, c2.Value
		if !stlbasic.Equal(v1.First, v2.First) || !stlbasic.Equal(v1.Second, v2.Second) {
			return false
		}
	}
	return true
}

func (_ *_StdLinkedHashMap[K, V]) NewWithIterator(iter stliter.Iterator[pair.Pair[K, V]]) any {
	self := _NewStdLinkedHashMapWithCapacity[K, V](iter.Length())
	for iter.Next() {
		item := iter.Value()
		self.Set(item.First, item.Second)
	}
	return self
}

func (self *_StdLinkedHashMap[K, V]) Iterator() stliter.Iterator[pair.Pair[K, V]] {
	return stliter.NewSliceIterator(self.KeyValues()...)
}

func (self *_StdLinkedHashMap[K, V]) MarshalJSON() ([]byte, error) {
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

func (self *_StdLinkedHashMap[K, V]) Length() uint {
	return self.kvs.Length()
}

// Set 插入键值对
func (self *_StdLinkedHashMap[K, V]) Set(k K, v V) V {
	if node := self.kvs.Get(k); node != nil {
		pv := node.Value.Second
		node.Value = pair.Pair[K, V]{First: k, Second: v}
		self.list.MoveToBack(node)
		return pv
	}
	self.kvs.Set(k, self.list.PushBack(pair.Pair[K, V]{First: k, Second: v}))
	var pv V
	return pv
}

// Get 获取值
func (self *_StdLinkedHashMap[K, V]) Get(k K, defaultValue ...V) V {
	node := self.kvs.Get(k)
	if node == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if node == nil {
		var v V
		return v
	}
	return node.Value.Second
}

// ContainKey 是否包含键
func (self *_StdLinkedHashMap[K, V]) ContainKey(k K) bool {
	return self.kvs.ContainKey(k)
}

// Remove 移除键值对
func (self *_StdLinkedHashMap[K, V]) Remove(k K, defaultValue ...V) V {
	node := self.kvs.Remove(k)
	if node == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if node == nil {
		var v V
		return v
	}
	return self.list.Remove(node).Second
}

// Clear 清空
func (self *_StdLinkedHashMap[K, V]) Clear() {
	self.list = list.New[pair.Pair[K, V]]()
	self.kvs.Clear()
}

// Empty 是否为空
func (self *_StdLinkedHashMap[K, V]) Empty() bool {
	return self.list.Len() == 0
}

// Keys 获取所有键
func (self *_StdLinkedHashMap[K, V]) Keys() []K {
	keys := make([]K, self.Length())
	var i uint
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		keys[i] = cursor.Value.First
		i++
	}
	return keys
}

// Values 获取所有值
func (self *_StdLinkedHashMap[K, V]) Values() []V {
	values := make([]V, self.Length())
	var i uint
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		values[i] = cursor.Value.Second
		i++
	}
	return values
}

// KeyValues 获取所有键值对
func (self *_StdLinkedHashMap[K, V]) KeyValues() []pair.Pair[K, V] {
	pairs := make([]pair.Pair[K, V], self.Length())
	var i uint
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		pairs[i] = pair.NewPair(cursor.Value.First, cursor.Value.Second)
		i++
	}
	return pairs
}

func (self *_StdLinkedHashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteString("LinkedHashMap{")
	var i int
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		buf.WriteString(fmt.Sprintf("%v", cursor.Value.First))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", cursor.Value.Second))
		if cursor.Next() != nil {
			buf.WriteString(", ")
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}

func (self *_StdLinkedHashMap[K, V]) getList() *list.List[pair.Pair[K, V]] {
	return self.list
}
