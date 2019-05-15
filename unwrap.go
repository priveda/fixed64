// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-15 16:46:15 5567C8                    priveda/fixed64/[unwrap.go]
// -----------------------------------------------------------------------------

package fixed64

// Unwrap returns the internal int64 used to store the fixed-point number.
func (n Fixed64) Unwrap() int64 {
	return n.i64
}

//end
