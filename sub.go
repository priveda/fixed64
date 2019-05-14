// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:56:01 42F23F                       priveda/fixed64/[sub.go]
// -----------------------------------------------------------------------------

package fixed64

// Sub subtracts one or more currency values from a currency object
// and returns the result. The object's value isn't changed.
func (ob Fixed64) Sub(subtract ...Fixed64) Fixed64 {
	for _, n := range subtract {
		var (
			a = ob.i64
			b = n.i64
			c = a - b
		)
		// check for overflow
		if c < minInt64 || (a < 0 && b > 0 && b > (-minInt64+a)) {
			return fixed64Overflow(true, EOverflow, ": ", a, " - ", b)
		}
		if c > maxInt64 || (a > 0 && b < 0 && b < (-maxInt64+a)) {
			return fixed64Overflow(false, EOverflow, ": ", a, " - ", b)
		}
		ob.i64 = c
	}
	return ob
}

//end
