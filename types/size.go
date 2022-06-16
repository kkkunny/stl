package types

import "fmt"

type Size uint64

const (
	Bit  Size = 1        // b
	Byte      = Bit << 3 // B

	KiloByte = Byte << 10     // KB
	MegaByte = KiloByte << 10 // MB
	GigaByte = MegaByte << 10 // GB
	TeraByte = GigaByte << 10 // TB

	Kilo = Byte * 1e3 // K
	Mega = Kilo * 1e3 // M
	Giga = Mega * 1e3 // G
	Tera = Giga * 1e3 // T
)

func (self Size) String() string {
	if self > TeraByte {
		return fmt.Sprintf("%.2fTB", self.TeraByte())
	} else if self > GigaByte {
		return fmt.Sprintf("%.2fGB", self.GigaByte())
	} else if self > MegaByte {
		return fmt.Sprintf("%.2fMB", self.MegaByte())
	} else if self > KiloByte {
		return fmt.Sprintf("%.2fKB", self.KiloByte())
	} else if self > Byte {
		return fmt.Sprintf("%dB", self.Byte())
	} else {
		return fmt.Sprintf("%db", self)
	}
}

func (self Size) Bit() uint64 {
	return uint64(self)
}

func (self Size) Byte() uint64 {
	return uint64(self) / uint64(Byte)
}

func (self Size) KiloByte() float64 {
	return float64(self) / float64(KiloByte)
}

func (self Size) MegaByte() float64 {
	return float64(self) / float64(MegaByte)
}

func (self Size) GigaByte() float64 {
	return float64(self) / float64(GigaByte)
}

func (self Size) TeraByte() float64 {
	return float64(self) / float64(TeraByte)
}

func (self Size) Kilo() float64 {
	return float64(self) / float64(Kilo)
}

func (self Size) Mega() float64 {
	return float64(self) / float64(Mega)
}

func (self Size) Giga() float64 {
	return float64(self) / float64(Giga)
}

func (self Size) Tera() float64 {
	return float64(self) / float64(Tera)
}
