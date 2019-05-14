// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 B5CD7B                       priveda/fixed64/[fmt.go]
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
func (ob Fixed64) Fmt(decimalPlaces int) string {
	var (
		retBuf  = bytes.NewBuffer(make([]byte, 0, 25))
		ws      = retBuf.WriteString
		wr      = retBuf.WriteRune
		intLen  = 0
		intPart = ob.i64 / 1E4         // integer part of the number
		decPart = ob.i64 - intPart*1E4 // decimal part (as an int)
		neg     = ob.i64 < 0           // is it negative? use absolute value
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
			n := intPart / power
			if n > 0 {
				write = true
			}
			intPart -= n * power
			power /= 10
			if !write {
				continue
			}
			wr(rune(n + 48))
			digits--
			if power > 0 && digits <= 0 {
				ws(",")
				digits = 3
			}
		}
	}
	// write fractional part
	if decimalPlaces != 0 {
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
			n := int64(0)
			if power > 0 {
				n = decPart / power
				decPart -= n * power
				power /= 10
			}
			wr(rune(n + 48))
			if unfixed && decPart == 0 {
				break
			}
		}
	}
	return retBuf.String()
}

//end
