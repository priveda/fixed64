// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 6AAE8B                 priveda/fixed64/[div_float.go]
// -----------------------------------------------------------------------------

package fixed64

// DivFloat divides a fixed-point number by one or more floating-point
// numbers and returns the result. The original number is not changed.
func (n Fixed64) DivFloat(divide ...float64) Fixed64 {
	for _, b := range divide {
		n.i64 *= 1E4
		n.i64 /= int64(b * 1E4)
	}
	return n
}

//end
