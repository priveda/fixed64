// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 A89647                 priveda/fixed64/[sub_float.go]
// -----------------------------------------------------------------------------

package fixed64

// SubFloat subtracts one or more floating-point numbers from a currency
// object and returns the result. The object's value isn't changed.
func (ob Fixed64) SubFloat(subtract ...float64) Fixed64 {
	for _, n := range subtract {
		ob.val -= int64(n * 1E4)
	}
	return ob
}

//end
