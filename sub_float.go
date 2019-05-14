// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:41:01 9CFB30                 priveda/fixed64/[sub_float.go]
// -----------------------------------------------------------------------------

package fixed64

// SubFloat subtracts one or more floating-point numbers from a fixed-point
// number and returns the result. The original number is not changed.
func (ob Fixed64) SubFloat(subtract ...float64) Fixed64 {
	for _, b := range subtract {
		ob.i64 -= int64(b * 1E4)
	}
	return ob
}

//end
