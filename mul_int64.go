// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                  priveda/fixed64/[mul_int64.go]
// (c) admin@priveda.com                                            License: MIT
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

// end
