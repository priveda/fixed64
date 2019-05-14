// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 7DE533                   priveda/fixed64/[float64.go]
// -----------------------------------------------------------------------------

package fixed64

// Float64 returns the currency value as a float64 value.
func (ob Fixed64) Float64() float64 {
	return float64(ob.val) / 1E4
}

//end
