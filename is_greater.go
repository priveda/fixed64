// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 0274E0                priveda/fixed64/[is_greater.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreater returns true if the object is greater than val.
func (ob Fixed64) IsGreater(val interface{}) bool {
	return ob.val > Fixed64Of(val).val
}

//end
