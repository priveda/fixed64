// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 19:00:28 482AE0               priveda/fixed64/[is_negative.go]
// -----------------------------------------------------------------------------

package fixed64

// IsNegative returns true if the number is negative.
func (n Fixed64) IsNegative() bool {
	return n.i64 < 0 && n.i64 != NaN
}

//end
