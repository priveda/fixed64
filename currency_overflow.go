// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:56:01 FD6132         priveda/fixed64/[currency_overflow.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"math"
)

// fixed64Overflow returns a negative (math.MinInt64)
// or positive (math.MaxInt64) overflow value.
//
// It also calls Error() to write an error in the log. The error is logged,
// but the function returns an int64 overflow value instead of an error.
//
// isNegative: should specify if the number is negative or positive.
// a: an array of values used to build the error message.
func fixed64Overflow(isNegative bool, a ...interface{}) Fixed64 {
	mod.Error(a...)
	if isNegative {
		return Fixed64{math.MinInt64}
	}
	return Fixed64{math.MaxInt64}
}

//end
