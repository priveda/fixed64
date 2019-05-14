// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 536FB7                   priveda/fixed64/[sub_int.go]
// -----------------------------------------------------------------------------

package fixed64

// SubInt subtracts one or more integer values from a fixed-point number
// and returns the result. The original number is not changed.
func (n Fixed64) SubInt(subtract ...int) Fixed64 {
	for _, b := range subtract {
		n.i64 -= int64(b) * 1E4
	}
	return n
}

//end
