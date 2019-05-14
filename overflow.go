// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 E9F781                  priveda/fixed64/[overflow.go]
// -----------------------------------------------------------------------------

package fixed64

// Overflow returns 1 if the number contains a positive overflow,
// -1 if it contains a negative overflow, or zero if there is no overflow or
// underflow. Overflow and underflows occur when an arithmeric operation's
// result exceeds the storage capacity of Fixed64.
func (ob Fixed64) Overflow() int {
	if ob.i64 > maxInt64 {
		return 1
	}
	if ob.i64 < minInt64 {
		return -1
	}
	return 0
}

//end
