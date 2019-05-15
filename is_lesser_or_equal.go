// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-15 16:28:01 9235F0        priveda/fixed64/[is_lesser_or_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsLesserOrEqual returns true if the number is lesser or equal to value.
func (n Fixed64) IsLesserOrEqual(value interface{}) bool {
	return n.i64 <= New(value).i64
}

//end
