// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:17:01 97CCAA                    priveda/fixed64/[module.go]
// -----------------------------------------------------------------------------

package fixed64

// # Constants:
//   IntLimit
//   minInt64
//   maxInt64
//
// # Fixed64 Type:
//   Fixed64 struct
//
// # Fixed64 Factories:
//   Fixed64Of(value interface{}) Fixed64
//   Fixed64OfS(s string) Fixed64
//   Fixed64Raw(raw int64) Fixed64
//   ParseFixed64(s string) (Fixed64, error)
//
// # String Output:
//   (n Fixed64) GoString() string
//   (n Fixed64) Fmt(decimalPlaces int) string
//   (n Fixed64) InWordsEN(fmt string) string
//   (n Fixed64) String() string
//
// # Division:
//   (n Fixed64) Div(divide ...Fixed64) Fixed64
//   (n Fixed64) DivFloat(divide ...float64) Fixed64
//   (n Fixed64) DivInt(divide ...int) Fixed64
//
// # Multiplication:
//   (n Fixed64) Mul(multiply ...Fixed64) Fixed64
//   (n Fixed64) MulFloat(multiply ...float64) Fixed64
//   (n Fixed64) MulInt(multiply ...int) Fixed64
//
// # Addition:
//   (n Fixed64) Add(add ...Fixed64) Fixed64
//   (n Fixed64) AddFloat(add ...float64) Fixed64
//   (n Fixed64) AddInt(add ...int) Fixed64
//
// # Subtraction:
//   (n Fixed64) Sub(subtract ...Fixed64) Fixed64
//   (n Fixed64) SubFloat(subtract ...float64) Fixed64
//   (n Fixed64) SubInt(subtract ...int) Fixed64
//
// # Information:
//   (n Fixed64) Float64() float64
//   (n Fixed64) Int() int
//   (n Fixed64) Int64() int64
//   (n Fixed64) IsEqual(val interface{}) bool
//   (n Fixed64) IsGreater(val interface{}) bool
//   (n Fixed64) IsGreaterOrEqual(val interface{}) bool
//   (n Fixed64) IsLesser(val interface{}) bool
//   (n Fixed64) IsLesserOrEqual(val interface{}) bool
//   (n Fixed64) IsNegative() bool
//   (n Fixed64) IsZero() bool
//   (n Fixed64) Overflow() int
//   (n Fixed64) Raw() int64
//
// # JSON:
//   (n Fixed64) MarshalJSON() ([]byte, error)
//   (n *Fixed64) UnmarshalJSON(data []byte) error
//
// # Helper Function
//   fixed64Overflow(isNegative bool, a ...interface{}) Fixed64

// -----------------------------------------------------------------------------

const (
	// EInvalidValue error occurs when a string or another type
	// can not be parsed or converted to the required data type.
	EInvalidValue = "Invalid value"

	// EOverflow indicates an arithmetic overflow.
	EOverflow = "Overflow"

	// ENilReceiver indicates a method call on a nil object.
	ENilReceiver = "nil receiver"
)

// -----------------------------------------------------------------------------
// # Function Proxy Variables (for mocking)

type thisMod struct {
	Error func(args ...interface{}) error
}

var mod = thisMod{
	Error: logError,
}

// ModReset restores all mocked functions to the original standard functions.
func (ob *thisMod) Reset() {
	ob.Error = logError
}

// logError __
func logError(args ...interface{}) error {
	return nil
}

//end
