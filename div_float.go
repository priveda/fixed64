// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 7E047F                 priveda/fixed64/[div_float.go]
// -----------------------------------------------------------------------------

package fixed64

// DivFloat divides a currency object by one or more floating-point
// numbers and returns the result. The object's value isn't changed.
func (ob Fixed64) DivFloat(divide ...float64) Fixed64 {
	for _, n := range divide {
		ob.val *= 1E4
		ob.val /= int64(n * 1E4)
	}
	return ob
}

//end
