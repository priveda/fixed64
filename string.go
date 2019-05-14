// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:42:36 354E9A                    priveda/fixed64/[string.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"strconv"
	"strings"
)

// String returns a string representing the currency value
// and implements the Stringer Interface.
func (ob Fixed64) String() string {
	var (
		i    = ob.i64 / 1E4             // integer value
		d    = ob.i64 - i*1E4           // decimal value
		sint = strconv.FormatInt(i, 10) // integer part
		sdec string                     // decimal part
	)
	if d != 0 { // format decimals
		if d < 0 { // adjust for negative value
			d = -d
			if i == 0 {
				sint = "-" + sint
			}
		}
		sdec = "." + strings.TrimRight(
			strconv.FormatInt(d+1E4, 10)[1:],
			"0",
		)
	}
	return sint + sdec
}

//end
