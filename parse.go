// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                      priveda/fixed64/[parse.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

import (
	"fmt"
	"unicode"
)

// Parse converts a numeric string with up to 4 decimals to a Fixed64.
//
// If the input string is whitespace-only or is zero-length,
// returns a zero-value Fixed64{0} and no error.
//
// If the string is not numeric, returns NaN and an error (which is logged).
//
func Parse(s string) (Fixed64, error) {
	var (
		hasMinus bool  // has minus sign (negative result)
		hasDP    bool  // has decimal point?
		count    int   //
		decs     int   // number of decimals
		ret      int64 // returned value
	)
	for _, r := range s {
		switch {
		case r == '+':
			{
				count--
			}
		case r == '-':
			{
				hasMinus = true
			}
		case r == '.':
			{
				hasDP = true
			}
		case r >= '0' && r <= '9':
			{
				if hasDP {
					decs++
					if decs > 4 {
						break
					}
				}
				ret *= 10
				ret += int64(r - '0')
			}
		case r == ',', unicode.IsSpace(r), r == '\a', r == '\b':
			{
				count--
			}
		default:
			erm := EInvalidValue + fmt.Sprintf(": Fixed(%q)", s)
			return Fixed64{NaN}, mod.Error(erm)
		}
		count++
	}
	if count == 0 {
		return Fixed64{0}, nil
	}
	for decs < 4 {
		decs++
		ret *= 10
	}
	if hasMinus {
		ret = -ret
	}
	return Fixed64{ret}, nil
}

// end
