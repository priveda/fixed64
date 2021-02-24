// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                       priveda/fixed64/[wrap.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// Wrap initializes a fixed-point number from a scaled 64-bit integer.
// The decimal point is moved left 4 decimal places. For example,
// an input value of 15500 results in a fixed-point value of 1.55
func Wrap(i64 int64) Fixed64 {
	return Fixed64{i64: i64}
}

// end
