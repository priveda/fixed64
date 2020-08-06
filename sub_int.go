// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2020-08-06 23:34:16 BCAF8F                   priveda/fixed64/[sub_int.go]
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
