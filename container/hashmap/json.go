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
	for iter := self.KeyValues().Iterator(); iter.Next(); {
		pair := iter.Value()
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
		if iter.HasNext() {
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
