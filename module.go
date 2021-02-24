// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number                     priveda/fixed64/[module.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

// # Fixed64 Type:
//   Fixed64 struct
//
// # Fixed64 Factories:
//   New(value interface{}) Fixed64
//   NewE(value interface{}) (Fixed64, error)
//   Parse(s string) (Fixed64, error)
//   Wrap(i64 int64) Fixed64
//
// # String Output:
//   (n Fixed64) GoString() string
//   (n Fixed64) Fmt(decimalPlaces int) string
//   (n Fixed64) InWordsEN(fmt string) string
//   (n Fixed64) String() string
//
// # Division:
//   (n Fixed64) Div(nums ...Fixed64) Fixed64
//   (n Fixed64) DivFloat(nums ...float64) Fixed64
//   (n Fixed64) DivInt(nums ...int) Fixed64
//   (n Fixed64) DivInt64(nums ...int64) Fixed64
//
// # Multiplication:
//   (n Fixed64) Mul(nums ...Fixed64) Fixed64
//   (n Fixed64) MulFloat(nums ...float64) Fixed64
//   (n Fixed64) MulInt(nums ...int) Fixed64
//   (n Fixed64) MulInt64(nums ...int64) Fixed64
//
// # Addition:
//   (n Fixed64) Add(nums ...Fixed64) Fixed64
//   (n Fixed64) AddFloat(nums ...float64) Fixed64
//   (n Fixed64) AddInt(nums ...int) Fixed64
//   (n Fixed64) AddInt64(nums ...int64) Fixed64
//
// # Subtraction:
//   (n Fixed64) Sub(nums ...Fixed64) Fixed64
//   (n Fixed64) SubFloat(nums ...float64) Fixed64
//   (n Fixed64) SubInt(nums ...int) Fixed64
//   (n Fixed64) SubInt64(nums ...int64) Fixed64
//
// # Information:
//   (n Fixed64) Float64() float64
//   (n Fixed64) Int() int
//   (n Fixed64) Int64() int64
//   (n Fixed64) IsEqual(value interface{}) bool
//   (n Fixed64) IsGreater(value interface{}) bool
//   (n Fixed64) IsGreaterOrEqual(value interface{}) bool
//   (n Fixed64) IsLesser(value interface{}) bool
//   (n Fixed64) IsLesserOrEqual(value interface{}) bool
//   (n Fixed64) IsNegative() bool
//   (n Fixed64) IsZero() bool
//   (n Fixed64) Overflow() int
//   (n Fixed64) Unwrap() int64
//
// # JSON:
//   (n Fixed64) MarshalJSON() ([]byte, error)
//   (n *Fixed64) UnmarshalJSON(data []byte) error
//
// # Helper Function
//   fixed64Overflow(isNegative bool, args ...interface{}) Fixed64

// -----------------------------------------------------------------------------
// # Function Proxy Variables (for mocking)

// mockable defines mockable functions used by this package
type mockable struct {
	Error func(args ...interface{}) error
}

// mod is the single instance of mockable
var mod = mockable{
	Error: logError,
}

// Reset restores all mocked functions to the original standard functions.
func (m *mockable) Reset() {
	m.Error = logError
}

// logError logs error conditions to a log file.
// (by default it does nothing)
func logError(args ...interface{}) error {
	return nil
}

// end
