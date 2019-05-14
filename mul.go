// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:41:00 79E694                       priveda/fixed64/[mul.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"math/big"
)

// Mul multiplies a fixed-point number by one or more fixed-point
// numbers and returns the result. The original number is not changed.
func (ob Fixed64) Mul(multiply ...Fixed64) Fixed64 {
	for _, nb := range multiply {
		var (
			a = ob.i64
			b = nb.i64
		)
		// return zero if either number is zero
		if a == 0 || b == 0 {
			return Fixed64{0}
		}
		// if direct multiplication will overflow int64, use big.Int
		lim := maxInt64 / a
		if lim < 0 {
			lim = -lim
		}
		if b >= lim || b <= -lim {
			c := big.NewInt(a)
			c.Mul(c, big.NewInt(b))
			c.Div(c, big1E4)
			//
			// if result can't be stored in Fixed64, return overflow
			//
			// TODO: IsInt64() is not available in older Go versions ``
			overflow := !c.IsInt64()
			var ret int64
			if !overflow {
				ret = c.Int64()
				if ret < minInt64 || ret > maxInt64 {
					overflow = true
				}
			}
			if overflow {
				return fixed64Overflow(
					(a < 0 || b < 0) && (a > 0 || b > 0),
					EOverflow, ": ", a, " * ", b, " = ", c,
				)
			}
			return Fixed64{ret}
		}
		ob.i64 = a * b / 1E4
	}
	return ob
}

//end
