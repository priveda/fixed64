// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 BC1F8C                   priveda/fixed64/[add_int.go]
// -----------------------------------------------------------------------------

package fixed64

// AddInt adds one or more integer values to a currency object
// and returns the result. The object's value isn't changed.
func (ob Fixed64) AddInt(add ...int) Fixed64 {
	for _, n := range add {
		ob = ob.Add(Fixed64{int64(n) * 1E4})
	}
	return ob
}

//end
