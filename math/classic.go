package math

// Pourcent value
func Pourcent(a, b int64) float32 {
	return (float32(a) / float32(b)) / 100
}

// Diff value
func Diff(a, b int64) int64 {
	return (a - b)
}
