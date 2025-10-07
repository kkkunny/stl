package optional

import "fmt"

func (op Optional[T]) String() string {
	if op.IsNone() {
		return "None"
	}
	return fmt.Sprintf("%v", *op.data)
}
