// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 7284F4       priveda/fixed64/[is_greater_or_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreaterOrEqual returns true if the number is greater or equal to value.
func (n Fixed64) IsGreaterOrEqual(value interface{}) bool {
	return n.i64 >= Fixed64Of(value).i64
}

//end
