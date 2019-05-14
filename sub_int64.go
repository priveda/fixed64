// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:08 6EF566                 priveda/fixed64/[sub_int64.go]
// -----------------------------------------------------------------------------

package fixed64

// SubInt64 subtracts one or more 64-bit integers from a fixed-point
// number and returns the result. The original number is not changed.
func (n Fixed64) SubInt64(nums ...int64) Fixed64 {
	for _, num := range nums {
		n.i64 -= num * 1E4
	}
	return n
}

//end
