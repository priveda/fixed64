// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 0E7434                   priveda/fixed64/[mul_int.go]
// -----------------------------------------------------------------------------

package fixed64

// MulInt multiplies a fixed-point number by one or more integer values
// and returns the result. The original number is not changed.
func (n Fixed64) MulInt(multiply ...int) Fixed64 {
	for _, b := range multiply {
		n = n.Mul(Fixed64{int64(b * 1E4)})
	}
	return n
}

//end
