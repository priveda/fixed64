// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 38051E               priveda/fixed64/[is_negative.go]
// -----------------------------------------------------------------------------

package fixed64

// IsNegative returns true if the value of the currency object is negative.
func (ob Fixed64) IsNegative() bool {
	return ob.i64 < 0
}

//end
