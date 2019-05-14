// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 BCA5F3        priveda/fixed64/[is_lesser_or_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsLesserOrEqual returns true if the number is lesser or equal to value.
func (n Fixed64) IsLesserOrEqual(value interface{}) bool {
	return n.i64 <= Fixed64Of(value).i64
}

//end
