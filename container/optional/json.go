package optional

import "encoding/json"

func (op Optional[T]) MarshalJSON() ([]byte, error) {
	if op.IsNone() {
		return []byte("null"), nil
	}
	return json.Marshal(*op.data)
}

func (op *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		op.data = nil
		return nil
	}
	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	op.data = &v
	return nil
}
