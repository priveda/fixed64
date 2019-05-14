// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 7FDC10                   priveda/fixed64/[is_zero.go]
// -----------------------------------------------------------------------------

package fixed64

// IsZero returns true if the value of the currency object is zero.
func (ob Fixed64) IsZero() bool {
	return ob.val == 0
}

//end
