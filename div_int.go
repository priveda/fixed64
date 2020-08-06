// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2020-08-06 23:34:16 0BC1AF                   priveda/fixed64/[div_int.go]
// -----------------------------------------------------------------------------

package fixed64

// DivInt divides a fixed-point number by one or more integer values
// and returns the result. The original number is not changed.
func (n Fixed64) DivInt(nums ...int) Fixed64 {
	for _, num := range nums {
		n.i64 *= 1e4
		n.i64 /= (int64(num) * 1e4)
	}
	return n
}

//end
