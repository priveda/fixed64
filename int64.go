// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 92743A                     priveda/fixed64/[int64.go]
// -----------------------------------------------------------------------------

package fixed64

// Int64 returns the fixed-point number as an int64 value.
func (n Fixed64) Int64() int64 {
	return n.i64 / 1E4
}

//end
