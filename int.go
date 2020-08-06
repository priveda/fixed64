// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2020-08-06 23:34:16 D24ACD                       priveda/fixed64/[int.go]
// -----------------------------------------------------------------------------

package fixed64

// Int returns the fixed-point number as an int value.
func (n Fixed64) Int() int64 {
	return n.i64 / 1e4
}

//end
