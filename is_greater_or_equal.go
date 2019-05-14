// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 189DF2       priveda/fixed64/[is_greater_or_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreaterOrEqual returns true if the object is greater or equal to val.
func (ob Fixed64) IsGreaterOrEqual(val interface{}) bool {
	return ob.i64 >= Fixed64Of(val).i64
}

//end
