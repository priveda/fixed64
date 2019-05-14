// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 2E6DA2                 priveda/fixed64/[add_int64.go]
// -----------------------------------------------------------------------------

package fixed64

// AddInt64 adds one or more 64-bit integers to a fixed-point number
// and returns the result. The original number is not changed.
func (n Fixed64) AddInt64(nums ...int64) Fixed64 {
	for _, num := range nums {
		n = n.Add(Fixed64{num * 1E4})
	}
	return n
}

//end
