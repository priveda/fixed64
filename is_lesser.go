// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 333EEF                 priveda/fixed64/[is_lesser.go]
// -----------------------------------------------------------------------------

package fixed64

// IsLesser returns true if the number is lesser than value.
func (n Fixed64) IsLesser(val interface{}) bool {
	return n.i64 < Fixed64Of(val).i64
}

//end
