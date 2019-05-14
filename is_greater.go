// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 4E0D09                priveda/fixed64/[is_greater.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreater returns true if the number is greater than value.
func (n Fixed64) IsGreater(val interface{}) bool {
	return n.i64 > Fixed64Of(val).i64
}

//end
