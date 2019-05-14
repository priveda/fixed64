// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 7C99AF                       priveda/fixed64/[raw.go]
// -----------------------------------------------------------------------------

package fixed64

// Raw returns the raw int64 used to store the currency value
func (ob Fixed64) Raw() int64 {
	return ob.i64
}

//end
