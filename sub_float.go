// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                  priveda/fixed64/[sub_float.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// SubFloat subtracts one or more floating-point numbers from a fixed-point
// number and returns the result. The original number is not changed.
func (n Fixed64) SubFloat(nums ...float64) Fixed64 {
	for _, num := range nums {
		n.i64 -= int64(num * 1e4)
	}
	return n
}

//end
