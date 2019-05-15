// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-15 16:28:01 494BD2       priveda/fixed64/[is_greater_or_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreaterOrEqual returns true if the number is greater or equal to value.
func (n Fixed64) IsGreaterOrEqual(value interface{}) bool {
	return n.i64 >= New(value).i64
}

//end
