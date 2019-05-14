// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:08 8BEB0C                 priveda/fixed64/[sub_float.go]
// -----------------------------------------------------------------------------

package fixed64

// SubFloat subtracts one or more floating-point numbers from a fixed-point
// number and returns the result. The original number is not changed.
func (n Fixed64) SubFloat(nums ...float64) Fixed64 {
	for _, num := range nums {
		n.i64 -= int64(num * 1E4)
	}
	return n
}

//end
