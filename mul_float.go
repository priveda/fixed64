// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 558858                 priveda/fixed64/[mul_float.go]
// -----------------------------------------------------------------------------

package fixed64

// MulFloat multiplies a currency object by one or more floating-point
// numbers and returns the result. The object's value isn't changed.
func (ob Fixed64) MulFloat(multiply ...float64) Fixed64 {
	a := float64(ob.val)
	for _, b := range multiply {
		// check for negative or positive overflow
		lim := MaxFixed64I64 / a
		if lim < 0 {
			lim = -lim
		}
		if b < -lim || b > lim {
			return currencyOverflow(
				(a < 0 || b < 0) && (a > 0 || b > 0),
				"Overflow: ", a, " * ", b, " = ", a*b,
			)
		}
		// multiply using int64, if there is no overflow
		ob.val = int64(a * b)
	}
	return ob
}

//end
