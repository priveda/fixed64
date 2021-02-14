// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                        priveda/fixed64/[fmt.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

import (
	"bytes"
)

// Fmt returns the fixed-point number as a a string
// delimited with commas (grouped every three digits)
// and having the specified number of decimal places.
// When decimalPlaces is negative, the resulting
// number's decimals will vary.
func (n Fixed64) Fmt(decimalPlaces int) string {
	if n.i64 == NaN {
		return ""
	}
	var (
		retBuf  = bytes.NewBuffer(make([]byte, 0, 25))
		ws      = retBuf.WriteString
		wr      = retBuf.WriteRune
		intLen  = 0
		intPart = n.i64 / 1e4         // integer part of the number
		decPart = n.i64 - intPart*1e4 // decimal part (as an int)
		neg     = n.i64 < 0           // is it negative? use absolute value
	)
	if neg {
		intPart = -intPart
		ws("-")
	}
	// calculate length of number's integer part
	for limit := int64(0); intLen < 15; intLen++ {
		if intPart <= limit {
			break
		}
		limit *= 10
		limit += 9
	}
	// write delimited integer part
	if intLen == 0 {
		ws("0")
	} else {
		var (
			write  = false
			power  = int64(100000000000000) // 10^14
			digits = intLen % 3
		)
		if digits == 0 {
			digits = 3
		}
		for power > 0 {
			x := intPart / power
			if x > 0 {
				write = true
			}
			intPart -= x * power
			power /= 10
			if !write {
				continue
			}
			wr(rune(x + 48))
			digits--
			if power > 0 && digits <= 0 {
				ws(",")
				digits = 3
			}
		}
	}
	if decimalPlaces != 0 {
		fixedNumFmtWriteDecimals(decimalPlaces, decPart, ws)
	}
	return retBuf.String()
}

// fixedNumFmtWriteDecimals writes the fractional part of a number
// It is only called by Fixed64.Fmt()
func fixedNumFmtWriteDecimals(
	decimalPlaces int,
	decPart int64,
	ws func(string) (int, error),
) {
	power := int64(1000) // 10^3
	unfixed := decPart > 0 && decimalPlaces < 0
	if unfixed {
		decimalPlaces = 4
	}
	if decimalPlaces > 0 {
		ws(".")
	}
	for decimalPlaces > 0 {
		decimalPlaces--
		x := int64(0)
		if power > 0 {
			x = decPart / power
			decPart -= x * power
			power /= 10
		}
		ws(string(rune(x + 48)))
		if unfixed && decPart == 0 {
			break
		}
	}
}

// end
