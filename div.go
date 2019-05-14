// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 1AA03A                       priveda/fixed64/[div.go]
// -----------------------------------------------------------------------------

package fixed64

// Div divides a currency object by one or more currency values
// and returns the result. The object's value isn't changed.
func (ob Fixed64) Div(divide ...Fixed64) Fixed64 {
	for _, n := range divide {
		ob.i64 *= 1E4
		ob.i64 /= n.i64
	}
	return ob
}

//end
