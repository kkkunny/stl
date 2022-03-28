package types

import (
	"math"
	"unsafe"
)

type I8 int8

func (self I8) Hash() int32 {
	return int32(self)
}

func (self I8) Compare(dst I8) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type I16 int16

func (self I16) Hash() int32 {
	return int32(self)
}

func (self I16) Compare(dst I16) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type I32 int32

func (self I32) Hash() int32 {
	return int32(self)
}

func (self I32) Compare(dst I32) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type I64 int64

func (self I64) Hash() int32 {
	return int32(self ^ (self >> 32))
}

func (self I64) Compare(dst I64) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type Isize int

func (self Isize) Hash() int32 {
	return int32(self ^ (self >> 32))
}

func (self Isize) Compare(dst Isize) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type U8 uint8

func (self U8) Hash() int32 {
	return int32(*(*int8)(unsafe.Pointer(&self)))
}

func (self U8) Compare(dst U8) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type U16 uint16

func (self U16) Hash() int32 {
	return int32(*(*int16)(unsafe.Pointer(&self)))
}

func (self U16) Compare(dst U16) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type U32 uint32

func (self U32) Hash() int32 {
	return *(*int32)(unsafe.Pointer(&self))
}

func (self U32) Compare(dst U32) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type U64 uint64

func (self U64) Hash() int32 {
	hash := *(*int64)(unsafe.Pointer(&self))
	return int32(hash ^ (hash >> 32))
}

func (self U64) Compare(dst U64) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type Usize uint

func (self Usize) Hash() int32 {
	hash := *(*int64)(unsafe.Pointer(&self))
	return int32(hash ^ (hash >> 32))
}

func (self Usize) Compare(dst Usize) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type F32 float32

func (self F32) Hash() int32 {
	return int32(math.Float32bits(float32(self)))
}

func (self F32) Compare(dst F32) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type F64 float64

func (self F64) Hash() int32 {
	bits := math.Float64bits(float64(self))
	return int32(bits ^ (bits >> 32))
}

func (self F64) Compare(dst F64) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type Bool bool

func (self Bool) Hash() int32 {
	if self {
		return 1231
	} else {
		return 1237
	}
}

type Char rune

func (self Char) Hash() int32 {
	return int32(self)
}

func (self Char) Compare(dst Char) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}

type String string

func (self String) Hash() int32 {
	var hash int32
	for _, c := range self {
		hash = hash*31 + int32(c)
	}
	return hash
}

func (self String) Compare(dst String) int {
	if self < dst {
		return -1
	} else if self > dst {
		return 1
	} else {
		return 0
	}
}
