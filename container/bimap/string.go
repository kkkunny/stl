package bimap

import (
	"fmt"
	"strings"
)

func (self BiMap[K, V]) String() string {
	var buf strings.Builder
	buf.WriteString("BiMap{")
	for i, pair := range self.KeyValues() {
		buf.WriteString(fmt.Sprintf("%v", pair.First))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprintf("%v", pair.Second))
		if i < int(self.Length())-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
