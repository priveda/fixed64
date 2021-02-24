// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                        priveda/fixed64/[int.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// Int returns the fixed-point number as an int value.
func (n Fixed64) Int() int64 {
	return n.i64 / 1e4
}

// end
