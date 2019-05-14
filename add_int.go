// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 E59B1F                   priveda/fixed64/[add_int.go]
// -----------------------------------------------------------------------------

package fixed64

// AddInt adds one or more integer values to a fixed-point number
// and returns the result. The original number is not changed.
func (n Fixed64) AddInt(add ...int) Fixed64 {
	for _, b := range add {
		n = n.Add(Fixed64{int64(b) * 1E4})
	}
	return n
}

//end
