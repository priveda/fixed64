// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:53:42 EC3A79                priveda/fixed64/[fixed64_of.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"fmt"
	"reflect"
)

// Fixed64Of converts any compatible type to a Fixed64. This includes all
// simple numeric types and strings, as well as pointers to these types.
//
// When given a non-numeric string, logs an error and returns Fixed64{NaN}
// which is a fixed-number type with a null value.
//
// When given a nil pointer, also returns Fixed64{NaN}.
//
// Note that fmt.Stringer() is not automatically handled as
// a string, to avoid hidden conversions and possible bugs.
//
func Fixed64Of(value interface{}) Fixed64 {
	switch v := value.(type) {
	case Fixed64:
		{
			return v
		}
	case int:
		{
			return Fixed64Of(int64(v))
		}
	case int64:
		{
			if v < -IntLimit || v > IntLimit {
				return fixed64Overflow(v < 0, EOverflow, ": ", v)
			}
			return Fixed64{int64(v) * 1E4}
		}
	case uint:
		{
			return Fixed64Of(uint64(v))
		}
	case uint64:
		{
			if v > IntLimit {
				return fixed64Overflow(false, EOverflow, "uint64: ", v)
			}
			return Fixed64{int64(v) * 1E4}
		}
	case float64:
		{
			if v < -float64(IntLimit)-0.9999 ||
				v > float64(IntLimit)+0.9999 {
				return fixed64Overflow(v < 0, EOverflow, "float64: ", v)
			}
			return Fixed64{int64(v * 1E4)}
		}
	case string:
		ret, _ := ParseFixed64(v)
		return ret
	}
	ret, done := fixed64Of2(value) // try to convert other, less-common types
	if done {
		return ret
	}
	erm := fmt.Sprint("Type", reflect.TypeOf(value), "not handled; =", value)
	mod.Error(erm)
	return Fixed64{NaN}
}

// fixed64Of2 is a helper function for Fixed64Of().
// It converts 8, 16 and 32-bit numeric types.
func fixed64Of2(value interface{}) (Fixed64, bool) {
	switch v := value.(type) {
	case int32:
		{
			return Fixed64Of(int64(v)), true // call fn. to check range
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
			return Fixed64Of(uint64(v)), true // call fn. to check range
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
			return fixed64Overflow(v < 0, EOverflow, "float32: ", v), true
		}
		return Fixed64{int64(float64(v) * 1E4)}, true
	}
	return fixed64Of3(value) // try to convert pointer types
}

// fixed64Of3 is an indirect helper function for Fixed64Of().
// It handles conversion from common pointer types and returns
// Fixed64{NaN} when given a nil pointer.
func fixed64Of3(value interface{}) (Fixed64, bool) {
	if value == nil {
		return Fixed64{NaN}, true
	}
	rval := reflect.ValueOf(value)
	if rval.Kind() == reflect.Ptr && rval.IsNil() {
		return Fixed64{NaN}, true
	}
	switch v := value.(type) {
	case *int64:
		{
			return Fixed64Of(*v), true
		}
	case *int:
		{
			return Fixed64Of(int64(*v)), true
		}
	case *uint64:
		{
			return Fixed64Of(*v), true
		}
	case *uint:
		{
			return Fixed64Of(uint64(*v)), true
		}
	case *float64:
		{
			return Fixed64Of(*v), true
		}
	case *string:
		{
			return Fixed64Of(*v), true
		}
	}
	return fixed64Of4(value)
}

// fixed64Of4 is an indirect helper function for Fixed64Of().
// It handles conversion from pointers to smaller numeric types.
func fixed64Of4(value interface{}) (Fixed64, bool) {
	switch v := value.(type) {
	case *int32:
		{
			return Fixed64Of(*v), true
		}
	case *int16:
		{
			return Fixed64Of(*v), true
		}
	case *int8:
		{
			return Fixed64Of(*v), true
		}
	case *uint32:
		{
			return Fixed64Of(*v), true
		}
	case *uint16:
		{
			return Fixed64Of(*v), true
		}
	case *uint8:
		{
			return Fixed64Of(*v), true
		}
	case *float32:
		{
			return Fixed64Of(*v), true
		}
	}
	// value could not be converted by these conversion functions,
	// pass false down the chain to Fixed64Of()
	return Fixed64{NaN}, false
}

//end
