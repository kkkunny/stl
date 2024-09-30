package hashset

import (
	"github.com/kkkunny/stl/container/hashmap"
)

// Add 插入值
func (self *HashSet[T]) Add(v T) bool {
	exist := (*hashmap.HashMap[T, struct{}])(self).ContainKey(v)
	if exist {
		return false
	}
	(*hashmap.HashMap[T, struct{}])(self).Set(v, struct{}{})
	return true
}

// Contain 是否包含值
func (self HashSet[T]) Contain(v T) bool {
	return hashmap.HashMap[T, struct{}](self).ContainKey(v)
}

// Remove 移除值
func (self *HashSet[T]) Remove(v T) bool {
	exist := (*hashmap.HashMap[T, struct{}])(self).ContainKey(v)
	if !exist {
		return false
	}
	(*hashmap.HashMap[T, struct{}])(self).Remove(v)
	return true
}

// Clear 清空
func (self *HashSet[T]) Clear() {
	(*hashmap.HashMap[T, struct{}])(self).Clear()
}

// Empty 是否为空
func (self HashSet[T]) Empty() bool {
	return hashmap.HashMap[T, struct{}](self).Empty()
}

// ToSlice 转成切片
func (self HashSet[T]) ToSlice() []T {
	return hashmap.HashMap[T, struct{}](self).Keys()
}
