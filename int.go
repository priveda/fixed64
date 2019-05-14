// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 F1D09F                       priveda/fixed64/[int.go]
// -----------------------------------------------------------------------------

package fixed64

// Int returns the fixed-point number as an int value.
func (ob Fixed64) Int() int64 {
	return ob.i64 / 1E4
}

//end
