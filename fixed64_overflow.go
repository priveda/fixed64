// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:58:50 1F2754          priveda/fixed64/[fixed64_overflow.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"math"
)

// fixed64Overflow returns a negative (math.MinInt64+1)
// or positive (math.MaxInt64) overflow value.
//
// It also logs an error. The error is logged, but the function
// returns an int64 overflow value instead of an NaN.
//
// isNegative: should specify if the number is negative or positive.
// ar: an array of values used to build the error message.
//
func fixed64Overflow(isNegative bool, ar ...interface{}) Fixed64 {
	mod.Error(ar...)
	if isNegative {
		return Fixed64{math.MinInt64 + 1}
	}
	return Fixed64{math.MaxInt64}
}

//end
