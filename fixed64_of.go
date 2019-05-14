// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:56:01 444C66                priveda/fixed64/[fixed64_of.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"fmt"
	"reflect"
)

// Fixed64Of converts any compatible value type to a Fixed64.
// This includes all numeric types and strings. If a string is
// not numeric, logs an error and sets the Fixed64 to zero.
func Fixed64Of(value interface{}) Fixed64 {
	switch val := value.(type) {
	case Fixed64:
		{
			return val
		}

	// integers
	case int8:
		{
			return Fixed64{int64(val) * 1E4}
		}
	case int16:
		{
			return Fixed64{int64(val) * 1E4}
		}
	case int32: //                              int32 and int could exceed range
		{
			return Fixed64Of(int64(val))
		}
	case int:
		{
			return Fixed64Of(int64(val))
		}
	case int64:
		{
			if val < -IntLimit || val > IntLimit {
				return fixed64Overflow(val < 0, EOverflow, ": ", val)
			}
			return Fixed64{int64(val) * 1E4}
		}
	// unsigned integers
	case uint8:
		{
			return Fixed64{int64(val) * 1E4}
		}
	case uint16:
		{
			return Fixed64{int64(val) * 1E4}
		}
	case uint32: //                           uint32 and uint could exceed range
		{
			return Fixed64Of(uint64(val))
		}
	case uint:
		{
			return Fixed64Of(uint64(val))
		}
	case uint64:
		{
			if val > IntLimit {
				return fixed64Overflow(false, EOverflow, "uint64: ", val)
			}
			return Fixed64{int64(val) * 1E4}
		}
	// float
	case float32:
		{
			if val < -float32(IntLimit)-0.9999 ||
				val > float32(IntLimit)+0.9999 {
				return fixed64Overflow(val < 0, EOverflow, "float32: ", val)
			}
			return Fixed64{int64(float64(val) * 1E4)}
		}
	case float64:
		{
			if val < -float64(IntLimit)-0.9999 ||
				val > float64(IntLimit)+0.9999 {
				return fixed64Overflow(val < 0, EOverflow, "float64: ", val)
			}
			return Fixed64{int64(val * 1E4)}
		}
	// integer pointers
	case *int:
		if val != nil {
			return Fixed64Of(int64(*val))
		}
	case *int8:
		if val != nil {
			return Fixed64Of(*val)
		}
	case *int16:
		if val != nil {
			return Fixed64Of(*val)
		}
	case *int32:
		if val != nil {
			return Fixed64Of(*val)
		}
	case *int64:
		if val != nil {
			return Fixed64Of(*val)
		}
	// unsigned integer pointers
	case *uint:
		if val != nil {
			return Fixed64Of(uint64(*val))
		}
	case *uint8:
		if val != nil {
			return Fixed64Of(*val)
		}
	case *uint16:
		if val != nil {
			return Fixed64Of(*val)
		}
	case *uint32:
		if val != nil {
			return Fixed64Of(*val)
		}
	case *uint64:
		if val != nil {
			return Fixed64Of(*val)
		}
	// float pointers
	case *float32:
		if val != nil {
			return Fixed64Of(*val)
		}
	case *float64:
		if val != nil {
			return Fixed64Of(*val)
		}
	// strings
	case string:
		{
			return Fixed64OfS(val)
		}
	case *string:
		if val != nil {
			return Fixed64OfS(*val)
		}
	case fmt.Stringer:
		{
			return Fixed64OfS(val.String())
		}
	}
	mod.Error("Type", reflect.TypeOf(value), "not handled; =", value)
	return Fixed64{}
}

//end
