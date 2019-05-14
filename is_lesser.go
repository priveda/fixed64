// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 4D5B6B                 priveda/fixed64/[is_lesser.go]
// -----------------------------------------------------------------------------

package fixed64

// IsLesser returns true if the number is lesser than value.
func (ob Fixed64) IsLesser(val interface{}) bool {
	return ob.i64 < Fixed64Of(val).i64
}

//end
