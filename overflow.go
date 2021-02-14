// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                   priveda/fixed64/[overflow.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// Overflow returns 1 if the number contains a positive overflow,
// -1 if it contains a negative overflow,
// or 0 if there is no overflow.
//
// Overflows occur when an arithmeric operation's
// result exceeds the storage capacity of Fixed64.
//
func (n Fixed64) Overflow() int {
	if n.i64 > maxInt64 {
		return 1
	}
	if n.i64 < minInt64 {
		return -1
	}
	return 0
}

//end
