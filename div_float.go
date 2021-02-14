// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                  priveda/fixed64/[div_float.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// DivFloat divides a fixed-point number by one or more floating-point
// numbers and returns the result. The original number is not changed.
func (n Fixed64) DivFloat(nums ...float64) Fixed64 {
	for _, num := range nums {
		n.i64 *= 1e4
		n.i64 /= int64(num * 1e4)
	}
	return n
}

//end
