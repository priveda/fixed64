// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                  priveda/fixed64/[constants.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// Fixed64 represents a fixed-point number with up to four decimal places.
// It is stored internally as a 64-bit integer. This gives it a range
// from -922,337,203,685,477.5808 to 922,337,203,685,477.5807
type Fixed64 struct {
	i64 int64
}

// end
