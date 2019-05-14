// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 E3912D                       priveda/fixed64/[int.go]
// -----------------------------------------------------------------------------

package fixed64

// Int returns the fixed-point number as an int value.
func (n Fixed64) Int() int64 {
	return n.i64 / 1E4
}

//end
