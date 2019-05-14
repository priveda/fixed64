// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 C9C295                   priveda/fixed64/[mul_int.go]
// -----------------------------------------------------------------------------

package fixed64

// MulInt multiplies a fixed-point number by one or more integer values
// and returns the result. The original number is not changed.
func (n Fixed64) MulInt(nums ...int) Fixed64 {
	for _, num := range nums {
		n = n.Mul(Fixed64{int64(num * 1E4)})
	}
	return n
}

//end
