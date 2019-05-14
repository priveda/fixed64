// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 9A255F               priveda/fixed64/[is_negative.go]
// -----------------------------------------------------------------------------

package fixed64

// IsNegative returns true if the number is negative.
func (ob Fixed64) IsNegative() bool {
	return ob.i64 < 0
}

//end
