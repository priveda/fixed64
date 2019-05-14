// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 A99FAE                   priveda/fixed64/[float64.go]
// -----------------------------------------------------------------------------

package fixed64

// Float64 returns the fixed-point number as a float64 value.
func (ob Fixed64) Float64() float64 {
	return float64(ob.i64) / 1E4
}

//end
