package stlos

import (
	"fmt"
)

const (
	decimalSep = 1000
	binarySep  = 1024
)

// Size 大小
type Size uint64

const (
	Bit  Size = 1       // 比特、位
	Byte      = Bit * 8 // 字节
)

// 十进制 bit
const (
	Kbit Size = Bit * decimalSep
	Mbit      = Kbit * decimalSep
	Gbit      = Mbit * decimalSep
	Tbit      = Gbit * decimalSep
	Pbit      = Tbit * decimalSep
	// Ebit      = Pbit * decimalSep
)

// 十进制 byte
const (
	Kb Size = Byte * decimalSep
	Mb      = Kb * decimalSep
	Gb      = Mb * decimalSep
	Tb      = Gb * decimalSep
	Pb      = Tb * decimalSep
	// Eb      = Pb * decimalSep
)

// 二进制 bit
const (
	Kibit Size = Bit * binarySep
	Mibit      = Kibit * binarySep
	Gibit      = Mibit * binarySep
	Tibit      = Gibit * binarySep
	Pibit      = Tibit * binarySep
	// Eibit      = Pibit * binarySep
)

// 二进制 byte
const (
	Kib Size = Byte * binarySep
	Mib      = Kib * binarySep
	Gib      = Mib * binarySep
	Tib      = Gib * binarySep
	Pib      = Tib * binarySep
	// Eib      = Pib * binarySep
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

func (self Size) SimpleString() string {
	switch {
	// case self/Eib > 0:
	// 	return fmt.Sprintf("%.2f Eib", float64(self/Pib)/float64(binarySep))
	case self/Pib > 0:
		return fmt.Sprintf("%.2f Pib", float64(self/Tib)/float64(binarySep))
	case self/Tib > 0:
		return fmt.Sprintf("%.2f Tib", float64(self/Gib)/float64(binarySep))
	case self/Gib > 0:
		return fmt.Sprintf("%.2f Gib", float64(self/Mib)/float64(binarySep))
	case self/Mib > 0:
		return fmt.Sprintf("%.2f Mib", float64(self/Kib)/float64(binarySep))
	case self/Kib > 0:
		return fmt.Sprintf("%.2f Kib", float64(self)/float64(Kib))
	case self/Byte > 0:
		return fmt.Sprintf("%.2f Byte", float64(self)/float64(Byte))
	default:
		return fmt.Sprintf("%d Bit", self)
	}
}
