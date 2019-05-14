// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 7F4570                 priveda/fixed64/[is_lesser.go]
// -----------------------------------------------------------------------------

package fixed64

// IsLesser returns true if the object is lesser than val.
func (ob Fixed64) IsLesser(val interface{}) bool {
	return ob.val < Fixed64Of(val).val
}

//end
