package hashmap

import (
	"fmt"
	"strings"
)

func (self HashMap[K, V]) String() string {
	self.init()

	var buf strings.Builder
	buf.WriteString("HashMap{")
	for iter := self.KeyValues().Iterator(); iter.Next(); {
		pair := iter.Value()
		buf.WriteString(fmt.Sprintf("%v", pair.First))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", pair.Second))
		if iter.HasNext() {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
