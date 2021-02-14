// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                  priveda/fixed64/[is_lesser.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// IsLesser returns true if the number is lesser than value.
func (n Fixed64) IsLesser(value interface{}) bool {
	return n.i64 < New(value).i64
}

//end
