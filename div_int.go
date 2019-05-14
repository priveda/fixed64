// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 182FED                   priveda/fixed64/[div_int.go]
// -----------------------------------------------------------------------------

package fixed64

// DivInt divides a currency object by one or more integer values
// and returns the result. The object's value isn't changed.
func (ob Fixed64) DivInt(divide ...int) Fixed64 {
	for _, val := range divide {
		ob.i64 *= 1E4
		ob.i64 /= (int64(val) * 1E4)
	}
	return ob
}

//end
