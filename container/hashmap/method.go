package hashmap

import (
	stlbasic "github.com/kkkunny/stl/basic"
	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/pair"
)

const (
	initialBucketCapacity uint    = 10  // 初始桶容量
	expandLoadFactor      float64 = 0.8 // 装载因子上限
	expandMultiple        float64 = 2   // 扩容倍数
)

// 初始化
func (self *HashMap[K, V]) init() {
	if self.buckets != nil {
		return
	}
	buckets := make([]bucket[K, V], initialBucketCapacity)
	self.buckets = &buckets
}

// 扩容
func (self *HashMap[K, V]) expand() {
	loadFactor := float64(self.Length()) / float64(self.Capacity())
	if loadFactor < expandLoadFactor {
		return
	}
	newSelf := NewHashMapWithCapacity[K, V](uint(float64(self.Capacity()) * expandMultiple))
	for _, bkt := range *self.buckets {
		// TODO: 优化
		newSelf.Set(bkt.Key, bkt.Value)
	}
	*self = newSelf
}

// 获取距离
func (self *HashMap[K, V]) getDistance(hash uint64, curIndex uint) uint {
	capacity := self.Capacity()
	targetIndex := uint(hash % uint64(capacity))
	if curIndex < targetIndex {
		return capacity - targetIndex + curIndex
	} else {
		return curIndex - targetIndex
	}
}

// 寻找桶
func (self *HashMap[K, V]) findBucket(hash uint64, k K) *bucket[K, V] {
	index := uint(hash % uint64(self.Capacity()))
	for {
		bkt := &(*self.buckets)[index]
		if !bkt.Used {
			return nil
		} else if stlbasic.Equal(k, bkt.Key) {
			return bkt
		} else if self.getDistance(hash, index) > self.getDistance(bkt.Hash, index) {
			return nil
		}
		if index == self.Capacity()-1 {
			index = 0
		} else {
			index++
		}
	}
}

// 插入桶（不更改length）
func (self *HashMap[K, V]) insertBucket(newBKT bucket[K, V]) (bool, V) {
	index := uint(newBKT.Hash % uint64(self.Capacity()))
	for {
		bkt := &(*self.buckets)[index]
		if !bkt.Used {
			*bkt = newBKT
			var v V
			return false, v
		} else if stlbasic.Equal(newBKT.Key, bkt.Key) {
			prev := bkt.Value
			bkt.Value = newBKT.Value
			return true, prev
		} else if self.getDistance(newBKT.Hash, index) > self.getDistance(bkt.Hash, index) {
			prevBKT := *bkt
			*bkt = newBKT
			self.insertBucket(prevBKT)
			var v V
			return false, v
		}
		if index == self.Capacity()-1 {
			index = 0
		} else {
			index++
		}
	}
}

// Set 插入键值对
func (self *HashMap[K, V]) Set(k K, v V) V {
	self.init()
	self.expand()

	used, prevValue := self.insertBucket(bucket[K, V]{
		Used:  true,
		Hash:  stlbasic.Hash(k),
		Key:   k,
		Value: v,
	})
	if !used {
		self.length++
	}
	return prevValue
}

// Get 获取值
func (self HashMap[K, V]) Get(k K, defaultValue ...V) V {
	self.init()
	bkt := self.findBucket(stlbasic.Hash(k), k)
	if bkt == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if bkt == nil {
		var v V
		return v
	}
	return bkt.Value
}

// ContainKey 是否包含键
func (self HashMap[K, V]) ContainKey(k K) bool {
	self.init()
	return self.findBucket(stlbasic.Hash(k), k) != nil
}

// Remove 移除键值对
func (self *HashMap[K, V]) Remove(k K, defaultValue ...V) V {
	self.init()

	bkt := self.findBucket(stlbasic.Hash(k), k)
	if bkt == nil && len(defaultValue) > 0 {
		return defaultValue[0]
	} else if bkt == nil {
		var v V
		return v
	}

	prevValue := bkt.Value
	self.length--

	index := uint(bkt.Hash % uint64(self.Capacity()))
	for {
		bkt = &(*self.buckets)[index]
		var nextBKT *bucket[K, V]
		if index+1 < self.Capacity() {
			nextBKT = &(*self.buckets)[index+1]
		} else {
			nextBKT = &(*self.buckets)[0]
		}

		if !nextBKT.Used || self.getDistance(nextBKT.Hash, index) == 0 {
			bkt.Used = false
			break
		} else {
			*bkt = *nextBKT
		}
		if index == self.Capacity()-1 {
			index = 0
		} else {
			index++
		}
	}
	return prevValue
}

// Clear 清空
func (self *HashMap[K, V]) Clear() {
	self.buckets = nil
	self.length = 0
	self.init()
}

// Empty 是否为空
func (self HashMap[K, V]) Empty() bool {
	return self.length == 0
}

// Keys 获取所有键
func (self HashMap[K, V]) Keys() dynarray.DynArray[K] {
	self.init()

	keys := dynarray.NewDynArrayWithLength[K](self.Length())
	var i uint
	for _, bkt := range *self.buckets {
		if !bkt.Used {
			continue
		}
		keys.Set(i, bkt.Key)
		i++
	}
	keys.Shuffle()
	return keys
}

// Values 获取所有值
func (self HashMap[K, V]) Values() dynarray.DynArray[V] {
	self.init()

	values := dynarray.NewDynArrayWithLength[V](self.Length())
	var i uint
	for _, bkt := range *self.buckets {
		if !bkt.Used {
			continue
		}
		values.Set(i, bkt.Value)
		i++
	}
	values.Shuffle()
	return values
}

// KeyValues 获取所有键值对
func (self HashMap[K, V]) KeyValues() dynarray.DynArray[pair.Pair[K, V]] {
	self.init()

	pairs := dynarray.NewDynArrayWithLength[pair.Pair[K, V]](self.Length())
	var i uint
	for _, bkt := range *self.buckets {
		if !bkt.Used {
			continue
		}
		pairs.Set(i, pair.NewPair(bkt.Key, bkt.Value))
		i++
	}
	pairs.Shuffle()
	return pairs
}
