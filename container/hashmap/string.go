package hashmap

import (
	"fmt"
	"strings"
)

func (self HashMap[K, V]) String() string {
	self.init()

	var buf strings.Builder
	buf.WriteString("HashMap{")
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
