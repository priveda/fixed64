// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 B74CA4                     priveda/fixed64/[int64.go]
// -----------------------------------------------------------------------------

package fixed64

// Int64 returns the fixed-point number as an int64 value.
func (ob Fixed64) Int64() int64 {
	return ob.i64 / 1E4
}

//end
