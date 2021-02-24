// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number         priveda/fixed64/[is_lesser_or_equal.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// IsLesserOrEqual returns true if the number is lesser or equal to value.
func (n Fixed64) IsLesserOrEqual(value interface{}) bool {
	return n.i64 <= New(value).i64
}

// end
