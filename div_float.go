// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 6AEE86                 priveda/fixed64/[div_float.go]
// -----------------------------------------------------------------------------

package fixed64

// DivFloat divides a fixed-point number by one or more floating-point
// numbers and returns the result. The original number is not changed.
func (ob Fixed64) DivFloat(divide ...float64) Fixed64 {
	for _, n := range divide {
		ob.i64 *= 1E4
		ob.i64 /= int64(n * 1E4)
	}
	return ob
}

//end
