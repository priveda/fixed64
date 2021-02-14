// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                  priveda/fixed64/[add_int64.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// AddInt64 adds one or more 64-bit integers to a fixed-point number
// and returns the result. The original number is not changed.
func (n Fixed64) AddInt64(nums ...int64) Fixed64 {
	for _, num := range nums {
		n = n.Add(Fixed64{num * 1e4})
	}
	return n
}

// end
