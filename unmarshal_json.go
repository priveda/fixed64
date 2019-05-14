// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 16:21:35 F5299D            priveda/fixed64/[unmarshal_json.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"encoding/json"
)

// UnmarshalJSON unmarshals a JSON description of zr.Fixed64.
// This method alters the object's value.
func (ob *Fixed64) UnmarshalJSON(data []byte) error {
	//   ^  don't remove pointer receiver, it is necessary
	if ob == nil {
		return mod.Error(ENilReceiver)
	}
	var (
		num float64
		err = json.Unmarshal(data, &num)
	)
	if err != nil {
		return err
	}
	ob.val = int64(num * 1E4)
	return nil
}

//end
