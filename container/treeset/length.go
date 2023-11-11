package treeset

import (
	"github.com/kkkunny/stl/container/treemap"
)

func (self TreeSet[T]) Length() uint {
	return treemap.TreeMap[T, struct{}](self).Length()
}
