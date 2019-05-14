// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 31856C                   priveda/fixed64/[mul_int.go]
// -----------------------------------------------------------------------------

package fixed64

// MulInt multiplies a fixed-point number by one or more integer values
// and returns the result. The original number is not changed.
func (ob Fixed64) MulInt(multiply ...int) Fixed64 {
	for _, n := range multiply {
		ob = ob.Mul(Fixed64{int64(n * 1E4)})
	}
	return ob
}

//end
