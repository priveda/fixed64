// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:41:00 1B8BB3                   priveda/fixed64/[add_int.go]
// -----------------------------------------------------------------------------

package fixed64

// AddInt adds one or more integer values to a fixed-point number
// and returns the result. The original number is not changed.
func (ob Fixed64) AddInt(add ...int) Fixed64 {
	for _, b := range add {
		ob = ob.Add(Fixed64{int64(b) * 1E4})
	}
	return ob
}

//end
