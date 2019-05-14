// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 8E0674                priveda/fixed64/[is_greater.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreater returns true if the number is greater than value.
func (n Fixed64) IsGreater(value interface{}) bool {
	return n.i64 > Fixed64Of(value).i64
}

//end
