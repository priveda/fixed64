// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                    priveda/fixed64/[is_zero.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// IsZero returns true if the number is equal to zero.
func (n Fixed64) IsZero() bool {
	return n.i64 == 0
}

// end
