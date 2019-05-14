// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:24:20 416202                 priveda/fixed64/[add_int64.go]
// -----------------------------------------------------------------------------

package fixed64

// AddInt64 adds one or more 64-bit integers to a fixed-point number
// and returns the result. The original number is not changed.
func (n Fixed64) AddInt64(nums ...int64) Fixed64 {
	for _, b := range nums {
		n = n.Add(Fixed64{b * 1E4})
	}
	return n
}

//end
