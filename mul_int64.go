// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2020-08-06 23:34:16 CCE86E                 priveda/fixed64/[mul_int64.go]
// -----------------------------------------------------------------------------

package fixed64

// MulInt64 multiplies a fixed-point number by one or more 64-bit integers
// and returns the result. The original number is not changed.
func (n Fixed64) MulInt64(nums ...int64) Fixed64 {
	for _, num := range nums {
		n = n.Mul(Fixed64{num * 1e4})
	}
	return n
}

//end
