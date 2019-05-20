// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-20 01:32:42 F5562D                       priveda/fixed64/[sub.go]
// -----------------------------------------------------------------------------

package fixed64

// Sub subtracts one or more fixed-point numbers from a fixed-point
// number and returns the result. The original number is not changed.
func (n Fixed64) Sub(nums ...Fixed64) Fixed64 {
	for _, num := range nums {
		var (
			a = n.i64
			b = num.i64
			c = a - b
		)
		// check for overflow
		if c < minInt64 || (a < 0 && b > 0 && b > (-minInt64+a)) {
			return fixed64Overflow(true, a, " - ", b)
		}
		if c > maxInt64 || (a > 0 && b < 0 && b < (-maxInt64+a)) {
			return fixed64Overflow(false, a, " - ", b)
		}
		n.i64 = c
	}
	return n
}

//end
