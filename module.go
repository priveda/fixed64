// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:27:00 FC6A2D                    priveda/fixed64/[module.go]
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
//
// # String Output:
//   (ob Fixed64) GoString() string
//   (ob Fixed64) Fmt(decimalPlaces int) string
//   (ob Fixed64) InWordsEN(fmt string) string
//   (ob Fixed64) String() string
//
// # Division:
//   (ob Fixed64) Div(divide ...Fixed64) Fixed64
//   (ob Fixed64) DivFloat(divide ...float64) Fixed64
//   (ob Fixed64) DivInt(divide ...int) Fixed64
//
// # Multiplication:
//   (ob Fixed64) Mul(multiply ...Fixed64) Fixed64
//   (ob Fixed64) MulFloat(multiply ...float64) Fixed64
//   (ob Fixed64) MulInt(multiply ...int) Fixed64
//
// # Addition:
//   (ob Fixed64) Add(add ...Fixed64) Fixed64
//   (ob Fixed64) AddFloat(add ...float64) Fixed64
//   (ob Fixed64) AddInt(add ...int) Fixed64
//
// # Subtraction:
//   (ob Fixed64) Sub(subtract ...Fixed64) Fixed64
//   (ob Fixed64) SubFloat(subtract ...float64) Fixed64
//   (ob Fixed64) SubInt(subtract ...int) Fixed64
//
// # Information:
//   (ob Fixed64) Float64() float64
//   (ob Fixed64) Int() int
//   (ob Fixed64) Int64() int64
//   (ob Fixed64) IsEqual(val interface{}) bool
//   (ob Fixed64) IsGreater(val interface{}) bool
//   (ob Fixed64) IsGreaterOrEqual(val interface{}) bool
//   (ob Fixed64) IsLesser(val interface{}) bool
//   (ob Fixed64) IsLesserOrEqual(val interface{}) bool
//   (ob Fixed64) IsNegative() bool
//   (ob Fixed64) IsZero() bool
//   (ob Fixed64) Overflow() int
//   (ob Fixed64) Raw() int64
//
// # JSON:
//   (ob Fixed64) MarshalJSON() ([]byte, error)
//   (ob *Fixed64) UnmarshalJSON(data []byte) error
//
// # Helper Function
//   currencyOverflow(isNegative bool, a ...interface{}) Fixed64

// -----------------------------------------------------------------------------

const (
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
