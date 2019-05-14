// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 1AD5CE                  priveda/fixed64/[is_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsEqual returns true if the number is equal to value.
func (n Fixed64) IsEqual(value interface{}) bool {
	return n.i64 == Fixed64Of(value).i64
}

//end
