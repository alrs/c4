package asset

import "math/big"

// MarshalJSON adds output support for package encoding/json.
func (id *ID) MarshalJSON() ([]byte, error) {
	bigID := big.Int(*id)
	if bigID.Cmp(big.NewInt(0)) == 0 {
		return []byte(`""`), nil
	}
	return []byte(`"` + id.String() + `"`), nil
}

// UnmarshalJSON adds parsing support for package encoding/json.
func (id *ID) UnmarshalJSON(data []byte) error {
	// UnmarshalJSON includes quotes in the data so we remove them
	id2, err := ParseBytesID(data[1 : len(data)-1])
	*id = *id2
	return err
}

// IDSlice.MarshalJSON adds json output for a slice of IDs.
func (ids *IDSlice) MarshalJSON() ([]byte, error) {
	ids_output := ""
	for _, id := range *ids {
		ids_output = ids_output + id.String()
	}
	return []byte(`"` + ids_output + `"`), nil
}

// IDSlice.UnmarshalJSON adds json parsing for a slice of IDs.
func (ids *IDSlice) UnmarshalJSON(data []byte) error {
	// UnmarshalJSON includes quotes in the data so we remove them
	ids2, err := ParseBytesIDSlice(data[1 : len(data)-1])
	*ids = *ids2
	return err
}
