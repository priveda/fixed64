// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number        priveda/fixed64/[is_greater_or_equal.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// IsGreaterOrEqual returns true if the number is greater or equal to value.
func (n Fixed64) IsGreaterOrEqual(value interface{}) bool {
	return n.i64 >= New(value).i64
}

// end
