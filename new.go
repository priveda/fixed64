// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-20 01:48:26 51BC17                       priveda/fixed64/[new.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"strings"
)

// New converts any compatible value type to Fixed64. This includes simple
// numeric types and strings, as well as pointers to these types.
//
// - Dereferences pointers to evaluate the pointed-to type.
// - Converts nil to Fixed64{NaN}.
// - Converts signed and unsigned integers, and floats to Fixed64.
// - Converts numeric strings to Fixed64.
//   If a string is zero-length or whitespace, returns Fixed64{0}
//   If a string is not numeric, returns Fixed64{NaN}
// - Converts boolean true to 1, false to 0.
//
// If the value can not be converted to Fixed64, returns
// Fixed64{NaN} and logs an error (when logging is active).
//
// When given a nil pointer, also returns Fixed64{NaN}.
//
// Note: fmt.Stringer (or fmt.GoStringer) interfaces are not treated as
// strings to avoid bugs from implicit conversion. Use the String method.
//
func New(value interface{}) Fixed64 {
	n, err := NewE(value)
	if err != nil && !strings.HasPrefix(err.Error(), EOverflow) {
		mod.Error(err)
	}
	return n
}

//end
