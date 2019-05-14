// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:24:20 0DDCC2                 priveda/fixed64/[mul_int64.go]
// -----------------------------------------------------------------------------

package fixed64

// MulInt64 multiplies a fixed-point number by one or more 64-bit integers
// and returns the result. The original number is not changed.
func (n Fixed64) MulInt64(nums ...int64) Fixed64 {
	for _, b := range nums {
		n = n.Mul(Fixed64{b * 1E4})
	}
	return n
}

//end
