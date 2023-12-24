package queue

import (
	"fmt"
	"strings"
)

// String 获取字符串
func (self PriorityQueue[T]) String() string {
	var buf strings.Builder
	buf.WriteString("PriorityQueue{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%d", iter.Value().First))
		buf.WriteByte(':')
		buf.WriteString(fmt.Sprintf("%v", iter.Value().Second))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
