// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 5816ED               priveda/fixed64/[fixed64_raw.go]
// -----------------------------------------------------------------------------

package fixed64

// Fixed64Raw initializes a currency value from a scaled 64-bit integer.
// The decimal point is moved left 4 decimal places.
// For example, a raw value of 15500 results in a currency value of 1.55
func Fixed64Raw(raw int64) Fixed64 {
	return Fixed64{i64: raw}
}

//end
