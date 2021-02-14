// -----------------------------------------------------------------------------
// 64-bit Fixed-precision Number             priveda/fixed64/[unmarshal_json.go]
// (c) admin@priveda.com                                            License: MIT
// -----------------------------------------------------------------------------

package fixed64

import (
	"encoding/json"
)

// UnmarshalJSON unmarshals a JSON description of Fixed64.
// This method alters the number's value.
func (n *Fixed64) UnmarshalJSON(data []byte) error {
	//  ^  don't remove pointer receiver, it is necessary
	if n == nil {
		return mod.Error(ENilReceiver)
	}
	var (
		num float64
		err = json.Unmarshal(data, &num)
	)
	if err != nil {
		return mod.Error(err)
	}
	n.i64 = int64(num * 1e4)
	return nil
}

// end
