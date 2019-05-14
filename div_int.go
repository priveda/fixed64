// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 AEFB66                   priveda/fixed64/[div_int.go]
// -----------------------------------------------------------------------------

package fixed64

// DivInt divides a fixed-point number by one or more integer values
// and returns the result. The original number is not changed.
func (n Fixed64) DivInt(divide ...int) Fixed64 {
	for _, val := range divide {
		n.i64 *= 1E4
		n.i64 /= (int64(val) * 1E4)
	}
	return n
}

//end
