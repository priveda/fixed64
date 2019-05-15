// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-15 16:46:15 8351D1                         memd/[fixed64_test.go]
// -----------------------------------------------------------------------------

package fixed64

// # Fixed64 Factories:
//   Test_New_
//   Test_Parse_
//   Test_Wrap_
//
// # String Output:
//   Test_Fixed64_Fmt_
//   Test_Fixed64_String_
//
// # Division:
//   Test_Fixed64_Div_
//   Test_Fixed64_DivFloat_
//   Test_Fixed64_DivInt_
//   Test_Fixed64_DivInt64_
//
// # Multiplication:
//   Test_Fixed64_Mul_
//   Test_Fixed64_MulFloat_
//   Test_Fixed64_MulInt_
//   Test_Fixed64_MulInt64_
//
// # Addition:
//   Test_Fixed64_Add_
//   Test_Fixed64_AddFloat_
//   Test_Fixed64_AddInt_
//   Test_Fixed64_AddInt64_
//
// # Subtraction:
//   Test_Fixed64_Sub_
//   Test_Fixed64_SubFloat_
//   Test_Fixed64_SubInt_
//   Test_Fixed64_SubInt64_
//
// # Information:
//   Test_Fixed64_Float64_
//   Test_Fixed64_Int_
//   Test_Fixed64_Int64_
//   Test_Fixed64_IsEqual_
//   Test_Fixed64_IsGreater_
//   Test_Fixed64_IsGreaterOrEqual_
//   Test_Fixed64_IsLesser_
//   Test_Fixed64_IsLesserOrEqual_
//   Test_Fixed64_IsNegative_
//   Test_Fixed64_IsZero_
//   Test_Fixed64_Overflow_
//   Test_Fixed64_Unwrap_
//
// # JSON:
//   Test_Fixed64_MarshalJSON_
//   Test_Fixed64_UnmarshalJSON_
//
// # Test Helper Functions
//   AF(nums ...float64) (ret []float64)
//   AI(nums ...int) (ret []int)
//   AN(nums ...Fixed64) (ret []Fixed64)

//  to test all items in fixed_num.go use:
//      go test --run Fixed64
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"testing"
)

// -----------------------------------------------------------------------------
// # Fixed64 Factories:

// go test --run Test_New_
func Test_New_(t *testing.T) {
	//
	// New(value interface{}) Fixed64
	//
	test := func(input interface{}, want Fixed64) {
		got := New(input)
		label := fmt.Sprintf("<-L%d: New(%#v)", testLine(), input)
		testGot(label, got.i64, want.i64, func(erm string) { t.Error(erm) })
	}
	// nil interface pointer
	{
		test(nil, Fixed64{NaN})
	}
	// highest/lowest possible fixed-point number
	test(int64(-922337203685476), Fixed64{-922337203685476 * 1E4})
	test(int64(922337203685476), Fixed64{922337203685476 * 1E4})
	//
	// strings
	test(" ", Fixed64{NaN})
	test("", Fixed64{NaN})
	test("-922337203685476", Fixed64{-922337203685476 * 1E4})
	test("1.01", Fixed64{1.01 * 1E4})
	test("922337203685476", Fixed64{922337203685476 * 1E4})
	//
	// integers
	test(int(-123456), Fixed64{-(123456 * 1E4)})
	test(int16(math.MaxInt16), Fixed64{32767 * 1E4})
	test(int16(math.MinInt16), Fixed64{-32768 * 1E4})
	test(int32(math.MaxInt32), Fixed64{2147483647 * 1E4})
	test(int32(math.MinInt32), Fixed64{-2147483648 * 1E4})
	test(int8(math.MaxInt8), Fixed64{127 * 1E4})
	test(int8(math.MinInt8), Fixed64{-128 * 1E4})
	// TODO: test with acceptable limit for int64
	test(int64(math.MaxInt32), Fixed64{2147483647 * 1E4})
	test(int64(math.MinInt32), Fixed64{-2147483648 * 1E4})
	//
	// unsigned integers
	test(uint(123456), Fixed64{123456 * 1E4})
	test(uint16(0), Fixed64{0})
	test(uint16(math.MaxUint16), Fixed64{65535 * 1E4})
	test(uint32(0), Fixed64{0})
	test(uint32(math.MaxUint32), Fixed64{4294967295 * 1E4})
	test(uint64(0), Fixed64{0})
	test(uint64(922337203685476), Fixed64{922337203685476 * 1E4})
	test(uint8(0), Fixed64{0})
	test(uint8(math.MaxUint8), Fixed64{255 * 1E4})
	//
	// floating-point numbers:
	//
	test(float32(-0.1), Fixed64{-1000})
	test(float32(0), Fixed64{0})
	test(float32(0.1), Fixed64{1000})
	test(float64(0), Fixed64{0})
	//
	test(float64(0.00001), Fixed64{0})
	test(float64(0.0001), Fixed64{1})
	test(float64(0.001), Fixed64{10})
	test(float64(0.01), Fixed64{100})
	test(float64(0.1), Fixed64{1000})
	//
	test(float64(-0.00001), Fixed64{0})
	test(float64(-0.0001), Fixed64{-1})
	test(float64(-0.001), Fixed64{-10})
	test(float64(-0.01), Fixed64{-100})
	test(float64(-0.1), Fixed64{-1000})
	//
	test(float32(1000), Fixed64{1000 * 1E4})
	test(float32(1234.5), Fixed64{1234.5 * 1E4})
	test(float64(1000), Fixed64{1000 * 1E4})
	test(float64(12345.6789), Fixed64{12345.6789 * 1E4})
	//
	// integer pointers
	{
		n := int(-123456)
		test(&n, Fixed64{-(123456 * 1E4)})
	}
	{
		n := int8(math.MaxInt8)
		test(&n, Fixed64{127 * 1E4})
	}
	{
		n := int8(math.MinInt8)
		test(&n, Fixed64{-128 * 1E4})
	}
	{
		n := int16(math.MaxInt16)
		test(&n, Fixed64{32767 * 1E4})
	}
	{
		n := int16(math.MinInt16)
		test(&n, Fixed64{-32768 * 1E4})
	}
	{
		n := int32(math.MaxInt32)
		test(&n, Fixed64{2147483647 * 1E4})
	}
	{
		n := int32(math.MinInt32)
		test(&n, Fixed64{-2147483648 * 1E4})
	}
	{
		n := int64(922337203685476)
		test(&n, Fixed64{922337203685476 * 1E4})
	}
	{
		n := int64(-922337203685476)
		test(&n, Fixed64{-922337203685476 * 1E4})
	}
	// unsigned integer pointers
	{
		n := uint(123456)
		test(&n, Fixed64{123456 * 1E4})
	}
	{
		n := uint8(math.MaxUint8)
		test(&n, Fixed64{255 * 1E4})
	}
	{
		n := uint8(0)
		test(&n, Fixed64{0})
	}
	{
		n := uint16(math.MaxUint16)
		test(&n, Fixed64{65535 * 1E4})
	}
	{
		n := uint16(0)
		test(&n, Fixed64{0})
	}
	{
		n := uint32(math.MaxUint32)
		test(&n, Fixed64{4294967295 * 1E4})
	}
	{
		n := uint32(0)
		test(&n, Fixed64{0})
	}
	// nil pointer to int
	{
		var ptr *int
		test(ptr, Fixed64{NaN})
	}
	// pointers to floating-point numbers
	{
		n := float32(1000)
		test(&n, Fixed64{1000 * 1E4})
	}
	{
		n := float32(1234.5)
		test(&n, Fixed64{1234.5 * 1E4})
	}
	{
		n := float64(1000)
		test(&n, Fixed64{1000 * 1E4})
	}
	{
		n := float64(12345.6789)
		test(&n, Fixed64{12345.6789 * 1E4})
	}
	// nil pointer to float
	{
		var ptr *float64
		test(ptr, Fixed64{NaN})
	}
	// pointer to string
	{
		s := "922337203685476"
		test(&s, Fixed64{922337203685476 * 1E4})
	}
	{
		s := "-922337203685476"
		test(&s, Fixed64{-922337203685476 * 1E4})
	}
	// nil pointer to string
	{
		var ptr *string
		test(ptr, Fixed64{NaN})
	}
	testErr := func(input interface{}, want Fixed64) {
		label := fmt.Sprintf("<-L%d: New(%#v)", testLine(), input)
		testErrors(label,
			func() error {
				got := New(input)
				testGot(label, got.i64, want.i64,
					func(erm string) { t.Error(erm) })
				return NoError
			},
			func(erm string) { t.Error(erm) },
		)
	}
	// test conditions that generate errors:
	{
		// non-numeric string
		testErr("abc", Fixed64{NaN})
	}
	{
		// wrong type
		testErr(true, Fixed64{NaN})
	}
	// overflow
	{
		// large floating-point numbers
		testErr(float32(-10e20), Fixed64{math.MinInt64 + 1})
		testErr(float32(-10e20), Fixed64{math.MinInt64 + 1})
		testErr(float64(10e20), Fixed64{math.MaxInt64})
		testErr(float64(10e20), Fixed64{math.MaxInt64})
		//
		// max int64
		testErr(int64(math.MaxInt64), Fixed64{math.MaxInt64})
		{
			n := int64(math.MaxInt64)
			testErr(&n, Fixed64{math.MaxInt64})
		}
		// min int64
		testErr(int64(math.MinInt64), Fixed64{math.MinInt64 + 1})
		{
			n := int64(math.MinInt64)
			testErr(&n, Fixed64{math.MinInt64 + 1})
		}
		// max uint64
		testErr(uint64(math.MaxUint64), Fixed64{math.MaxInt64})
		{
			n := uint64(math.MaxUint64)
			testErr(&n, Fixed64{math.MaxInt64})
		}
	}
}

// go test --run Test_Parse_
func Test_Parse_(t *testing.T) {
	//
	test := func(s string, wantI64 int64, wantError bool) {
		mockError(true)
		var (
			gotVal, gotErr = Parse(s)
		)
		gotErm, wantErm := "nil", "nil"
		if gotErr != nil {
			gotErm = "error"
		}
		if wantError {
			wantErm = "error"
		}
		var (
			got  = fmt.Sprintf("(%d, <%s>)", gotVal, gotErm)
			want = fmt.Sprintf("(%d, <%s>)", Fixed64{wantI64}, wantErm)
		)
		label := fmt.Sprintf("<-L%d: Parse(%q) =", testLine(), s)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
		mockError(false)
	}
	// spaces without any digits -> 0, ENullValue error
	test("   ", NaN, false)
	test("  ", NaN, false)
	test(" ", NaN, false)
	test("", NaN, false)
	test("\t", NaN, false)
	test("\t\t", NaN, false)
	//
	// zeros
	test(" +0 ", 0, false)
	test(" +0", 0, false)
	test(" +00 ", 0, false)
	test(" +00", 0, false)
	test(" -0 ", 0, false)
	test(" -0", 0, false)
	test(" -00 ", 0, false)
	test(" -00", 0, false)
	test(" 0 ", 0, false)
	test(" 0", 0, false)
	test(" 00 ", 0, false)
	test(" 00", 0, false)
	test("+0 ", 0, false)
	test("+0", 0, false)
	test("+00 ", 0, false)
	test("+00", 0, false)
	test("-0 ", 0, false)
	test("-0", 0, false)
	test("-00 ", 0, false)
	test("-00", 0, false)
	test("0 ", 0, false)
	test("0", 0, false)
	test("00 ", 0, false)
	test("00", 0, false)
	//
	test("0.12345", 0*1E4+1234, false)
	test("123456", 123456*1E4, false)
	test("abc", NaN, true)
}

// go test --run Test_Wrap_
func Test_Wrap_(t *testing.T) {
	//
	// Wrap(i64 int64) Fixed64
	//
	test := func(i64 int64, want Fixed64) {
		label := fmt.Sprintf("<-L%d: Wrap(%d) =", testLine(), i64)
		got := Wrap(i64)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(-1, Fixed64{-1})
	test(0, Fixed64{0})
	test(1, Fixed64{1})
	test(NaN, Fixed64{NaN})
}

// -----------------------------------------------------------------------------
// # String Output:

// go test --run Test_Fixed64_Fmt_
func Test_Fixed64_Fmt_(t *testing.T) {
	//
	// (n Fixed64) Fmt(decimalPlaces int) string
	//
	test := func(input interface{}, decimalPlaces int, want string) {
		var (
			n     = New(input)
			got   = n.Fmt(decimalPlaces)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.Fmt(%v) =",
				testLine(), n, decimalPlaces)
		)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	//
	// zeros with different number of decimal places
	test("0", 0, "0")
	test("0", 1, "0.0")
	test("0", 2, "0.00")
	test("0", 3, "0.000")
	test("0", 4, "0.0000")
	test("0", 5, "0.00000")
	//
	// variable decimal places
	test("1234.00000", -1, "1,234")
	test("1234.50000", -1, "1,234.5")
	test("1234.56000", -1, "1,234.56")
	test("1234.56700", -1, "1,234.567")
	test("1234.56780", -1, "1,234.5678")
	test("1234.56789", -1, "1,234.5678")
	//
	// more decimals than Fixed64's precision
	test("1234.00", 2, "1,234.00")
	test("1234.0000", 2, "1,234.00")
	test("1234.56", 2, "1,234.56")
	test("1234.5678", 0, "1,234")
	test("1234.5678", 1, "1,234.5")
	test("1234.5678", 2, "1,234.56")
	test("1234.5678", 3, "1,234.567")
	test("1234.5678", 4, "1,234.5678")
	test("1234.56780", 4, "1,234.5678")
	test("1234.56789", 4, "1,234.5678")
	test("1234.56789", 5, "1,234.56780")
	//
	// 0 digits
	test("", 0, "")
	//
	// 1 digit
	test(-1, 0, "-1")
	test(-9, 0, "-9")
	//
	// 2 digits
	test(-10, 0, "-10")
	test(-12, 0, "-12")
	test(-99, 0, "-99")
	//
	// 3 digits
	test(-100, 0, "-100")
	test(-123, 0, "-123")
	test(-201, 0, "-201")
	test(-999, 0, "-999")
	//
	// 4 digits
	test(-1000, 0, "-1,000")
	test(-1001, 0, "-1,001")
	test(-1234, 0, "-1,234")
	test(-9999, 0, "-9,999")
	//
	// 1 digit
	test(1, 0, "1")
	test(9, 0, "9")
	//
	// 2 digits
	test(10, 0, "10")
	test(12, 0, "12")
	test(99, 0, "99")
	//
	// 3 digits
	test(100, 0, "100")
	test(123, 0, "123")
	test(201, 0, "201")
	test(999, 0, "999")
	//
	// 4 digits
	test(1000, 0, "1,000")
	test(1001, 0, "1,001")
	test(1234, 0, "1,234")
	test(9999, 0, "9,999")
	//
	// large numbers
	test("-922337203685477", 0, "-922,337,203,685,477")
	test("100000000000000", 0, "100,000,000,000,000")
	test("22337203685477", 0, "22,337,203,685,477")
	test("900000000000000", 0, "900,000,000,000,000")
	test("922337203685477", 0, "922,337,203,685,477")
	test("99999999999999", 0, "99,999,999,999,999")
}

// go test --run Test_Fixed64_String_
func Test_Fixed64_String_(t *testing.T) {
	//
	// (n Fixed64) String() string
	//
	test := func(n Fixed64, want string) {
		var (
			old   = n
			got   = n.String()
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.String() =", testLine(), n)
		)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
	}
	// null value
	test(Fixed64{NaN}, "")
	//
	// zero
	test(Fixed64{0}, "0")
	//
	// integers
	test(Fixed64{-1 * 1E4}, "-1")
	test(Fixed64{-12 * 1E4}, "-12")
	test(Fixed64{-123 * 1E4}, "-123")
	test(Fixed64{-1234 * 1E4}, "-1234")
	test(Fixed64{-12345 * 1E4}, "-12345")
	test(Fixed64{1 * 1E4}, "1")
	test(Fixed64{12 * 1E4}, "12")
	test(Fixed64{123 * 1E4}, "123")
	test(Fixed64{1234 * 1E4}, "1234")
	test(Fixed64{12345 * 1E4}, "12345")
	//
	// numbers with zero int part and some decimals
	test(Fixed64{-1000}, "-0.1")
	test(Fixed64{-100}, "-0.01")
	test(Fixed64{-10}, "-0.001")
	test(Fixed64{-1}, "-0.0001")
	test(Fixed64{1000}, "0.1")
	test(Fixed64{100}, "0.01")
	test(Fixed64{10}, "0.001")
	test(Fixed64{1}, "0.0001")
	//
	test(Fixed64{-1200}, "-0.12")
	test(Fixed64{-1230}, "-0.123")
	test(Fixed64{-12345}, "-1.2345")
	test(Fixed64{-1234}, "-0.1234")
	test(Fixed64{-123}, "-0.0123")
	test(Fixed64{-12}, "-0.0012")
	test(Fixed64{-9999}, "-0.9999")
	test(Fixed64{1200}, "0.12")
	test(Fixed64{1230}, "0.123")
	test(Fixed64{1234}, "0.1234")
	test(Fixed64{123}, "0.0123")
	test(Fixed64{12}, "0.0012")
	test(Fixed64{9999}, "0.9999")
	//
	// numbers with both int part and decimals
	test(Fixed64{-10100}, "-1.01")
	test(Fixed64{-1234*1E4 - 5678}, "-1234.5678")
	test(Fixed64{-123456789}, "-12345.6789")
	test(Fixed64{-2000000000000000001}, "-200000000000000.0001")
	test(Fixed64{10100}, "1.01")
	test(Fixed64{1234*1E4 + 5678}, "1234.5678")
	test(Fixed64{123456789}, "12345.6789")
	test(Fixed64{12345}, "1.2345")
	test(Fixed64{2000000000000000001}, "200000000000000.0001")
	//
	// smallest and largest fixed-point numbers (about 922 trillion):
	test(Fixed64{-9223372036854769999}, "-922337203685476.9999")
	test(Fixed64{9223372036854769999}, "922337203685476.9999")
	//
	// these are overflowing values, but String() must print them anyway
	test(Fixed64{math.MaxInt64}, "922337203685477.5807")
	test(Fixed64{math.MinInt64 + 1}, "-922337203685477.5807")
}

// -----------------------------------------------------------------------------
// # Division:

// go test --run Test_Fixed64_Div_
func Test_Fixed64_Div_(t *testing.T) {
	//
	// (n Fixed64) Div(nums ...Fixed64) Fixed64
	//
	test := func(n Fixed64, nums []Fixed64, want Fixed64) {
		var (
			old   = n
			got   = n.Div(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.Div(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	//           12345.6789   /              1.0    =       12345.6789
	test(Fixed64{123456789}, AN(Fixed64{10000}), Fixed64{123456789})
}

// go test --run Test_Fixed64_DivFloat_
func Test_Fixed64_DivFloat_(t *testing.T) {
	//
	// (n Fixed64) DivFloat(nums ...float64) Fixed64
	//
	test := func(n Fixed64, nums []float64, want Fixed64) {
		var (
			old   = n
			got   = n.DivFloat(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.DivFloat(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	//           12345.6789    / 1.0   =       12345.6789
	test(Fixed64{123456789}, AF(1.0), Fixed64{123456789})
	//
	//           12345.6789 /    2.0   =       6172.8394
	test(Fixed64{123456789}, AF(2.0), Fixed64{61728394})
}

// go test --run Test_Fixed64_DivInt_
func Test_Fixed64_DivInt_(t *testing.T) {
	//
	// (n Fixed64) DivInt(nums ...int) Fixed64
	//
	test := func(n Fixed64, nums []int, want Fixed64) {
		var (
			old   = n
			got   = n.DivInt(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.DivInt(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	//
	// TODO: declaration comment
	//
	//       12345.6789 / 1   =               12345.6789
	test(Fixed64{123456789}, AI(1), Fixed64{123456789})
	//
	//       12345.6789 / 2   =               6172.8394
	test(Fixed64{123456789}, AI(2), Fixed64{61728394})
}

// go test --run Test_Fixed64_DivInt64_
func Test_Fixed64_DivInt64_(t *testing.T) {
	//
	// (n Fixed64) DivInt64(nums ...int64) Fixed64
	//
	test := func(n Fixed64, nums []int64, want Fixed64) {
		var (
			old   = n
			got   = n.DivInt64(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.DivInt64(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	//           12345.6789   /       1   =       12345.6789
	test(Fixed64{123456789}, []int64{1}, Fixed64{123456789})
	//
	//           12345.6789   /       2   =       6172.8394
	test(Fixed64{123456789}, []int64{2}, Fixed64{61728394})
	// TODO: add more test cases
}

// -----------------------------------------------------------------------------
// # Multiplication:

// go test --run Test_Fixed64_Mul_
func Test_Fixed64_Mul_(t *testing.T) {
	//
	// (n Fixed64) Mul(nums ...Fixed64) Fixed64
	//
	test := func(n Fixed64, nums []Fixed64, want Fixed64) {
		var (
			old   = n
			got   = n.Mul(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.Mul(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(New(-2), AN(New(2)), New(-4))
	test(New(0), AN(New(0)), New(0))
	test(New(1), AN(New(1)), New(1))
	test(New(1000), AN(New(1000)), New(1000000))
	test(New(2), AN(New(-2)), New(-4))
	test(New(2), AN(New(2)), New(4))
	//
	// these would overflow if big.Int is not used
	test(Fixed64{4611686018427384999},
		AN(New(2)), Fixed64{9223372036854769998})
	test(Fixed64{922337203685476999},
		AN(New(10)), Fixed64{9223372036854769990})
	//
	// tests that cause overflow
	testOver := func(n Fixed64, nums []Fixed64, want Fixed64) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.Mul(%#v)",
			testLine(), n, nums)
		testErrors(label,
			func() error {
				var (
					old = n
					got = n.Mul(nums...)
				)
				if n != old {
					t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
				}
				testGot(label, got, want, func(erm string) { t.Error(erm) })
				return NoError
			},
			func(erm string) { t.Error(erm) },
		)
	}
	// overflow: the result fits in int64, but not in acceptable Fixed64 range
	testOver(
		Fixed64{maxInt64 + 1},
		AN(New(1)),
		Fixed64{math.MaxInt64},
	)
	testOver(
		New(int64(-627199999000000)),
		AN(Fixed64{627199999000000}),
		Fixed64{math.MinInt64 + 1},
	)
	testOver(
		New(int64(-922337203685476)),
		AN(Fixed64{922337203685476}),
		Fixed64{math.MinInt64 + 1},
	)
	testOver(
		New(int64(627199999000000)),
		AN(Fixed64{627199999000000}),
		Fixed64{math.MaxInt64},
	)
	testOver(
		New(int64(922337203685476)),
		AN(Fixed64{922337203685476}),
		Fixed64{math.MaxInt64},
	)
}

// go test --run Test_Fixed64_Mul_2_
func Test_Fixed64_Mul_2_(t *testing.T) {
	//
	// (n Fixed64) Mul(nums ...Fixed64) Fixed64
	//
	test := func(want, num int64, nums ...int64) {
		var ar []Fixed64
		for _, i64 := range nums {
			ar = append(ar, Fixed64{i64})
		}
		var (
			got = Fixed64{num}.Mul(ar...)
		)
		label := fmt.Sprintf("<-L%d: Fixed64_Mul(%#v, %#v) =",
			testLine(), num, nums)
		testGot(label, got.i64, want, func(erm string) { t.Error(erm) })
	}
	const (
		IntLimit = 922337203685476
	)
	test(0, 0, 0)             // 0 = 0 x 0
	test(4*1E4, 2*1E4, 2*1E4) // 4.0000 = 2.0000 x 2.0000
	//
	// 922,337,203,681,802.9970 = square 30,370,004.9997
	test(9223372036818029970, 303700049997, 303700049997)
	//
	// test overflow
	testOver := func(want, num int64, nums ...int64) {
		var ar []Fixed64
		for _, i64 := range nums {
			ar = append(ar, Fixed64{i64})
		}
		mockError(true)
		var (
			got = Fixed64{num}.Mul(ar...)
		)
		if !gotError {
			t.Error("Failed to set mod.Error")
		}
		mockError(false)
		label := fmt.Sprintf("<-L%d: Fixed64{%#v}.MulInt64(%#v) =",
			testLine(), num, nums)
		testGot(label, got.i64, want, func(erm string) { t.Error(erm) })
	}
	//
	// 30,370,004.9998
	testOver(math.MinInt64+1, -303700049998, 303700049998)
	testOver(math.MinInt64+1, 303700049998, -303700049998)
	testOver(math.MaxInt64, 303700049998, 303700049998)
	//
	testOver(math.MinInt64+1, -IntLimit*1E4, IntLimit*1E4)
	testOver(math.MinInt64+1, IntLimit*1E4, -IntLimit*1E4)
	testOver(math.MaxInt64, IntLimit*1E4, IntLimit*1E4)
	//
	testOver(math.MinInt64+1, minInt64-1, 1E4)
	testOver(math.MinInt64+1, 1E4, minInt64-1)
	testOver(math.MaxInt64, maxInt64+1, 1E4)
	testOver(math.MaxInt64, 1E4, maxInt64+1)
	//
	testOver(math.MinInt64+1, -9004443332221119999, 9004443332221119999)
	testOver(math.MinInt64+1, 9004443332221119999, -9004443332221119999)
	testOver(math.MaxInt64, 9004443332221119999, 9004443332221119999)
}

// go test --run Test_Fixed64_MulFloat_
func Test_Fixed64_MulFloat_(t *testing.T) {
	//
	// (n Fixed64) MulFloat(nums ...float64) Fixed64
	//
	test := func(n Fixed64, nums []float64, want Fixed64) {
		var (
			old   = n
			got   = n.MulFloat(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.MulFloat(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	//            0   *  1.0   =        0
	test(Fixed64{0}, AF(1.0), Fixed64{0})
	//
	//            2   *        2.0   =        4
	test(Fixed64{2 * 1E4}, AF(2.0), Fixed64{4 * 1E4})
	//
	//           12345.6789   *  1   =       12345.6789
	test(Fixed64{123456789}, AF(1), Fixed64{123456789})
	//
	//           12345.6789   /  2   =       24691.3578
	test(Fixed64{123456789}, AF(2), Fixed64{246913578})
	//
	test(New(-1), AF(1e10), New(int64(-10000000000))) // -10 billion
	test(New(-1), AF(1e6), New(-1000000))             // -1 million
	test(New(-1), AF(1e7), New(-10000000))            // -10 million
	test(New(-1), AF(1e8), New(-100000000))           // -100 million
	test(New(-1), AF(1e9), New(-1000000000))          // -1 billion
	//
	test(New(1), AF(1e10), New(int64(10000000000))) // 10 billion
	test(New(1), AF(1e6), New(1000000))             // 1 million
	test(New(1), AF(1e7), New(10000000))            // 10 million
	test(New(1), AF(1e8), New(100000000))           // 100 million
	test(New(1), AF(1e9), New(1000000000))          // 1 billion
	//
	// Here, if Mul() didn't use big.Int, multiplication would overflow
	test(New(-1), AF(1e14), Fixed64{int64(-1e18)})
	test(New(1), AF(1e14), Fixed64{int64(1e18)})
	test(New(123), AF(1e9), Fixed64{int64(123 * 1e13)}) // 123 billion
	//
	// overflow
	testOver := func(n Fixed64, nums []float64, want Fixed64) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.MulFloat(%#v)",
			testLine(), n, nums)
		testErrors(label,
			func() error {
				var (
					old = n
					got = n.MulFloat(nums...)
				)
				if n != old {
					t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
				}
				testGot(label, got, want, func(erm string) { t.Error(erm) })
				return NoError
			},
			func(erm string) { t.Error(erm) },
		)
	}
	testOver(New(1), AF(1e20),
		Fixed64{math.MaxInt64})
	testOver(New(123), AF(float64(maxInt64)),
		Fixed64{math.MaxInt64})
	testOver(New(123), AF(float64(minInt64)-0.001),
		Fixed64{math.MinInt64 + 1})
}

// go test --run Test_Fixed64_MulInt_
func Test_Fixed64_MulInt_(t *testing.T) {
	//
	// (n Fixed64) MulInt(nums ...int) Fixed64
	//
	test := func(n Fixed64, nums []int, want Fixed64) {
		var (
			old   = n
			got   = n.MulInt(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.MulInt(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	//            0   *   1   =        0
	test(Fixed64{0}, AI(1), Fixed64{0})
	//
	//            2         *   2   =        4
	test(Fixed64{2 * 1E4}, AI(2), Fixed64{4 * 1E4})
	//
	//           12345.6789   *   1   =       12345.6789
	test(Fixed64{123456789}, AI(1), Fixed64{123456789})
	//
	//           12345.6789   /   2   =       24691.3578
	test(Fixed64{123456789}, AI(2), Fixed64{246913578})
	//
	// TODO: Try to cause failure in unit test:
	//       pass MaxInt64 * just above fixed-number limit.
}

// go test --run Test_Fixed64_MulInt64_
func Test_Fixed64_MulInt64_(t *testing.T) {
	//
	// (n Fixed64) MulInt64(nums ...int64) Fixed64
	//
	test := func(n Fixed64, nums []int64, want Fixed64) {
		var (
			old   = n
			got   = n.MulInt64(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.MulInt64(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(Fixed64{-123 * 1E4}, []int64{2}, Fixed64{-246 * 1E4})
	test(Fixed64{2 * 1E4}, []int64{2, 2, 2, 2, 2}, Fixed64{64 * 1E4})
	test(Fixed64{2 * 1E4}, []int64{2}, Fixed64{4 * 1E4})
	test(New(0), []int64{0}, Fixed64{0})
	test(New(0), []int64{1, 3, 5, 7, 11}, Fixed64{0})
	test(New(1), []int64{3, 5, 7}, Fixed64{105 * 1E4})
	// TODO: add more test cases
}

// -----------------------------------------------------------------------------
// # Addition:

// go test --run Test_Fixed64_Add_
func Test_Fixed64_Add_(t *testing.T) {
	//
	// (n Fixed64) Add(nums ...Fixed64) Fixed64
	//
	test := func(n Fixed64, nums []Fixed64, want Fixed64) {
		var (
			old   = n
			got   = n.Add(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.Add(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	// test additions within range of Fixed64:
	test(New(-1), AN(New(-1)), New(-2))
	test(New(0), AN(New(0)), New(0))
	test(New(1), AN(New(1)), New(2))
	test(New(1000), AN(New(1000)), New(2000))
	//
	// multiple additions:
	test(New(1),
		AN(New(1.1), New(2.02), New(3.003), New(4.0004)),
		New(11.1234))
	//
	// test addition close to the limits of Fixed64:
	//
	// largest + smallest number: must 'annihilate' to zero
	test(
		Fixed64{minInt64},     // -9223372036854769999
		AN(Fixed64{maxInt64}), //  9223372036854769999
		Fixed64{0},
	)
	test(
		Fixed64{maxInt64},     //  9223372036854769999
		AN(Fixed64{minInt64}), // -9223372036854769999
		Fixed64{0},
	)
	// -922337203685476 + -0.9999 = -922337203685476.9999
	test(
		Fixed64{-9223372036854760000},
		AN(Fixed64{-9999}),
		Fixed64{minInt64},
	)
	test(
		Fixed64{-9999},
		AN(Fixed64{-9223372036854760000}),
		Fixed64{minInt64},
	)
	// -922337203685476.9998 + -0.0001 = -922337203685476.9999
	test(
		Fixed64{-9223372036854769998},
		AN(Fixed64{-1}),
		Fixed64{minInt64},
	)
	test(
		Fixed64{-1},
		AN(Fixed64{-9223372036854769998}),
		Fixed64{minInt64},
	)
	// 922337203685476 + 0.9999 = 922337203685476.9999
	test(
		Fixed64{9223372036854760000},
		AN(Fixed64{9999}),
		Fixed64{maxInt64},
	)
	test(
		Fixed64{9999},
		AN(Fixed64{9223372036854760000}),
		Fixed64{maxInt64},
	)
	// 922337203685476.9998 + 0.0001 = 922337203685476.9999
	test(
		Fixed64{9223372036854769998},
		AN(Fixed64{1}),
		Fixed64{maxInt64},
	)
	test(
		Fixed64{1},
		AN(Fixed64{9223372036854769998}),
		Fixed64{maxInt64},
	)
	// test additions that will overflow Fixed64:
	testOver := func(n Fixed64, nums []Fixed64, want Fixed64) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.Add(%#v)",
			testLine(), n, nums)
		testErrors(label,
			func() error {
				var (
					old = n
					got = n.Add(nums...)
				)
				if n != old {
					t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
				}
				testGot(label, got, want, func(erm string) { t.Error(erm) })
				return NoError
			},
			func(erm string) { t.Error(erm) },
		)
	}
	// adding -0.0001 to the smallest number must overflow
	testOver(
		Fixed64{minInt64}, // -922337203685476.9999
		AN(Fixed64{-1}),   //               -0.0001
		Fixed64{math.MinInt64 + 1},
	)
	// adding 0.0001 to the largest number must overflow
	testOver(
		Fixed64{maxInt64}, //  922337203685476.9999
		AN(Fixed64{1}),    //                0.0001
		Fixed64{math.MaxInt64},
	)
	// adding two smallest numbers must overflow
	testOver(
		Fixed64{minInt64},     // -922337203685476.9999
		AN(Fixed64{minInt64}), // -922337203685476.9999
		Fixed64{math.MinInt64 + 1},
	)
	// adding two largest numbers must overflow
	testOver(
		Fixed64{maxInt64},     //  922337203685476.9999
		AN(Fixed64{maxInt64}), //  922337203685476.9999
		Fixed64{math.MaxInt64},
	)
}

// go test --run Test_Fixed64_AddFloat_
func Test_Fixed64_AddFloat_(t *testing.T) {
	//
	// (n Fixed64) AddFloat(nums ...float64) Fixed64
	//
	test := func(n Fixed64, nums []float64, want Fixed64) {
		var (
			old   = n
			got   = n.AddFloat(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.AddFloat(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(New(0), AF(0), New(0))
	test(New(1), AF(1), New(2))
	//
	// test multiple arguments
	test(New(1), AF(2, 3, 4), New(10))
	//
	// must overflow
	testOver := func(n Fixed64, nums []float64, want Fixed64) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.AddFloat(%#v)",
			testLine(), n, nums)
		testErrors(label,
			func() error {
				var (
					old = n
					got = n.AddFloat(nums...)
				)
				if n != old {
					t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
				}
				testGot(label, got, want, func(erm string) { t.Error(erm) })
				return NoError
			},
			func(erm string) { t.Error(erm) },
		)
	}
	testOver(New(1), AF(1e30), Fixed64{math.MaxInt64})
}

// go test --run Test_Fixed64_AddInt_
func Test_Fixed64_AddInt_(t *testing.T) {
	//
	// (n Fixed64) AddInt(nums ...int) Fixed64
	//
	test := func(n Fixed64, nums []int, want Fixed64) {
		var (
			old   = n
			got   = n.AddInt(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.AddInt(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(New(0), AI(0), New(0))
}

// go test --run Test_Fixed64_AddInt64_
func Test_Fixed64_AddInt64_(t *testing.T) {
	//
	// (n Fixed64) AddInt64(nums ...int64) Fixed64
	//
	test := func(n Fixed64, nums []int64, want Fixed64) {
		var (
			old   = n
			got   = n.AddInt64(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.AddInt64(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(Fixed64{-12345 * 1E4}, []int64{12345}, Fixed64{0})
	test(Fixed64{2 * 1E4}, []int64{2}, Fixed64{4 * 1E4})
	test(New(0), []int64{0}, Fixed64{0})
	test(New(0), []int64{1, 3, 5, 7, 11}, Fixed64{27 * 1E4})
	// TODO: add more test cases
}

// -----------------------------------------------------------------------------
// # Subtraction:

// go test --run Test_Fixed64_Sub_
func Test_Fixed64_Sub_(t *testing.T) {
	//
	// (n Fixed64) Sub(nums ...Fixed64) Fixed64
	//
	test := func(n Fixed64, nums []Fixed64, want Fixed64) {
		var (
			old   = n
			got   = n.Sub(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.Sub(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	// test subtractions within range of Fixed64:
	test(New(-2), AN(New(-1)), New(-1))
	test(New(0), AN(New(0)), New(0))
	test(New(0), AN(New(0.0001)), New(-0.0001))
	test(New(0.0001), AN(New(0.0001)), New(0))
	test(New(1), AN(New(1)), New(0))
	test(New(10), AN(New(1), New(2), New(3)), New(4))
	test(New(2), AN(New(1)), New(1))
	test(New(2000), AN(New(1000)), New(1000))
	test(New(int64(1234567891234)),
		AN(New(67891234)),
		New(int64(1234500000000)))
	//
	// multiple subtractions:
	test(New(11.1234),
		AN(New(4.0004), New(3.003), New(2.02), New(1.1)),
		New(1))
	//
	// test subtraction close to the limits of Fixed64:
	//
	// largest - largest number: must be to zero
	test(
		Fixed64{maxInt64},     // 9223372036854769999
		AN(Fixed64{maxInt64}), // 9223372036854769999
		Fixed64{0},
	)
	// -922337203685476.9999 - -0.9999 = -922337203685476
	test(
		Fixed64{-9223372036854769999},
		AN(Fixed64{-9999}),
		Fixed64{-9223372036854760000},
	)
	test(
		Fixed64{minInt64},
		AN(Fixed64{-9223372036854760000}),
		Fixed64{-9999},
	)
	// -922337203685476.9999 - -922337203685476.9998 = -0.0001
	test(
		Fixed64{minInt64},
		AN(Fixed64{-9223372036854769998}),
		Fixed64{-1},
	)
	// -922337203685476.9999 - 0.0001 = -922337203685476.9998
	test(
		Fixed64{minInt64},
		AN(Fixed64{-1}),
		Fixed64{-9223372036854769998},
	)
	// 922337203685476.9999 - 922337203685476 = 0.9999
	test(
		Fixed64{maxInt64},
		AN(Fixed64{9223372036854760000}),
		Fixed64{9999},
	)
	// 922337203685476.9999 - -0.9999 = 922337203685476
	test(
		Fixed64{9223372036854760000},
		AN(Fixed64{-9999}),
		Fixed64{maxInt64},
	)
	// 922337203685476.9999 - 922337203685476.9998 = 0.0001
	test(
		Fixed64{maxInt64},
		AN(Fixed64{9223372036854769998}),
		Fixed64{1},
	)
	// 922337203685476.9999 - 0.0001 = 922337203685476.9998
	test(
		Fixed64{maxInt64},
		AN(Fixed64{1}),
		Fixed64{9223372036854769998},
	)
	// test subtraction that will overflow Fixed64:
	testOver := func(n Fixed64, nums []Fixed64, want Fixed64) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.Sub(%#v)",
			testLine(), n, nums)
		testErrors(label,
			func() error {
				var (
					old = n
					got = n.Sub(nums...)
				)
				if n != old {
					t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
				}
				testGot(label, got, want, func(erm string) { t.Error(erm) })
				return NoError
			},
			func(erm string) { t.Error(erm) },
		)
	}
	// subtracting 0.0001 from the smallest number must overflow
	testOver(
		Fixed64{minInt64}, // -922337203685476.9999
		AN(Fixed64{1}),    //               -0.0001
		Fixed64{math.MinInt64 + 1},
	)
	// subtracting -0.0001 from the largest number must overflow
	testOver(
		Fixed64{maxInt64}, // 922337203685476.9999
		AN(Fixed64{-1}),   //               0.0001
		Fixed64{math.MaxInt64},
	)
	// subtracting largest from smallest number must overflow
	testOver(
		Fixed64{minInt64},     // -922337203685476.9999
		AN(Fixed64{maxInt64}), //  922337203685476.9999
		Fixed64{math.MinInt64 + 1},
	)
	// subtracting smallest from largest number must overflow
	testOver(
		Fixed64{maxInt64},     //  922337203685476.9999
		AN(Fixed64{minInt64}), // -922337203685476.9999
		Fixed64{math.MaxInt64},
	)
}

// go test --run Test_Fixed64_SubFloat_
func Test_Fixed64_SubFloat_(t *testing.T) {
	//
	// (n Fixed64) SubFloat(nums ...float64) Fixed64
	//
	test := func(n Fixed64, nums []float64, want Fixed64) {
		var (
			old   = n
			got   = n.SubFloat(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.SubFloat(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(New(0), AF(0), New(0))
	test(New(0), AF(123456789), New(-123456789))
	test(New(0), AF(2, 3, 5, 7, 11), New(-28))
	test(New(123456789), AF(123456789), New(0))
	test(New(28), AF(2, 3, 5, 7, 11), New(0))
	// TODO: add more test cases
}

// go test --run Test_Fixed64_SubInt_
func Test_Fixed64_SubInt_(t *testing.T) {
	//
	// (n Fixed64) SubInt(nums ...int) Fixed64
	//
	test := func(n Fixed64, nums []int, want Fixed64) {
		var (
			old   = n
			got   = n.SubInt(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.SubInt(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(New(0), AI(0), New(0))
	test(New(0), AI(123456789), New(-123456789))
	test(New(0), AI(2, 3, 5, 7, 11), New(-28))
	test(New(123456789), AI(123456789), New(0))
	test(New(28), AI(2, 3, 5, 7, 11), New(0))
	// TODO: add more test cases
}

// go test --run Test_Fixed64_SubInt64_
func Test_Fixed64_SubInt64_(t *testing.T) {
	//
	// (n Fixed64) SubInt64(nums ...int64) Fixed64
	//
	test := func(n Fixed64, nums []int64, want Fixed64) {
		var (
			old   = n
			got   = n.SubInt64(nums...)
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.SubInt64(%#v)",
				testLine(), n, nums)
		)
		if n != old {
			t.Errorf("%s mutated from <%s> to <%s>", label, old, n)
		}
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(Fixed64{-12345 * 1E4}, []int64{-12345}, Fixed64{0})
	test(Fixed64{12345 * 1E4}, []int64{12345}, Fixed64{0})
	test(Fixed64{4 * 1E4}, []int64{2}, Fixed64{2 * 1E4})
	test(New(0), []int64{0}, Fixed64{0})
	test(New(0), []int64{1, 3, 5, 7, 11}, Fixed64{-27 * 1E4})
	// TODO: add more test cases
}

// -----------------------------------------------------------------------------
// # Information:

// go test --run Test_Fixed64_Float64_
func Test_Fixed64_Float64_(t *testing.T) {
	//
	// (n Fixed64) Float64() float64
	//
	test := func(n Fixed64, want float64) {
		var (
			got   = n.Float64()
			label = fmt.Sprintf("<-L%d: Fixed64<%s>.Float64()=", testLine(), n)
		)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(Fixed64{-1}, -0.0001)
	test(Fixed64{1}, 0.0001)
	test(Fixed64{maxInt64}, 9.22337203685477e+14)
	test(Fixed64{minInt64}, -9.22337203685477e+14)
	test(New(0), 0.0)
	test(New(1234567890), 1234567890)
	test(New(987654321), 987654321)
}

// go test --run Test_Fixed64_Int_
func Test_Fixed64_Int_(t *testing.T) {
	//
	// (n Fixed64) Int() int64
	//
	test := func(n Fixed64, want int) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.Int() =", testLine(), n)
		got := n.Int()
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(New(-1), -1)
	test(New(0), 0)
	test(New(1), 1)
	// TODO: add more test cases
}

// go test --run Test_Fixed64_Int64_
func Test_Fixed64_Int64_(t *testing.T) {
	//
	// (n Fixed64) Int64() int64
	//
	test := func(n Fixed64, want int64) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.Int64() =", testLine(), n)
		got := n.Int64()
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(New(-1), -1)
	test(New(0), 0)
	test(New(1), 1)
	// TODO: add more test cases
}

// go test --run Test_Fixed64_IsEqual_
func Test_Fixed64_IsEqual_(t *testing.T) {
	//
	// (n Fixed64) IsEqual() bool
	//
	test := func(n, num Fixed64, want bool) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.IsEqual(%#v) =",
			testLine(), n, num)
		got := n.IsEqual(num)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(New(0), New(0), true)
	test(New(0), New(1), false)
	test(New(1), New(0), false)
	test(New(1), New(1), true)
	// TODO: add more test cases
}

// go test --run Test_Fixed64_IsGreater_
func Test_Fixed64_IsGreater_(t *testing.T) {
	//
	// (n Fixed64) IsGreater(value interface{}) bool
	//
	test := func(n, num Fixed64, want bool) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.IsGreater(%#v) =",
			testLine(), n, num)
		got := n.IsGreater(num)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	{
		test(New(0), New(-1), true)
		test(New(1), New(0), true)
		test(New(100), New(-200), true)
	}
	{
		test(New(-1), New(-1), false)
		test(New(0), New(1), false)
		test(New(1), New(1), false)
	}
	// TODO: add more test cases
}

// go test --run Test_Fixed64_IsGreaterOrEqual_
func Test_Fixed64_IsGreaterOrEqual_(t *testing.T) {
	//
	// (n Fixed64) IsGreaterOrEqual(value interface{}) bool
	//
	test := func(n, num Fixed64, want bool) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.IsGreaterOrEqual(%#v) =",
			testLine(), n, num)
		got := n.IsGreaterOrEqual(num)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	{
		test(New(-200), New(-200), true)
		test(New(0), New(-1), true)
		test(New(0), New(0), true)
		test(New(1), New(0), true)
		test(New(1), New(1), true)
		test(New(100), New(-200), true)
	}
	{
		test(New(-2), New(-1), false)
		test(New(0), New(1), false)
		test(New(1), New(2), false)
	}
	// TODO: add more test cases
}

// go test --run Test_Fixed64_IsLesser_
func Test_Fixed64_IsLesser_(t *testing.T) {
	//
	// (n Fixed64) IsLesser(value interface{}) bool
	//
	test := func(n, num Fixed64, want bool) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.IsLesser(%#v) =",
			testLine(), n, num)
		got := n.IsLesser(num)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	{
		test(New(-1), New(0), true)
		test(New(0), New(1), true)
		test(New(-200), New(100), true)
	}
	{
		test(New(-1), New(-1), false)
		test(New(1), New(0), false)
		test(New(1), New(1), false)
	}
	// TODO: add more test cases
}

// go test --run Test_Fixed64_IsLesserOrEqual_
func Test_Fixed64_IsLesserOrEqual_(t *testing.T) {
	//
	// (n Fixed64) IsLesserOrEqual(value interface{}) bool
	//
	test := func(n, num Fixed64, want bool) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.IsLesserOrEqual(%#v) =",
			testLine(), n, num)
		got := n.IsLesserOrEqual(num)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	{
		test(New(-200), New(-200), true)
		test(New(-1), New(0), true)
		test(New(0), New(0), true)
		test(New(0), New(1), true)
		test(New(1), New(1), true)
		test(New(-200), New(100), true)
	}
	{
		test(New(-1), New(-2), false)
		test(New(1), New(0), false)
		test(New(2), New(1), false)
	}
	// TODO: add more test cases
}

// go test --run Test_Fixed64_IsNegative_
func Test_Fixed64_IsNegative_(t *testing.T) {
	//
	// (n Fixed64) IsNegative() bool
	//
	test := func(n Fixed64, want bool) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.IsNegative() =",
			testLine(), n)
		got := n.IsNegative()
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	{
		test(Fixed64{minInt64}, true)
		test(New(-1), true)
		test(New(-1000), true)
	}
	{
		test(Fixed64{NaN}, false)
		test(Fixed64{maxInt64}, false)
		test(New(0), false)
		test(New(1), false)
		test(New(1000), false)
	}
	// TODO: add more test cases
}

// go test --run Test_Fixed64_IsZero_
func Test_Fixed64_IsZero_(t *testing.T) {
	//
	// (n Fixed64) IsZero() bool
	//
	test := func(n Fixed64, want bool) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.IsZero() =",
			testLine(), n)
		got := n.IsZero()
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	{
		test(Fixed64{0}, true)
		test(New("-0.00001"), true) // because this is > 4 decimal places
		test(New("0"), true)
		test(New("00"), true)
		test(New(0), true)
		test(New(float32(0)), true)
		test(New(float64(0)), true)
	}
	{
		test(Fixed64{-1}, false)
		test(Fixed64{1}, false)
		test(Fixed64{maxInt64}, false)
		test(Fixed64{minInt64}, false)
		test(New(""), false) // this is null, not zero
		test(New("-0.0001"), false)
		test(New("-922337203685476.9999"), false) // lowest fixed-point number
		test(New("0.0001"), false)
		test(New("1"), false)
		test(New("922337203685476.9999"), false) // highest fixed-point number
		test(New(-1), false)
		test(New(-1000), false)
		test(New(1), false)
		test(New(1000), false)
	}
	// TODO: add more test cases
	//
	// logged with errors
	// TODO: catch error here
	//       {"XYZ", true}, // non-numeric strings are initialized to zero
}

// go test --run Test_Fixed64_Overflow_
func Test_Fixed64_Overflow_(t *testing.T) {
	//
	// (n Fixed64) Overflow() int
	//
	test := func(n Fixed64, want int) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.Overflow() =", testLine(), n)
		got := n.Overflow()
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	{
		// no overflow
		test(Fixed64{0}, 0)
		test(Fixed64{maxInt64}, 0)
		test(Fixed64{minInt64}, 0)
	}
	{
		// negative overflow
		test(Fixed64{math.MinInt64}, -1)
		test(Fixed64{minInt64 - 1}, -1)
	}
	{
		// positive overflow
		test(Fixed64{math.MaxInt64}, 1)
		test(Fixed64{maxInt64 + 1}, 1)
	}
}

// go test --run Test_Fixed64_Unwrap_
func Test_Fixed64_Unwrap_(t *testing.T) {
	//
	// (n Fixed64) Unwrap() int64
	//
	test := func(n Fixed64, want int64) {
		label := fmt.Sprintf("<-L%d: Fixed64<%s>.Unwrap() =", testLine(), n)
		got := n.Unwrap()
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	test(Fixed64{-123456789 * 1E4}, -1234567890000)
	test(Fixed64{0}, 0)
	test(Fixed64{1 * 1E4}, 10000)
	test(New(-123456789), -1234567890000)
	test(New(0), 0)
	test(New(1), 10000)
}

// -----------------------------------------------------------------------------
// # JSON:

// go test --run Test_Fixed64_MarshalJSON_
func Test_Fixed64_MarshalJSON_(t *testing.T) {
	//
	// (n Fixed64) MarshalJSON() ([]byte, error)
	//
	test := func(input interface{}, want string) {
		type Struct struct{ Num Fixed64 }
		var box Struct
		box.Num = New(input)
		script, _ := json.MarshalIndent(box, "", " ")
		//                  ^ calls the number's MarshalIndent() method
		got := string(script)
		got = strings.Replace(got, "\n", "", -1)
		got = strings.Replace(got, "{ ", "{", -1)
		{
			label := fmt.Sprintf("<-L%d: Fixed64.MarshalJSON()..", testLine())
			testGot(label, got, want, func(erm string) { t.Error(erm) })
		}
	}
	test(-1, `{"Num": -1}`)
	test(0, `{"Num": 0}`)
	test(0.0001, `{"Num": 0.0001}`)
	test(0.001, `{"Num": 0.001}`)
	test(0.01, `{"Num": 0.01}`)
	test(0.1, `{"Num": 0.1}`)
	test(1, `{"Num": 1}`)
	//
	test(-100000, `{"Num": -100000}`)   // -100 thousand
	test(-1000000, `{"Num": -1000000}`) // -1 million
	test(100000, `{"Num": 100000}`)     // 100 thousand
	test(1000000, `{"Num": 1000000}`)   // 1 million
	//
	test(math.MaxInt32, `{"Num": 2147483647}`)
	test(math.MinInt32, `{"Num": -2147483648}`)
	//
	test("1,000,000,000,000.0001", `{"Num": 1000000000000.0001}`)
	test("2.0001", `{"Num": 2.0001}`)
	test(123.456, `{"Num": 123.456}`)
	//
	//          Tn  Bn  Millions
	test("1,234,567,890,123.4321", `{"Num": 1234567890123.4321}`)
	test("9,999,999,999,999.9999", `{"Num": 9999999999999.9999}`)
	//
	// overflow
	testNoErrors("no overflow",
		func() error {
			test(int64(-922337203685476), `{"Num": -922337203685476}`)
			return nil
		},
		func(erm string) { t.Error(erm) },
	)
	testErrors("overflow",
		func() error {
			test(int64(922337203685477), `{"Num": 922337203685477.5807}`)
			return NoError
		},
		func(erm string) { t.Error(erm) },
	)
}

// go test --run Test_Fixed64_UnmarshalJSON_
func Test_Fixed64_UnmarshalJSON_(t *testing.T) {
	//
	// (n *Fixed64) UnmarshalJSON(data []byte) error
	//
	{
		label := `Fixed64.UnmarshalJSON("1234567")`
		var n Fixed64
		err := n.UnmarshalJSON([]byte("1234567"))
		if err != nil {
			t.Errorf("%s error: %s", label, err)
		}
		got := n
		want := New(1234567)
		testGot(label, got, want, func(erm string) { t.Error(erm) })
	}
	{
		testErrors(`<nil>.UnmarshalJSON("1234567")`,
			func() error {
				var n *Fixed64
				err := n.UnmarshalJSON([]byte("1234567"))
				return err
			},
			func(erm string) { t.Error(erm) },
		)
	}
	{
		testErrors("`Fixed64{}.UnmarshalJSON(<nil>)",
			func() error {
				var n Fixed64
				err := n.UnmarshalJSON(nil)
				return err
			},
			func(erm string) { t.Error(erm) },
		)
	}
}

// -----------------------------------------------------------------------------
// # Test Helper Functions

// AF is a convenience function to create an array of float64 numbers.
// That is, instead of having to specify '[]float64{1.0, 2.0, 3.0}'
// you can just use 'AF(1.0, 2.0, 3.0)'.
func AF(nums ...float64) (ret []float64) {
	for _, num := range nums {
		ret = append(ret, num)
	}
	return ret
}

// AI is a convenience function to create an array of int numbers.
// That is, instead of having to specify '[]int{1, 2, 3}'
// you can just use 'AI(1, 2, 3)'.
func AI(nums ...int) (ret []int) {
	for _, num := range nums {
		ret = append(ret, num)
	}
	return ret
}

// AN is a convenience function to create an array of Fixed64 numbers.
// That is, instead of having to specify
// '[]Fixed64{Fixed64{10000}, Fixed64{20000}, Fixed64{30000}}'
// you can just use 'AN(1, 2, 3)'.
func AN(nums ...Fixed64) (ret []Fixed64) {
	for _, num := range nums {
		ret = append(ret, New(num))
	}
	return ret
}

//end
