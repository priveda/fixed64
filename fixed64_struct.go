// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 D76D40                 priveda/fixed64/[constants.go]
// -----------------------------------------------------------------------------

package fixed64

// Fixed64 represents a currency value with up to four decimal places.
// It is stored internally as a 64-bit integer. This gives it a range
// from -922,337,203,685,477.5808 to 922,337,203,685,477.5807
type Fixed64 struct {
	i64 int64
}

//end
