// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:27:00 3C0B92                 priveda/fixed64/[constants.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"math/big"
)

const (
	// IntLimit specifies the highest (and lowest
	// when negative) integer component that currency
	// can hold, about 922.33 trillion.
	// The exact number is 922(t)337(b)203(m)685,476.
	//
	// This limit exists because an int64, on which Fixed64 is based, can hold
	// up to 9,223,372,036,854,775,807. Out of this 4 digits are used for the
	// decimal part, i.e. 922,337,203,685,477.5807. The limit is set to this
	// number minus 1, so that all decimals from .0000 to .9999. can be used.
	IntLimit = 922337203685476

	// minInt64 is the lowest int64 value for a Fixed64.
	// This number is close to math.minInt64, but ensures it
	// can store 4 full decimals when scaled (the 9999 part).
	minInt64 int64 = -9223372036854769999

	// maxInt64 is the highest int64 value for a Fixed64.
	// This number is close to math.MaxInt64, but ensures it
	// can store 4 full decimals when scaled (the 9999 part).
	maxInt64 int64 = 9223372036854769999
)

var (
	// big1E4 scales the int64 to provide 4 decimal places.
	big1E4 = big.NewInt(1E4)
)

//end
