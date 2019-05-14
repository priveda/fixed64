// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 5C12A0                 priveda/fixed64/[constants.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"math/big"
)

const (
	// Fixed64IntLimit specifies the highest (and lowest
	// when negative) integer component that currency
	// can hold, about 922.33 trillion.
	// The exact number is 922(t)337(b)203(m)685,476.
	//
	// This limit exists because an int64, on which Fixed64 is based, can hold
	// up to 9,223,372,036,854,775,807. Out of this 4 digits are used for the
	// decimal part, i.e. 922,337,203,685,477.5807. The limit is set to this
	// number minus 1, so that all decimals from .0000 to .9999. can be used.
	Fixed64IntLimit = 922337203685476

	// MinFixed64I64 is the lowest internal value that Fixed64 can hold.
	MinFixed64I64 = -9223372036854769999

	// MaxFixed64I64 is the highest internal value that Fixed64 can hold.
	MaxFixed64I64 = 9223372036854769999
)

var (
	// big1E4 scales the int64 to provide 4 decimal places.
	big1E4 = big.NewInt(1E4)
)

//end
