package math

// Pourcent value
func Pourcent(a, b int32) float32 {
	if a > b {
		return 0
	} else if a == b {
		return 100
	}

	return (float32(a) / float32(b)) * 100
}

// Diff value
func Diff(a, b int32) int32 {
	return (a - b)
}
