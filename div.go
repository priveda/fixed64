// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 F879A1                       priveda/fixed64/[div.go]
// -----------------------------------------------------------------------------

package fixed64

// Div divides a fixed-point number by one or more other fixed-point
// numbers and returns the result. The original number is not changed.
func (ob Fixed64) Div(divide ...Fixed64) Fixed64 {
	for _, n := range divide {
		ob.i64 *= 1E4
		ob.i64 /= n.i64
	}
	return ob
}

//end
