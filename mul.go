// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 79DED9                       priveda/fixed64/[mul.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"math/big"
)

// Mul multiplies a currency object by one or more currency values
// and returns the result. The object's value isn't changed.
func (ob Fixed64) Mul(multiply ...Fixed64) Fixed64 {
	for _, cur := range multiply {
		var (
			a = ob.val
			b = cur.val
		)
		// return zero if either number is zero
		if a == 0 || b == 0 {
			return Fixed64{0}
		}
		// if direct multiplication will overflow int64, use big.Int
		lim := MaxFixed64I64 / a
		if lim < 0 {
			lim = -lim
		}
		if b >= lim || b <= -lim {
			n := big.NewInt(a)
			n.Mul(n, big.NewInt(b))
			n.Div(n, big1E4)
			//
			// if result can't be stored in Fixed64, return overflow
			//
			// TODO: IsInt64() is not available in older Go versions ``
			overflow := !n.IsInt64()
			var ret int64
			if !overflow {
				ret = n.Int64()
				if ret < MinFixed64I64 || ret > MaxFixed64I64 {
					overflow = true
				}
			}
			if overflow {
				return currencyOverflow(
					(a < 0 || b < 0) && (a > 0 || b > 0),
					EOverflow, ": ", a, " * ", b, " = ", n,
				)
			}
			return Fixed64{ret}
		}
		ob.val = a * b / 1E4
	}
	return ob
}

//end
