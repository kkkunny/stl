package linkedhashset

import (
	"github.com/kkkunny/stl/container/dynarray"
	"github.com/kkkunny/stl/container/linkedhashmap"
)

// Add 插入值
func (self *LinkedHashSet[T]) Add(v T) bool {
	exist := (*linkedhashmap.LinkedHashMap[T, struct{}])(self).ContainKey(v)
	if exist {
		return false
	}
	(*linkedhashmap.LinkedHashMap[T, struct{}])(self).Set(v, struct{}{})
	return true
}

// Contain 是否包含值
func (self LinkedHashSet[T]) Contain(v T) bool {
	return linkedhashmap.LinkedHashMap[T, struct{}](self).ContainKey(v)
}

// Remove 移除值
func (self *LinkedHashSet[T]) Remove(v T) bool {
	exist := (*linkedhashmap.LinkedHashMap[T, struct{}])(self).ContainKey(v)
	if !exist {
		return false
	}
	(*linkedhashmap.LinkedHashMap[T, struct{}])(self).Remove(v)
	return true
}

// Clear 清空
func (self *LinkedHashSet[T]) Clear() {
	(*linkedhashmap.LinkedHashMap[T, struct{}])(self).Clear()
}

// Empty 是否为空
func (self LinkedHashSet[T]) Empty() bool {
	return linkedhashmap.LinkedHashMap[T, struct{}](self).Empty()
}

// ToSlice 转成切片
func (self LinkedHashSet[T]) ToSlice() dynarray.DynArray[T] {
	return linkedhashmap.LinkedHashMap[T, struct{}](self).Keys()
}
