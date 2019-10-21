package types

import "encoding/json"

// StringEncode for JSON
func StringEncode(a string) string {
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	// Trim the beginning and trailing " character
	return string(b[1 : len(b)-1])
}
