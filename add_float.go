// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                  priveda/fixed64/[add_float.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

import (
	"math"
)

// AddFloat adds one or more floating-point numbers to a fixed-point
// number and returns the result. The original number is not changed.
func (n Fixed64) AddFloat(nums ...float64) Fixed64 {
	const limit float64 = math.MaxInt64 / 1e4
	a := n.i64
	for _, b := range nums {
		//
		// check for overflow
		if b < -limit || b > limit {
			return fixed64Overflow(
				(a < 0 || b < 0) && (a > 0 || b > 0),
				a, " * ", b, " = ", float64(a)*b)
		}
		// use Add() because it has other overflow checks
		n = n.Add(Fixed64{int64(b * 1e4)})
	}
	return n
}

// end
