// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                    priveda/fixed64/[sub_int.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// SubInt subtracts one or more integer values from a fixed-point number
// and returns the result. The original number is not changed.
func (n Fixed64) SubInt(nums ...int) Fixed64 {
	for _, num := range nums {
		n.i64 -= int64(num) * 1e4
	}
	return n
}

//end
