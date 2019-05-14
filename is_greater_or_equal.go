// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 A06214       priveda/fixed64/[is_greater_or_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreaterOrEqual returns true if the number is greater or equal to value.
func (n Fixed64) IsGreaterOrEqual(val interface{}) bool {
	return n.i64 >= Fixed64Of(val).i64
}

//end
