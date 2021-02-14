// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                        priveda/fixed64/[div.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// Div divides a fixed-point number by one or more other fixed-point
// numbers and returns the result. The original number is not changed.
func (n Fixed64) Div(nums ...Fixed64) Fixed64 {
	for _, num := range nums {
		n.i64 *= 1e4
		n.i64 /= num.i64
	}
	return n
}

// end
