// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                    priveda/fixed64/[float64.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// Float64 returns the fixed-point number as a float64 value.
func (n Fixed64) Float64() float64 {
	return float64(n.i64) / 1e4
}

// end
