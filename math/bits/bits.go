package stlbits

import (
	"strings"
	"unsafe"

	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Bits 位列表
type Bits []Bit

func NewFromIntegerWithLength[T constraints.Integer](v T, length uint64) Bits {
	self := make(Bits, length)
	for i := int(length) - 1; i >= 0; i-- {
		vv := v >> 1
		self[i] = NewBit(v - vv<<1)
		v = vv
	}
	return self
}

func NewFromInteger[T constraints.Integer](v T) Bits {
	return NewFromIntegerWithLength(v, uint64(unsafe.Sizeof(v))*8).Reduce()
}

func (self Bits) String() string {
	var buf strings.Builder
	for _, bit := range self {
		buf.WriteString(bit.String())
	}
	return buf.String()
}

func (self Bits) Length() uint {
	return uint(len(self))
}

func (self Bits) Equal(dst Bits) bool {
	for i := 0; i < len(self); i++ {
		if self[i] != dst[i] {
			return false
		}
	}
	return true
}

func (self Bits) Clone() Bits {
	dst := make(Bits, len(self), cap(self))
	copy(dst, self)
	return dst
}

func (self Bits) Hash() uint64 {
	var v uint64
	for _, bit := range self {
		if bit {
			v = (v << 1) | 1
		} else {
			v = (v << 1) | 0
		}
	}
	return v
}

func (self Bits) Order(dst Bits) int {
	var selfCpy, dstCpy Bits
	if len(self) < len(dst) {
		selfCpy, dstCpy = make([]Bit, len(dst)), dst
		copy(selfCpy[len(dst)-len(self):], self)
	} else if len(self) == len(dst) {
		selfCpy, dstCpy = self, dst
	} else {
		selfCpy, dstCpy = self, make([]Bit, len(self))
		copy(dstCpy[len(self)-len(dst):], dst)
	}

	for i := 0; i < len(selfCpy); i++ {
		if order := selfCpy[i].Order(dstCpy[i]); order > 0 {
			return 1
		} else if order < 0 {
			return -1
		}
	}
	return 0
}

// Reduce 缩小
func (self Bits) Reduce() Bits {
	for i, bit := range self {
		if bit {
			return self[i:]
		}
	}
	return self
}

// Reverse 反码
func (self Bits) Reverse() Bits {
	return lo.Map(self, func(item Bit, _ int) Bit {
		return item.Not()
	})
}

// Complement 补码
func (self Bits) Complement() Bits {
	self = self.Reverse()
	var add bool
	for i := len(self) - 1; i >= 0; i-- {
		if i == len(self)-1 || add {
			add = self[i] == true
			self[i] = !self[i]
		} else {
			break
		}
	}
	return self
}

func (self Bits) SignedInteger() int64 {
	neg := self[0]

	if neg {
		self = self.Complement()
	}
	num := int64(self[1:].UnsignedInteger())

	if !neg {
		return num
	} else {
		return 0 - num
	}
}

func (self Bits) UnsignedInteger() uint64 {
	var num uint64
	for i, bit := range self {
		if i != 0 {
			num <<= 1
		}
		if bit {
			num |= 1
		}
	}
	return num
}

func (self Bits) Not() {

}

func (self Bits) Default() Bits {
	return make(Bits, 0)
}
