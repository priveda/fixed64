// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 0D21E2        priveda/fixed64/[is_lesser_or_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsLesserOrEqual returns true if the object is lesser or equal to val.
func (ob Fixed64) IsLesserOrEqual(val interface{}) bool {
	return ob.val <= Fixed64Of(val).val
}

//end
