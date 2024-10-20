package linkedlist

import (
	"fmt"
	"strings"
)

// String 获取字符串
func (self LinkedList[T]) String() string {
	var buf strings.Builder
	buf.WriteString("LinkedList{")
	for cursor := self.root; cursor != nil; cursor = cursor.Next {
		buf.WriteString(fmt.Sprintf("%v", cursor.Value))
		if cursor.Next != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
