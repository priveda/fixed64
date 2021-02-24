// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                  priveda/fixed64/[go_string.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// GoString outputs the value as a Go language string,
// It implements the fmt.GoStringer interface.
func (n Fixed64) GoString() string {
	return "fixed64.New(" + n.String() + ")"
}

// end
