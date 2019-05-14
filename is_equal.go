// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 EE3024                  priveda/fixed64/[is_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsEqual returns true if the number is equal to value.
func (ob Fixed64) IsEqual(val interface{}) bool {
	return ob.i64 == Fixed64Of(val).i64
}

//end
