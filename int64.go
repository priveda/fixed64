// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                      priveda/fixed64/[int64.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// Int64 returns the fixed-point number as an int64 value.
func (n Fixed64) Int64() int64 {
	return n.i64 / 1e4
}

// end
