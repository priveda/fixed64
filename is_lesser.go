// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 FA010A                 priveda/fixed64/[is_lesser.go]
// -----------------------------------------------------------------------------

package fixed64

// IsLesser returns true if the number is lesser than value.
func (n Fixed64) IsLesser(value interface{}) bool {
	return n.i64 < Fixed64Of(value).i64
}

//end
