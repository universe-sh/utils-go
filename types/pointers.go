package types

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
