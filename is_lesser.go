// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 3C8AAA                 priveda/fixed64/[is_lesser.go]
// -----------------------------------------------------------------------------

package fixed64

// IsLesser returns true if the object is lesser than val.
func (ob Fixed64) IsLesser(val interface{}) bool {
	return ob.i64 < Fixed64Of(val).i64
}

//end
