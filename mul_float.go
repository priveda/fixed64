// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                  priveda/fixed64/[mul_float.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// MulFloat multiplies a fixed-point number by one or more floating-point
// numbers and returns the result. The original number is not changed.
func (n Fixed64) MulFloat(nums ...float64) Fixed64 {
	a := float64(n.i64)
	for _, b := range nums {
		//
		// check for negative or positive overflow
		limit := float64(maxInt64) / a
		if limit < 0 {
			limit = -limit
		}
		if b < -limit || b > limit {
			return fixed64Overflow(
				(a < 0 || b < 0) && (a > 0 || b > 0),
				a, " * ", b, " = ", a*b)
		}
		// multiply using int64, if there is no overflow
		n.i64 = int64(a * b)
	}
	return n
}

//end
