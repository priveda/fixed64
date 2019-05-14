// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:41:00 CFDCF8                       priveda/fixed64/[sub.go]
// -----------------------------------------------------------------------------

package fixed64

// Sub subtracts one or more fixed-point numbers from a fixed-point
// number and returns the result. The original number is not changed.
func (ob Fixed64) Sub(subtract ...Fixed64) Fixed64 {
	for _, nb := range subtract {
		var (
			a = ob.i64
			b = nb.i64
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
