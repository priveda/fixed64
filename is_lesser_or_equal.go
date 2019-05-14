// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 E58B46        priveda/fixed64/[is_lesser_or_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsLesserOrEqual returns true if the number is lesser or equal to value.
func (ob Fixed64) IsLesserOrEqual(val interface{}) bool {
	return ob.i64 <= Fixed64Of(val).i64
}

//end
