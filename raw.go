// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 E99291                       priveda/fixed64/[raw.go]
// -----------------------------------------------------------------------------

package fixed64

// Raw returns the internal int64 used to store the fixed-point number.
func (n Fixed64) Raw() int64 {
	return n.i64
}

//end
