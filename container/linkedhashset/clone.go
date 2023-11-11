package linkedhashset

import (
	"github.com/kkkunny/stl/container/linkedhashmap"
)

func (self LinkedHashSet[T]) Clone() LinkedHashSet[T] {
	return LinkedHashSet[T](linkedhashmap.LinkedHashMap[T, struct{}](self).Clone())
}
