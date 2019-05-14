// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 029239                  priveda/fixed64/[is_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsEqual returns true if the number is equal to value.
func (n Fixed64) IsEqual(val interface{}) bool {
	return n.i64 == Fixed64Of(val).i64
}

//end
