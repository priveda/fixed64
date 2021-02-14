// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                        priveda/fixed64/[sub.go]
// (c) admin@priveda.com                                            License: MIT
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
