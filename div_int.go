// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 8488CF                   priveda/fixed64/[div_int.go]
// -----------------------------------------------------------------------------

package fixed64

// DivInt divides a currency object by one or more integer values
// and returns the result. The object's value isn't changed.
func (ob Fixed64) DivInt(divide ...int) Fixed64 {
	for _, val := range divide {
		ob.val *= 1E4
		ob.val /= (int64(val) * 1E4)
	}
	return ob
}

//end
