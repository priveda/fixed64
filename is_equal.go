// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 8E7BC8                  priveda/fixed64/[is_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsEqual returns true if the value of the currency object is negative.
func (ob Fixed64) IsEqual(val interface{}) bool {
	return ob.i64 == Fixed64Of(val).i64
}

//end
