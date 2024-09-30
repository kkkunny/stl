package hashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (self HashMap[K, V]) MarshalJSON() ([]byte, error) {
	self.init()

	var buf bytes.Buffer
	err := buf.WriteByte('{')
	if err != nil {
		return nil, err
	}
	for i, pair := range self.KeyValues() {
		_, err = buf.WriteString(fmt.Sprintf("\"%+v\"", pair.First))
		if err != nil {
			return nil, err
		}
		err = buf.WriteByte(':')
		if err != nil {
			return nil, err
		}
		vs, err := json.Marshal(pair.Second)
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(vs)
		if err != nil {
			return nil, err
		}
		if i < int(self.Length())-1 {
			err = buf.WriteByte(',')
			if err != nil {
				return nil, err
			}
		}
	}
	err = buf.WriteByte('}')
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
