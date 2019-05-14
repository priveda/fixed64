// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 97ECB8                priveda/fixed64/[is_greater.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreater returns true if the number is greater than value.
func (ob Fixed64) IsGreater(val interface{}) bool {
	return ob.i64 > Fixed64Of(val).i64
}

//end
