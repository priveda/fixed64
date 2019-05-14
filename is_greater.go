// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 41BB3A                priveda/fixed64/[is_greater.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreater returns true if the object is greater than val.
func (ob Fixed64) IsGreater(val interface{}) bool {
	return ob.i64 > Fixed64Of(val).i64
}

//end
