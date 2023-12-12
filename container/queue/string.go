package queue

import (
	"fmt"
	"strings"
)

// String 获取字符串
func (self Queue[T]) String() string {
	var buf strings.Builder
	buf.WriteString("Queue{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
