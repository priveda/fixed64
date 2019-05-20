// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-20 01:43:40 08C97A                       priveda/fixed64/[new.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"fmt"
	"reflect"
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
	switch v := value.(type) {
	case Fixed64:
		{
			return v
		}
	case int:
		{
			return New(int64(v))
		}
	case int64:
		{
			if v == NaN {
				return Fixed64{NaN}
			}
			if v < -IntLimit || v > IntLimit {
				return fixed64Overflow(v < 0, v)
			}
			return Fixed64{int64(v) * 1E4}
		}
	case uint:
		{
			return New(uint64(v))
		}
	case uint64:
		{
			if v > IntLimit {
				return fixed64Overflow(false, v)
			}
			return Fixed64{int64(v) * 1E4}
		}
	case float64:
		{
			if v < -float64(IntLimit)-0.9999 ||
				v > float64(IntLimit)+0.9999 {
				return fixed64Overflow(v < 0, v)
			}
			return Fixed64{int64(v * 1E4)}
		}
	case string:
		ret, _ := Parse(v)
		return ret
	}
	ret, done := new2(value) // try to convert other, less-common types
	if done {
		return ret
	}
	erm := fmt.Sprint("Type", reflect.TypeOf(value), "not handled; =", value)
	mod.Error(erm)
	return Fixed64{NaN}
}

// new2 is a helper function for New().
// It converts 8, 16 and 32-bit numeric types.
func new2(value interface{}) (Fixed64, bool) {
	switch v := value.(type) {
	case int32:
		{
			return New(int64(v)), true // call fn. to check range
		}
	case int16:
		{
			return Fixed64{int64(v) * 1E4}, true
		}
	case int8:
		{
			return Fixed64{int64(v) * 1E4}, true
		}
	case uint32:
		{
			return New(uint64(v)), true // call fn. to check range
		}
	case uint16:
		{
			return Fixed64{int64(v) * 1E4}, true
		}
	case uint8:
		{
			return Fixed64{int64(v) * 1E4}, true
		}
	case float32:
		if v < -float32(IntLimit)-0.9999 ||
			v > float32(IntLimit)+0.9999 {
			return fixed64Overflow(v < 0, v), true
		}
		return Fixed64{int64(float64(v) * 1E4)}, true
	}
	return new3(value) // try to convert pointer types
}

// new3 is an indirect helper function for New().
// It handles conversion from common pointer types and returns
// Fixed64{NaN} when given a nil pointer.
func new3(value interface{}) (Fixed64, bool) {
	if value == nil {
		return Fixed64{NaN}, true
	}
	rval := reflect.ValueOf(value)
	if rval.Kind() == reflect.Ptr && rval.IsNil() {
		return Fixed64{NaN}, true
	}
	switch v := value.(type) {
	case *int:
		{
			return New(int64(*v)), true
		}
	case *int64:
		{
			return New(*v), true
		}
	case *uint:
		{
			return New(uint64(*v)), true
		}
	case *uint64:
		{
			return New(*v), true
		}
	case *float64:
		{
			return New(*v), true
		}
	case *string:
		return New(*v), true
	}
	return new4(value)
}

// new4 is an indirect helper function for New().
// It handles conversion from pointers to smaller numeric types.
func new4(value interface{}) (Fixed64, bool) {
	switch v := value.(type) {
	case *int32:
		{
			return New(*v), true
		}
	case *int16:
		{
			return New(*v), true
		}
	case *int8:
		{
			return New(*v), true
		}
	case *uint32:
		{
			return New(*v), true
		}
	case *uint16:
		{
			return New(*v), true
		}
	case *uint8:
		{
			return New(*v), true
		}
	case *float32:
		return New(*v), true
	}
	// value could not be converted by these conversion
	// functions, pass false down the chain to New()
	return Fixed64{NaN}, false
}

//end
