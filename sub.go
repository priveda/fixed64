// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 E229A8                       priveda/fixed64/[sub.go]
// -----------------------------------------------------------------------------

package fixed64

// Sub subtracts one or more currency values from a currency object
// and returns the result. The object's value isn't changed.
func (ob Fixed64) Sub(subtract ...Fixed64) Fixed64 {
	for _, n := range subtract {
		var (
			a = ob.val
			b = n.val
			c = a - b
		)
		// check for overflow
		if c < MinFixed64I64 || (a < 0 && b > 0 && b > (-MinFixed64I64+a)) {
			return currencyOverflow(true, EOverflow, ": ", a, " - ", b)
		}
		if c > MaxFixed64I64 || (a > 0 && b < 0 && b < (-MaxFixed64I64+a)) {
			return currencyOverflow(false, EOverflow, ": ", a, " - ", b)
		}
		ob.val = c
	}
	return ob
}

//end
