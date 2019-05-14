// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 B11E8D                 priveda/fixed64/[sub_float.go]
// -----------------------------------------------------------------------------

package fixed64

// SubFloat subtracts one or more floating-point numbers from a fixed-point
// number and returns the result. The original number is not changed.
func (ob Fixed64) SubFloat(subtract ...float64) Fixed64 {
	for _, n := range subtract {
		ob.i64 -= int64(n * 1E4)
	}
	return ob
}

//end
