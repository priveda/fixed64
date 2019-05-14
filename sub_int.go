// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 A1F60D                   priveda/fixed64/[sub_int.go]
// -----------------------------------------------------------------------------

package fixed64

// SubInt subtracts one or more integer values from a currency object
// and returns the result. The object's value isn't changed.
func (ob Fixed64) SubInt(subtract ...int) Fixed64 {
	for _, n := range subtract {
		ob.val -= int64(n) * 1E4
	}
	return ob
}

//end
