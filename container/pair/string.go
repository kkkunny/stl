package pair

import "fmt"

func (self Pair[T, F]) String() string {
	return fmt.Sprintf("(%v, %v)", self.First, self.Second)
}
