// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-20 01:46:11 C0654E                     priveda/fixed64/[new_e.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// NewE converts any compatible value to Fixed64. This includes simple
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
// Fixed64{NaN} and an error. Does not log the error.
//
// When given a nil pointer, also returns Fixed64{NaN}.
//
// Note: fmt.Stringer (or fmt.GoStringer) interfaces are not treated as
// strings to avoid bugs from implicit conversion. Use the String method.
//
func NewE(value interface{}) (Fixed64, error) {
	switch v := value.(type) {
	case int, int64, int32, int16, int8:
		{
			n := reflect.ValueOf(value).Int()
			if n == NaN {
				return Fixed64{NaN}, nil
			}
			if n < -IntLimit || n > IntLimit {
				return fixed64Overflow(n < 0, n), errors.New(EOverflow)
			}
			return Fixed64{n * 1E4}, nil
		}
	case float64, float32:
		{
			n := reflect.ValueOf(value).Float()
			if n < -float64(IntLimit)-0.9999 ||
				n > float64(IntLimit)+0.9999 {
				return fixed64Overflow(n < 0, n), errors.New(EOverflow)
			}
			return Fixed64{int64(n * 1E4)}, nil
		}
	case string:
		{
			n, err := Parse(v)
			return n, err
		}
	case uint, uint64, uint32, uint16, uint8:
		{
			n := reflect.ValueOf(value).Uint()
			if n > IntLimit {
				return fixed64Overflow(false, n), errors.New(EOverflow)
			}
			return Fixed64{int64(n) * 1E4}, nil
		}
	case Fixed64:
		{
			return v, nil
		}
	case bool:
		{
			if v {
				return Fixed64{1 * 1E4}, nil
			}
			return Fixed64{0}, nil
		}
	case nil:
		return Fixed64{NaN}, nil
	}
	// if not converted yet, try to dereference pointer, then convert
	ret := Fixed64{NaN}
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return Fixed64{NaN}, nil
		}
		var err error
		ret, err = NewE(v.Elem().Interface())
		if err == nil || strings.HasPrefix(err.Error(), EOverflow) {
			return ret, nil
		}
	}
	erm := fmt.Sprintf("Can not convert %s to Fixed64: %v",
		reflect.TypeOf(value), value)
	return ret, errors.New(erm)
}

//end
