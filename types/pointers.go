package types

import "encoding/base64"

// PBase64toString pointer value
func PBase64toString(value string) *string {
	var (
		dec []byte
		err error
	)

	if dec, err = base64.StdEncoding.DecodeString(value); err != nil {
		return nil
	}

	return PString(string(dec))
}

// PString pointer value
func PString(value string) *string {
	return &value
}

// PBool pointer value
func PBool(value bool) *bool {
	return &value
}

// PInt64 pointer value
func PInt64(value int64) *int64 {
	return &value
}
