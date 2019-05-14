// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 A879C3                       priveda/fixed64/[add.go]
// -----------------------------------------------------------------------------

package fixed64

// Add adds one or more currency values and returns a new Fixed64 object.
// The value in the object to which this method is applied isn't changed.
// If there is an overflow, sets the Fixed64's internal value to
// math.MinInt64 or math.MaxInt64 depending on if the result is negative.
func (ob Fixed64) Add(add ...Fixed64) Fixed64 {
	for _, itm := range add {
		var (
			a = ob.val
			b = itm.val
			c = a + b
		)
		// check for overflow
		if c < MinFixed64I64 || (a < 0 && b < 0 && b < (MinFixed64I64-a)) {
			return currencyOverflow(true, EOverflow, ": ", a, " + ", b)
		}
		if c > MaxFixed64I64 || (a > 0 && b > 0 && b > (MaxFixed64I64-a)) {
			return currencyOverflow(false, EOverflow, ": ", a, " + ", b)
		}
		ob.val = c
	}
	return ob
}

//end
