package treeset

import (
	"fmt"
	"strings"
)

func (self TreeSet[T]) String() string {
	var buf strings.Builder
	buf.WriteString("TreeSet{")
	for iter := self.Iterator(); iter.Next(); {
		buf.WriteString(fmt.Sprintf("%v", iter.Value()))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
