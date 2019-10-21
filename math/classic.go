package math

// Pourcent value
func Pourcent(a, b int64) float64 {
	if a > b {
		return 0
	} else if a == b {
		return 100
	}

	return (float64(a) / float64(b)) * 100
}

// Diff value
func Diff(a, b int64) int64 {
	return (a - b)
}
