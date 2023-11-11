package treeset

import (
	"github.com/kkkunny/stl/container/treemap"
)

func (self TreeSet[T]) Clone() TreeSet[T] {
	return TreeSet[T](treemap.TreeMap[T, struct{}](self).Clone())
}
