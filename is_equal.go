// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                   priveda/fixed64/[is_equal.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// IsEqual returns true if the number is equal to value.
func (n Fixed64) IsEqual(value interface{}) bool {
	return n.i64 == New(value).i64
}

//end
