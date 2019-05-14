// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 5D8A3A                 priveda/fixed64/[go_string.go]
// -----------------------------------------------------------------------------

package fixed64

// GoString outputs the value as a Go language string,
// It implements the fmt.GoStringer interface.
func (n Fixed64) GoString() string {
	return "zr.Fixed64Of(" + n.String() + ")"
}

//end
