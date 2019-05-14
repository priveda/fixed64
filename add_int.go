// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 96A575                   priveda/fixed64/[add_int.go]
// -----------------------------------------------------------------------------

package fixed64

// AddInt adds one or more integer values to a fixed-point number
// and returns the result. The original number is not changed.
func (n Fixed64) AddInt(nums ...int) Fixed64 {
	for _, num := range nums {
		n = n.Add(Fixed64{int64(num) * 1E4})
	}
	return n
}

//end
