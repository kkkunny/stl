package dynarray

import (
	"fmt"
	"strings"
)

// String 获取字符串
func (self DynArray[T]) String() string {
	self.init()
	var buf strings.Builder
	buf.WriteString("DynArray{")
	for i, v := range *self.data {
		buf.WriteString(fmt.Sprintf("%v", v))
		if i < len(*self.data)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
