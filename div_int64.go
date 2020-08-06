// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2020-08-06 23:34:16 156168                 priveda/fixed64/[div_int64.go]
// -----------------------------------------------------------------------------

package fixed64

// DivInt64 divides a fixed-point number by one or more 64-bit integers
// and returns the result. The original number is not changed.
func (n Fixed64) DivInt64(nums ...int64) Fixed64 {
	for _, num := range nums {
		n.i64 *= 1e4
		n.i64 /= (num * 1e4)
	}
	return n
}

//end
