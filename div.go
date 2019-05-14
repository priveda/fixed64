// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 E37B4E                       priveda/fixed64/[div.go]
// -----------------------------------------------------------------------------

package fixed64

// Div divides a currency object by one or more currency values
// and returns the result. The object's value isn't changed.
func (ob Fixed64) Div(divide ...Fixed64) Fixed64 {
	for _, n := range divide {
		ob.val *= 1E4
		ob.val /= n.val
	}
	return ob
}

//end
