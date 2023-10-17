package pair

import (
	"fmt"

	stlbasic "github.com/kkkunny/stl/basic"
)

// Pair 对
type Pair[T, F any] struct {
	First  T
	Second F
}

func NewPair[T, F any](first T, second F) Pair[T, F] {
	return Pair[T, F]{
		First:  first,
		Second: second,
	}
}

func (self Pair[T, F]) Debug(prefix uint) string {
	return fmt.Sprintf("pair(%s, %s)", stlbasic.Debug(self.First, prefix), stlbasic.Debug(self.Second, prefix))
}
