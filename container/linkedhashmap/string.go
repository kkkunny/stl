package linkedhashmap

import (
	"fmt"
	"strings"
)

func (self LinkedHashMap[K, V]) String() string {
	// TODO: 优化
	var buf strings.Builder
	buf.WriteString("LinkedHashMap{")
	var i int
	for cursor := self.list.Front(); cursor != nil; cursor = cursor.Next() {
		buf.WriteString(fmt.Sprintf("%v", cursor.Value.First))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", cursor.Value.Second))
		if cursor.Next() != nil {
			buf.WriteString(", ")
		}
		i++
	}
	buf.WriteByte('}')
	return buf.String()
}
