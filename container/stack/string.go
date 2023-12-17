package stack

import (
	"fmt"
	"strings"
)

// String 获取字符串
func (self Stack[T]) String() string {
	var buf strings.Builder
	buf.WriteString("Stack{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
