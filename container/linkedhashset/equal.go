package linkedhashset

import (
	"github.com/kkkunny/stl/container/linkedhashmap"
)

func (self LinkedHashSet[T]) Equal(dst LinkedHashSet[T]) bool {
	return linkedhashmap.LinkedHashMap[T, struct{}](self).Equal(linkedhashmap.LinkedHashMap[T, struct{}](dst))
}
