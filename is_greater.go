// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                 priveda/fixed64/[is_greater.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// IsGreater returns true if the number is greater than value.
func (n Fixed64) IsGreater(value interface{}) bool {
	return n.i64 > New(value).i64
}

//end
