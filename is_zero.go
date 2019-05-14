// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 5869F9                   priveda/fixed64/[is_zero.go]
// -----------------------------------------------------------------------------

package fixed64

// IsZero returns true if the number is equal to zero.
func (n Fixed64) IsZero() bool {
	return n.i64 == 0
}

//end
