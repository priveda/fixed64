// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-15 16:28:01 5297B3                 priveda/fixed64/[go_string.go]
// -----------------------------------------------------------------------------

package fixed64

// GoString outputs the value as a Go language string,
// It implements the fmt.GoStringer interface.
func (n Fixed64) GoString() string {
	return "fixed64.New(" + n.String() + ")"
}

//end
