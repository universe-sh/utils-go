package math

// MinValue value
func MinValue(a map[string]int64) string {
	var (
		tmpString  string
		tmpInteger int64 = 9223372036854775807
	)

	for k, v := range a {
		if tmpInteger > v {
			tmpInteger = v
			tmpString = k
		}
	}

	return tmpString
}
