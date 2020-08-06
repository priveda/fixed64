// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2020-08-06 23:34:16 3764C2                   priveda/fixed64/[float64.go]
// -----------------------------------------------------------------------------

package fixed64

// Float64 returns the fixed-point number as a float64 value.
func (n Fixed64) Float64() float64 {
	return float64(n.i64) / 1e4
}

//end
