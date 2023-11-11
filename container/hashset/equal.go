package hashset

import (
	"github.com/kkkunny/stl/container/hashmap"
)

func (self HashSet[T]) Equal(dst HashSet[T]) bool {
	return hashmap.HashMap[T, struct{}](self).Equal(hashmap.HashMap[T, struct{}](dst))
}
