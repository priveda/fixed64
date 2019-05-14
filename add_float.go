// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 A505E0                 priveda/fixed64/[add_float.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"math"
)

// AddFloat adds one or more floating-point numbers to a fixed-point
// number and returns the result. The original number is not changed.
func (ob Fixed64) AddFloat(add ...float64) Fixed64 {
	const lim float64 = math.MaxInt64 / 1E4
	a := ob.i64
	for _, b := range add {
		//
		// check for overflow
		if b < -lim || b > lim {
			return fixed64Overflow(
				(a < 0 || b < 0) && (a > 0 || b > 0),
				EOverflow, ": ", a, " * ", b, " = ", float64(a)*b,
			)
		}
		// use Add() because it has other overflow checks
		ob = ob.Add(Fixed64{int64(b * 1E4)})
	}
	return ob
}

//end
