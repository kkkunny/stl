package hashmap

import stlbasic "github.com/kkkunny/stl/basic"

func (self HashMap[K, V]) Equal(dst HashMap[K, V]) bool {
	self.init()
	dst.init()

	if self.buckets == dst.buckets {
		return true
	} else if self.Length() != dst.Length() {
		return false
	}

	for iter := self.KeyValues().Iterator(); iter.Next(); {
		pair := iter.Value()
		if !dst.ContainKey(pair.First) {
			return false
		} else if !stlbasic.Equal(pair.Second, dst.Get(pair.First)) {
			return false
		}
	}
	return true
}
