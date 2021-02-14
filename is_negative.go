// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                priveda/fixed64/[is_negative.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// IsNegative returns true if the number is negative.
func (n Fixed64) IsNegative() bool {
	return n.i64 < 0 && n.i64 != NaN
}

// end
