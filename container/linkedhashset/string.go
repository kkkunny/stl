package linkedhashset

import (
	"fmt"
	"strings"
)

func (self LinkedHashSet[T]) String() string {
	var buf strings.Builder
	buf.WriteString("LinkedHashSet{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
