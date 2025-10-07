package stlbasic

import "golang.org/x/exp/constraints"

// Number 数字
type Number interface {
	constraints.Integer | constraints.Float
}
