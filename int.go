// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 78EFCC                       priveda/fixed64/[int.go]
// -----------------------------------------------------------------------------

package fixed64

// Int returns the currency value as an int value.
func (ob Fixed64) Int() int64 {
	return ob.i64 / 1E4
}

//end
