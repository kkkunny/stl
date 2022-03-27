package types

type I8 int8

func (self I8) Hash() Usize {
	return Usize(self)
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

func (self I16) Hash() Usize {
	return Usize(self)
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

func (self I32) Hash() Usize {
	return Usize(self)
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

func (self I64) Hash() Usize {
	return Usize(self)
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

func (self Isize) Hash() Usize {
	return Usize(self)
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

func (self U8) Hash() Usize {
	return Usize(self)
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

func (self U16) Hash() Usize {
	return Usize(self)
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

func (self U32) Hash() Usize {
	return Usize(self)
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

func (self U64) Hash() Usize {
	return Usize(self)
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

func (self Usize) Hash() Usize {
	return self
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

func (self F32) Hash() Usize {
	return Usize(self)
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

func (self F64) Hash() Usize {
	return Usize(self)
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

type String string

func (self String) Hash() Usize {
	var code Usize
	for _, c := range self {
		code += I32(c).Hash()
	}
	return code
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
