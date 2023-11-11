package treeset

import (
	"github.com/kkkunny/stl/container/treemap"
)

func (self TreeSet[T]) Equal(dst TreeSet[T]) bool {
	return treemap.TreeMap[T, struct{}](self).Equal(treemap.TreeMap[T, struct{}](dst))
}
