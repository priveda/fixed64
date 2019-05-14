// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 A6B06A                 priveda/fixed64/[go_string.go]
// -----------------------------------------------------------------------------

package fixed64

// GoString outputs the value as a Go language string,
// It implements the fmt.GoStringer interface.
func (ob Fixed64) GoString() string {
	return "zr.Fixed64Of(" + ob.String() + ")"
}

//end
