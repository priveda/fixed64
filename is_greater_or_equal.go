// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 80B6DD       priveda/fixed64/[is_greater_or_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreaterOrEqual returns true if the number is greater or equal to value.
func (ob Fixed64) IsGreaterOrEqual(val interface{}) bool {
	return ob.i64 >= Fixed64Of(val).i64
}

//end
