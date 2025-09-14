// 参考https://www.toolhelper.cn/Digit/UnitConvert?tab=byte

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

// 位

const (
	Bit Size = 1 // 位

	// 十进制

	Kbit = Bit * decimalSep
	Mbit = Kbit * decimalSep
	Gbit = Mbit * decimalSep
	Tbit = Gbit * decimalSep

	// 二进制

	Kibit = Bit * binarySep
	Mibit = Kibit * binarySep
	Gibit = Mibit * binarySep
	Tibit = Gibit * binarySep
)

// 字节

const (
	Byte = Bit * 8 // 字节

	// 十进制 byte

	KB = Byte * decimalSep
	MB = KB * decimalSep
	GB = MB * decimalSep
	TB = GB * decimalSep

	// 二进制 byte

	KiB = Byte * binarySep
	MiB = KiB * binarySep
	GiB = MiB * binarySep
	TiB = GiB * binarySep
)

type UnitType uint8

const (
	UnitTypeXbit = iota
	UnitTypeXibit
	UnitTypeXB
	UnitTypeXiB
)

var unitType2Name = [...][6]string{
	{"bit", "byte", "Kbit", "Mbit", "Gbit", "Tbit"},
	{"bit", "byte", "Kibit", "Mibit", "Gibit", "Tibit"},
	{"b", "B", "KB", "MB", "GB", "TB"},
	{"b", "B", "KiB", "MiB", "GiB", "TiB"},
}

var unitType2Value = [...][6]Size{
	{Bit, Byte, Kbit, Mbit, Gbit, Tbit},
	{Bit, Byte, Kibit, Mibit, Gibit, Tibit},
	{Bit, Byte, KB, MB, GB, TB},
	{Bit, Byte, KiB, MiB, GiB, TiB},
}

func (s Size) ToString(u UnitType) string {
	names, values := unitType2Name[u], unitType2Value[u]
	for i := len(names) - 1; i >= 0; i-- {
		if s/values[i] > 0 {
			return fmt.Sprintf("%.2f %s", float64(s)/float64(values[i]), names[i])
		}
	}
	return ""
}

func (s Size) String() string {
	return s.ToString(UnitTypeXiB)
}
