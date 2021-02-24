// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                     priveda/fixed64/[unwrap.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// Unwrap returns the internal int64 used to store the fixed-point number.
func (n Fixed64) Unwrap() int64 {
	return n.i64
}

// end
