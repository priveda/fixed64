// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 06BF22                   priveda/fixed64/[float64.go]
// -----------------------------------------------------------------------------

package fixed64

// Float64 returns the fixed-point number as a float64 value.
func (n Fixed64) Float64() float64 {
	return float64(n.i64) / 1E4
}

//end
