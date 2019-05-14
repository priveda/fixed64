// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 A2D78A                       priveda/fixed64/[int.go]
// -----------------------------------------------------------------------------

package fixed64

// Int returns the currency value as an int value.
func (ob Fixed64) Int() int64 {
	return ob.val / 1E4
}

//end
