// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-15 16:28:01 6D8CE5                priveda/fixed64/[is_greater.go]
// -----------------------------------------------------------------------------

package fixed64

// IsGreater returns true if the number is greater than value.
func (n Fixed64) IsGreater(value interface{}) bool {
	return n.i64 > New(value).i64
}

//end
