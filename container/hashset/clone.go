package hashset

import "github.com/kkkunny/stl/container/hashmap"

func (self HashSet[T]) Clone() HashSet[T] {
	return HashSet[T](hashmap.HashMap[T, struct{}](self).Clone())
}
