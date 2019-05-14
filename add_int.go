// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:30:23 BD2B7F                   priveda/fixed64/[add_int.go]
// -----------------------------------------------------------------------------

package fixed64

// AddInt adds one or more integer values to a fixed-point number
// and returns the result. The original number is not changed.
func (ob Fixed64) AddInt(add ...int) Fixed64 {
	for _, n := range add {
		ob = ob.Add(Fixed64{int64(n) * 1E4})
	}
	return ob
}

//end
