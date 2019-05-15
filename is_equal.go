// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-15 16:28:01 47C270                  priveda/fixed64/[is_equal.go]
// -----------------------------------------------------------------------------

package fixed64

// IsEqual returns true if the number is equal to value.
func (n Fixed64) IsEqual(value interface{}) bool {
	return n.i64 == New(value).i64
}

//end
