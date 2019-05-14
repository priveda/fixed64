// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 90470B                  priveda/fixed64/[is_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsEqual returns true if the value of the currency object is negative.
func (ob Fixed64) IsEqual(val interface{}) bool {
	return ob.val == Fixed64Of(val).val
}

//end
