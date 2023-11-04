package linkedlist

import stlbasic "github.com/kkkunny/stl/basic"

// Equal 比较相等
func (self LinkedList[T]) Equal(dst LinkedList[T]) bool {
	if self.root == dst.root {
		return true
	}

	if self.Length() != dst.Length() {
		return false
	}

	for c1, c2 := self.root, dst.root; c1 != nil && c2 != nil; c1, c2 = c1.Next, c2.Next {
		if !stlbasic.Equal(c1.Value, c2.Value) {
			return false
		}
	}
	return true
}
