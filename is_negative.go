// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 F20EB4               priveda/fixed64/[is_negative.go]
// -----------------------------------------------------------------------------

package fixed64

// IsNegative returns true if the number is negative.
func (n Fixed64) IsNegative() bool {
	return n.i64 < 0
}

//end
