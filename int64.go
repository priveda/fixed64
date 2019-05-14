// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 1653FA                     priveda/fixed64/[int64.go]
// -----------------------------------------------------------------------------

package fixed64

// Int64 returns the currency value as an int64 value.
func (ob Fixed64) Int64() int64 {
	return ob.i64 / 1E4
}

//end
