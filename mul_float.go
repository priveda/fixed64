// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:56:01 D07BBC                 priveda/fixed64/[mul_float.go]
// -----------------------------------------------------------------------------

package fixed64

// MulFloat multiplies a currency object by one or more floating-point
// numbers and returns the result. The object's value isn't changed.
func (ob Fixed64) MulFloat(multiply ...float64) Fixed64 {
	a := float64(ob.i64)
	for _, b := range multiply {
		//
		// check for negative or positive overflow
		lim := float64(maxInt64) / a
		if lim < 0 {
			lim = -lim
		}
		if b < -lim || b > lim {
			return fixed64Overflow(
				(a < 0 || b < 0) && (a > 0 || b > 0),
				"Overflow: ", a, " * ", b, " = ", a*b,
			)
		}
		// multiply using int64, if there is no overflow
		ob.i64 = int64(a * b)
	}
	return ob
}

//end
