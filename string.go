// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 17:45:40 6255C7                    priveda/fixed64/[string.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"strconv"
	"strings"
)

// String returns a string representing the fixed-point number
// with up to 4 decimals and implements the Stringer interface.
func (n Fixed64) String() string {
	var (
		i    = n.i64 / 1E4              // integer value
		d    = n.i64 - i*1E4            // decimal value
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
