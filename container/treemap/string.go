package treemap

import (
	"fmt"
	"strings"

	"github.com/HuKeping/rbtree"
)

// String 转成字符串
func (self TreeMap[K, V]) String() string {
	self.init()
	var buf strings.Builder
	buf.WriteString("TreeMap{")
	if !self.Empty() {
		var i uint
		self.tree.Ascend(self.tree.Min(), func(item rbtree.Item) bool {
			node := item.(*entry[K, V])
			buf.WriteString(fmt.Sprintf("%v", node.First))
			buf.WriteString(": ")
			buf.WriteString(fmt.Sprintf("%v", node.Second))
			if i < self.Length()-1 {
				buf.WriteString(", ")
			}
			i++
			return true
		})
	}
	buf.WriteByte('}')
	return buf.String()
}
