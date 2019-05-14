// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 640E26                       priveda/fixed64/[raw.go]
// -----------------------------------------------------------------------------

package fixed64

// Raw returns the internal int64 used to store the fixed-point number.
func (ob Fixed64) Raw() int64 {
	return ob.i64
}

//end
