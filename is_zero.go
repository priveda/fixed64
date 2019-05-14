// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 6CF380                   priveda/fixed64/[is_zero.go]
// -----------------------------------------------------------------------------

package fixed64

// IsZero returns true if the value of the currency object is zero.
func (ob Fixed64) IsZero() bool {
	return ob.i64 == 0
}

//end
