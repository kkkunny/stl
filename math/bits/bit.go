package stlbits

import "golang.org/x/exp/constraints"

// Bit 位
type Bit bool

func NewBit[T constraints.Integer](v T) Bit {
	return v != 0
}

func (self Bit) String() string {
	if self {
		return "1"
	} else {
		return "0"
	}
}

func (self Bit) Equal(dst Bit) bool {
	return self == dst
}

func (self Bit) Clone() Bit {
	return self
}

func (self Bit) Hash() uint64 {
	if self {
		return 1
	} else {
		return 0
	}
}

func (self Bit) Compare(dst Bit) int {
	if self.Hash() < dst.Hash() {
		return -1
	} else if self.Hash() == dst.Hash() {
		return 0
	} else {
		return 1
	}
}

func (self Bit) Not() Bit {
	return !self
}

func (self Bit) Default() Bit {
	return false
}
