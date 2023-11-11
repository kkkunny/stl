package hashset

import "github.com/kkkunny/stl/container/hashmap"

func (self HashSet[T]) Length() uint {
	return hashmap.HashMap[T, struct{}](self).Length()
}
