// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2020-08-06 23:34:16 A3AFDA                     priveda/fixed64/[int64.go]
// -----------------------------------------------------------------------------

package fixed64

// Int64 returns the fixed-point number as an int64 value.
func (n Fixed64) Int64() int64 {
	return n.i64 / 1e4
}

//end
