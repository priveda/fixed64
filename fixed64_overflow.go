// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number           priveda/fixed64/[fixed64_overflow.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

import (
	"math"
)

// fixed64Overflow returns a negative (math.MinInt64+1)
// or positive (math.MaxInt64) overflow value.
//
// It also calls Error() to write an error in the log. The error is logged,
// but the function returns an overflow value instead of Fixed64{NaN}.
//
// isNegative: should specify if the number is negative or positive.
// args: an array of values used to build the error message.
func fixed64Overflow(isNegative bool, args ...interface{}) Fixed64 {
	ar := []interface{}{EOverflow + ": "}
	ar = append(ar, args...)
	mod.Error(ar...)
	if isNegative {
		return Fixed64{math.MinInt64 + 1}
	}
	return Fixed64{math.MaxInt64}
}

//end
