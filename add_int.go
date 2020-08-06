// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2020-08-06 23:34:16 7FF363                   priveda/fixed64/[add_int.go]
// -----------------------------------------------------------------------------

package fixed64

// AddInt adds one or more integer values to a fixed-point number
// and returns the result. The original number is not changed.
func (n Fixed64) AddInt(nums ...int) Fixed64 {
	for _, num := range nums {
		n = n.Add(Fixed64{int64(num) * 1e4})
	}
	return n
}

//end
