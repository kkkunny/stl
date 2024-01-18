package dynarray

import "encoding/json"

func (self DynArray[T]) MarshalJSON() ([]byte, error) {
	self.init()
	return json.Marshal(*self.data)
}

func (self *DynArray[T]) UnmarshalJSON(data []byte) error {
	self.init()
	return json.Unmarshal(data, self.data)
}
