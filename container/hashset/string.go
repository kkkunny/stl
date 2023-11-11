package hashset

import (
	"fmt"
	"strings"
)

func (self HashSet[T]) String() string {
	var buf strings.Builder
	buf.WriteString("HashSet{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
