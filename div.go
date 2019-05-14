// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 3EDFA8                       priveda/fixed64/[div.go]
// -----------------------------------------------------------------------------

package fixed64

// Div divides a fixed-point number by one or more other fixed-point
// numbers and returns the result. The original number is not changed.
func (n Fixed64) Div(divide ...Fixed64) Fixed64 {
	for _, b := range divide {
		n.i64 *= 1E4
		n.i64 /= b.i64
	}
	return n
}

//end
