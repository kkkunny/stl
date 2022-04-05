package table

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/kkkunny/stl/list"
	. "github.com/kkkunny/stl/types"
)

type hashMapEntry[K Hasher, V any] struct {
	Entry[K, V]
	next *hashMapEntry[K, V]
}

// 哈希表
type HashMap[K Hasher, V any] struct {
	length  Usize
	buckets []*hashMapEntry[K, V]
}

// 新建哈希表
func NewHashMap[K Hasher, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		length:  0,
		buckets: make([]*hashMapEntry[K, V], 16),
	}
}

// 转成字符串 O(N)
func (self *HashMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteByte('{')
	var index Usize
	for iter := self.Iterator(); iter.HasValue(); iter.Next() {
		buf.WriteString(fmt.Sprintf("%v: %v", iter.Key(), iter.Value()))
		if index < self.length-1 {
			buf.WriteString(", ")
		}
		index++
	}
	buf.WriteByte('}')
	return buf.String()
}

// 获取长度 O(1)
func (self *HashMap[K, V]) Length() Usize {
	return self.length
}

// 是否为空 O(1)
func (self *HashMap[K, V]) Empty() bool {
	return self.length == 0
}

// 哈希
func (self *HashMap[K, V]) getHash(k K) int32 {
	hash := k.Hash()
	hash ^= (hash >> 20) ^ (hash >> 12)
	return hash ^ (hash >> 7) ^ (hash >> 4)
}

// 获取哈希所在索引
func (self *HashMap[K, V]) getIndexFromHash(k K) Usize {
	hash := self.getHash(k)
	return Usize(int(hash) % cap(self.buckets))
}

// 检查扩容
func (self *HashMap[K, V]) checkExpand() {
	if float64(self.length)/float64(cap(self.buckets)) > 0.75 {
		self.expandSize()
	}
}

// 扩容 O(N)
func (self *HashMap[K, V]) expandSize() {
	newSize := cap(self.buckets) * 2
	oldBuckets := self.buckets
	self.buckets = make([]*hashMapEntry[K, V], newSize)
	self.length = 0
	for _, head := range oldBuckets {
		for cursor := head; cursor != nil; cursor = cursor.next {
			self.Set(cursor.Key, cursor.Value)
		}
	}
}

// 设置键值对 O(1)-O(N)
func (self *HashMap[K, V]) Set(k K, v V) {
	index := self.getIndexFromHash(k)
	head := self.buckets[index]
	if head == nil {
		self.buckets[index] = &hashMapEntry[K, V]{
			Entry: Entry[K, V]{
				Key:   k,
				Value: v,
			},
		}
	} else {
		var prev *hashMapEntry[K, V]
		for ; head != nil; head = head.next {
			if head.Key.Hash() == k.Hash() {
				head.Value = v
				return
			}
			prev = head
		}
		prev.next = &hashMapEntry[K, V]{
			Entry: Entry[K, V]{
				Key:   k,
				Value: v,
			},
		}
	}
	self.length++
	self.checkExpand()
}

// 获取值 O(1)
func (self *HashMap[K, V]) Get(k K, d V) V {
	index := self.getIndexFromHash(k)
	head := self.buckets[index]
	if head != nil {
		for ; head != nil; head = head.next {
			if head.Key.Hash() == k.Hash() {
				return head.Value
			}
		}
	}
	return d
}

// 移除键值对 O(1)
func (self *HashMap[K, V]) Remove(k K, d V) V {
	index := self.getIndexFromHash(k)
	head := self.buckets[index]
	if head != nil {
		var prev *hashMapEntry[K, V]
		for ; head != nil; head = head.next {
			if head.Key.Hash() == k.Hash() {
				break
			}
			prev = head
		}
		if head != nil {
			elem := head.Value
			if prev == nil {
				self.buckets[index] = head.next
			} else {
				prev.next = head.next
			}
			self.length--
			return elem
		}
	}
	return d
}

// 是否存在键 O(1)
func (self *HashMap[K, V]) ContainKey(k K) bool {
	index := self.getIndexFromHash(k)
	head := self.buckets[index]
	if head != nil {
		for ; head != nil; head = head.next {
			if head.Key.Hash() == k.Hash() {
				return true
			}
		}
	}
	return false
}

// 获取键 O(N)
func (self *HashMap[K, V]) Keys() *list.ArrayList[K] {
	var index Usize
	keys := list.NewArrayList[K](self.length, self.length)
	for _, head := range self.buckets {
		for ; head != nil; head = head.next {
			keys.Set(index, head.Key)
			index++
		}
	}
	return keys
}

// 获取值 O(N)
func (self *HashMap[K, V]) Values() *list.ArrayList[V] {
	var index Usize
	values := list.NewArrayList[V](self.length, self.length)
	for _, head := range self.buckets {
		for ; head != nil; head = head.next {
			values.Set(index, head.Value)
			index++
		}
	}
	return values
}

// 清空 O(1)
func (self *HashMap[K, V]) Clear() {
	self.length = 0
	self.buckets = make([]*hashMapEntry[K, V], cap(self.buckets))
}

// 克隆 O(N)
func (self *HashMap[K, V]) Clone() *HashMap[K, V] {
	cpy := NewHashMap[K, V]()
	for _, head := range self.buckets {
		for ; head != nil; head = head.next {
			cpy.Set(head.Key, head.Value)
		}
	}
	return cpy
}

// 过滤 O(N)
func (self *HashMap[K, V]) Filter(f func(k K, v V) bool) *HashMap[K, V] {
	hm := NewHashMap[K, V]()
	for _, head := range self.buckets {
		for ; head != nil; head = head.next {
			if f(head.Key, head.Value) {
				hm.Set(head.Key, head.Value)
			}
		}
	}
	return hm
}

// 获取迭代器
func (self *HashMap[K, V]) Iterator() *HashMapIterator[K, V] {
	data := make([]*hashMapEntry[K, V], self.length)
	var index Usize
	for _, head := range self.buckets {
		for ; head != nil; head = head.next {
			data[index] = head
			index++
		}
	}
	iter := &HashMapIterator[K, V]{data: data}
	iter.shuffle()
	return iter
}

// 迭代器
type HashMapIterator[K Hasher, V any] struct {
	data  []*hashMapEntry[K, V]
	index Usize
}

// 是否存在值
func (self *HashMapIterator[K, V]) HasValue() bool {
	return 0 <= self.index && self.index < Usize(len(self.data))
}

// 下一个
func (self *HashMapIterator[K, V]) Next() {
	self.index++
}

// 获取键
func (self *HashMapIterator[K, V]) Key() K {
	return self.data[self.index].Key
}

// 获取值
func (self *HashMapIterator[K, V]) Value() V {
	return self.data[self.index].Value
}

func (self *HashMapIterator[K, V]) shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < len(self.data); i++ {
		randIndex := r.Intn(len(self.data))
		self.data[i], self.data[randIndex] = self.data[randIndex], self.data[i]
	}
}
