// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 9565BB                       priveda/fixed64/[raw.go]
// -----------------------------------------------------------------------------

package fixed64

// Raw returns the raw int64 used to store the currency value
func (ob Fixed64) Raw() int64 {
	return ob.val
}

//end
