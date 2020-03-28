package math

// MinValue value
func MinValue(a map[string]int32) string {
	var (
		tmpString  string
		tmpInteger int32 = 2147483647
	)

	for k, v := range a {
		if tmpInteger > v {
			tmpInteger = v
			tmpString = k
		}
	}

	return tmpString
}
