package stlos

import (
	"fmt"
)

// Size 大小
type Size uint64

const (
	Bit  Size = 1       // 比特、位
	Byte      = Bit * 8 // 字节
)

// 十进制 bit
const (
	Kbit Size = Bit * 1000
	Mbit      = Kbit * 1000
	Gbit      = Mbit * 1000
	Tbit      = Gbit * 1000
	Pbit      = Tbit * 1000
	// Ebit      = Pbit * 1000
)

// 十进制 byte
const (
	Kb Size = Byte * 1000
	Mb      = Kb * 1000
	Gb      = Mb * 1000
	Tb      = Gb * 1000
	Pb      = Tb * 1000
	// Eb      = Pb * 1000
)

// 二进制 bit
const (
	Kibit Size = Bit * 1024
	Mibit      = Kibit * 1024
	Gibit      = Mibit * 1024
	Tibit      = Gibit * 1024
	Pibit      = Tibit * 1024
	// Eibit      = Pibit * 1024
)

// 二进制 byte
const (
	Kib Size = Byte * 1024
	Mib      = Kib * 1024
	Gib      = Mib * 1024
	Tib      = Gib * 1024
	Pib      = Tib * 1024
	// Eib      = Pib * 1024
)

func (self Size) String() string {
	var prefix string
	var more Size

	switch {
	// case self/Eib > 0:
	// 	prefix = fmt.Sprintf("%d Eib", self/Eib)
	// 	more = self % Eib
	case self/Pib > 0:
		prefix = fmt.Sprintf("%d Pib", self/Pib)
		more = self % Pib
	case self/Tib > 0:
		prefix = fmt.Sprintf("%d Tib", self/Tib)
		more = self % Tib
	case self/Gib > 0:
		prefix = fmt.Sprintf("%d Gib", self/Gib)
		more = self % Gib
	case self/Mib > 0:
		prefix = fmt.Sprintf("%d Mib", self/Mib)
		more = self % Mib
	case self/Kib > 0:
		prefix = fmt.Sprintf("%d Kib", self/Kib)
		more = self % Kib
	case self/Byte > 0:
		prefix = fmt.Sprintf("%d Byte", self/Byte)
		more = self % Byte
	default:
		return fmt.Sprintf("%d Bit", self)
	}

	if more == 0 {
		return prefix
	}
	return fmt.Sprintf("%s %s", prefix, more.String())
}

func (self Size) Clone() Size {
	return self
}

func (self Size) Equal(dst Size) bool {
	return self == dst
}

func (self Size) Hash() uint64 {
	return uint64(self)
}
