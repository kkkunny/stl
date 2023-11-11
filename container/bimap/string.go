package bimap

import (
	"fmt"
	"strings"
)

func (self BiMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteString("BiMap{")
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
