package linkedhashset

import (
	"encoding/json"
	"fmt"
)

func (self LinkedHashSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.ToSlice())
}

func (self *LinkedHashSet[T]) UnmarshalJSON(data []byte) error {
	fmt.Println(string(data))
	var slice []T
	if err := json.Unmarshal(data, &slice); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(slice)
	*self = NewLinkedHashSetWith[T](slice...)
	return nil
}
