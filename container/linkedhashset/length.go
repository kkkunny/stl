package linkedhashset

import (
	"github.com/kkkunny/stl/container/linkedhashmap"
)

func (self LinkedHashSet[T]) Length() uint {
	return linkedhashmap.LinkedHashMap[T, struct{}](self).Length()
}
