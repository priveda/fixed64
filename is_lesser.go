// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-15 16:28:01 3C2960                 priveda/fixed64/[is_lesser.go]
// -----------------------------------------------------------------------------

package fixed64

// IsLesser returns true if the number is lesser than value.
func (n Fixed64) IsLesser(value interface{}) bool {
	return n.i64 < New(value).i64
}

//end
