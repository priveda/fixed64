// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 6708ED                   priveda/fixed64/[is_zero.go]
// -----------------------------------------------------------------------------

package fixed64

// IsZero returns true if the number is equal to zero.
func (ob Fixed64) IsZero() bool {
	return ob.i64 == 0
}

//end
